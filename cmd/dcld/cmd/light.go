package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/spf13/cobra"
	dbm "github.com/tendermint/tm-db"

	tmcfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	tmmath "github.com/tendermint/tendermint/libs/math"
	"github.com/tendermint/tendermint/light"
	lproxy "github.com/tendermint/tendermint/light/proxy"
	lrpc "github.com/tendermint/tendermint/light/rpc"
	dbs "github.com/tendermint/tendermint/light/store/db"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	rpcserver "github.com/tendermint/tendermint/rpc/jsonrpc/server"
)

// mostly copied from https://github.com/tendermint/tendermint/blob/master/cmd/tendermint/commands/light.go

const (
	FlagListenAddr       = "laddr"
	FlagPrimary          = "primary"
	FlagPrimaryShort     = "p"
	FlagWitness          = "witnesses"
	FlagWitnessShort     = "w"
	FlagHeight           = "height"
	FlagHash             = "hash"
	FlagDir              = "dir"
	FlagDirShort         = "d"
	FlagLogLevel         = "log-level"
	FlagSeq              = "sequential"
	FlagTrustLevel       = "trust-level"
	FlagMaxConn          = "max-open-connections"
	FlagTrustPeriod      = "trusting-period"
	FlagStartTimeout     = "start-timeout"
	FlagTLSCertFile      = "tls-cert-file"
	FlagTLSCertFileShort = "c"
	FlagTLSKeyFile       = "tls-key-file"
	FlagTLSKeyFileShort  = "k"
)

// LightCmd represents the base command when called without any subcommands.
var LightCmd = &cobra.Command{
	Use:   "light [chainID]",
	Short: "Run a light client proxy server, verifying Tendermint rpc",
	Long: `Run a light client proxy server, verifying Tendermint rpc.

All calls that can be tracked back to a block header by a proof
will be verified before passing them back to the caller. Other than
that, it will present the same interface as a full Tendermint node.

Furthermore to the chainID, a fresh instance of a light client will
need a primary RPC address, a trusted hash and height and witness RPC addresses
(if not using sequential verification). To restart the node, thereafter
only the chainID is required.

When /abci_query is called, the Merkle key path format is:

	/{store name}/{key}

Please verify with your application that this Merkle key format is used (true
for applications built w/ Cosmos SDK).
`,
	RunE: runProxy,
	Args: cobra.ExactArgs(1),
	Example: `dcld light dcltestnet -p http://52.57.29.196:26657 -w http://public-seed-node.cosmoshub.certus.one:26657
	--height 962118 --hash 28B97BE9F6DE51AC69F70E0B7BFD7E5C9CD1A595B7DC31AFF27C50D4948020CD`,
}

var (
	listenAddr         string
	primaryAddr        string
	witnessAddrsJoined string
	chainID            string
	dir                string
	maxOpenConnections int

	sequential     bool
	trustingPeriod time.Duration
	trustedHeight  int64
	trustedHash    []byte
	trustLevelStr  string

	logLevel string
	// logFormat string.

	primaryKey   = []byte("primary")
	witnessesKey = []byte("witnesses")

	startTimeout int64

	tlsCertFile string
	tlsKeyFile  string
)

func init() {
	LightCmd.Flags().StringVar(&listenAddr, FlagListenAddr, "tcp://0.0.0.0:8888",
		"serve the proxy on the given address")
	LightCmd.Flags().StringVarP(&primaryAddr, FlagPrimary, FlagPrimaryShort, "",
		"connect to a Tendermint node at this address")
	LightCmd.Flags().StringVarP(&witnessAddrsJoined, FlagWitness, FlagWitnessShort, "",
		"tendermint nodes to cross-check the primary node, comma-separated")
	LightCmd.Flags().StringVarP(&dir, FlagDir, FlagDirShort, os.ExpandEnv(filepath.Join("$HOME", ".tendermint-light")),
		"specify the directory")
	LightCmd.Flags().IntVar(
		&maxOpenConnections,
		FlagMaxConn,
		900,
		"maximum number of simultaneous connections (including WebSocket).")
	LightCmd.Flags().DurationVar(&trustingPeriod, FlagTrustPeriod, 168*time.Hour,
		"trusting period that headers can be verified within. Should be significantly less than the unbonding period")
	LightCmd.Flags().Int64Var(&trustedHeight, FlagHeight, 0, "Trusted header's height")
	LightCmd.Flags().BytesHexVar(&trustedHash, FlagHash, []byte{}, "Trusted header's hash")
	LightCmd.Flags().StringVar(&logLevel, FlagLogLevel, "info", "The logging level (debug|info|warn|error|fatal)")
	// LightCmd.Flags().StringVar(&logFormat, "log-format", log.LogFormatPlain, "The logging format (text|json)")
	LightCmd.Flags().StringVar(&trustLevelStr, FlagTrustLevel, "1/3",
		"trust level. Must be between 1/3 and 3/3",
	)
	LightCmd.Flags().BoolVar(&sequential, FlagSeq, false,
		"sequential verification. Verify all headers sequentially as opposed to using skipping verification",
	)
	LightCmd.Flags().Int64Var(&startTimeout, FlagStartTimeout, 0,
		"How many seconds to wait before starting the light client proxy. Mostly for test purposes when light client is started at the same time as the pool.")

	LightCmd.Flags().StringVarP(&tlsCertFile, FlagTLSCertFile, FlagTLSCertFileShort, "", "Path to the TLS certificate file")
	LightCmd.Flags().StringVarP(&tlsKeyFile, FlagTLSKeyFile, FlagTLSKeyFileShort, "", "Path to the TLS key file")
}

func runProxy(cmd *cobra.Command, args []string) error {
	time.Sleep(time.Duration(startTimeout) * time.Second)

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	var option log.Option
	if logLevel == "info" {
		option, _ = log.AllowLevel("info")
	} else {
		option, _ = log.AllowLevel("debug")
	}
	logger = log.NewFilter(logger, option)

	chainID = args[0]
	logger.Info("Creating client...", "chainID", chainID)

	witnessesAddrs := []string{}
	if witnessAddrsJoined != "" {
		witnessesAddrs = strings.Split(witnessAddrsJoined, ",")
	}

	lightDB, err := dbm.NewGoLevelDB("light-client-db", dir)
	if err != nil {
		return fmt.Errorf("can't create a db: %w", err)
	}
	// create a prefixed db on the chainID
	db := dbm.NewPrefixDB(lightDB, []byte(chainID))

	if primaryAddr == "" { // check to see if we can start from an existing state
		var err error
		primaryAddr, witnessesAddrs, err = checkForExistingProviders(db)
		if err != nil {
			return fmt.Errorf("failed to retrieve primary or witness from db: %w", err)
		}
		if primaryAddr == "" {
			return errors.New("no primary address was provided nor found. Please provide a primary (using -p)." +
				" Run the command: tendermint light --help for more information")
		}
	} else {
		err := saveProviders(db, primaryAddr, witnessAddrsJoined)
		if err != nil {
			logger.Error("Unable to save primary and or witness addresses", "err", err)
		}
	}

	trustLevel, err := tmmath.ParseFraction(trustLevelStr)
	if err != nil {
		return fmt.Errorf("can't parse trust level: %w", err)
	}

	options := []light.Option{light.Logger(logger)}

	if sequential {
		options = append(options, light.SequentialVerification())
	} else {
		options = append(options, light.SkippingVerification(trustLevel))
	}

	// should set either both hash and height, or none of them
	if trustedHeight == 0 && len(trustedHash) > 0 {
		return fmt.Errorf("hash is set but height is not (--%s) ", FlagHeight)
	}
	if trustedHeight != 0 && len(trustedHash) == 0 {
		return fmt.Errorf("height is set but hash is not (--%s) ", FlagHash)
	}

	// if trusted height and hash are not set, try to get it from primaries and witness
	if trustedHeight == 0 && len(trustedHash) == 0 {
		trustedHeight, trustedHash, err = getTrustedHeightAndHash(primaryAddr, witnessesAddrs)
		if err != nil {
			return err
		}
	}

	// Initiate the light client. If the trusted store already has blocks in it, this
	// will be used else we use the trusted options.
	c, err := light.NewHTTPClient(
		context.Background(),
		chainID,
		light.TrustOptions{
			Period: trustingPeriod,
			Height: trustedHeight,
			Hash:   trustedHash,
		},
		primaryAddr,
		witnessesAddrs,
		dbs.New(db, chainID),
		options...,
	)
	if err != nil {
		return err
	}

	config := tmcfg.DefaultConfig()
	cfg := rpcserver.DefaultConfig()
	cfg.MaxBodyBytes = config.RPC.MaxBodyBytes
	cfg.MaxHeaderBytes = config.RPC.MaxHeaderBytes
	cfg.MaxOpenConnections = maxOpenConnections
	// If necessary adjust global WriteTimeout to ensure it's greater than
	// TimeoutBroadcastTxCommit.
	// See https://github.com/tendermint/tendermint/issues/3435
	if cfg.WriteTimeout <= config.RPC.TimeoutBroadcastTxCommit {
		cfg.WriteTimeout = config.RPC.TimeoutBroadcastTxCommit + 10*time.Second
	}
	cfg.WriteTimeout = config.RPC.TimeoutBroadcastTxCommit + 20*time.Second
	cfg.ReadTimeout = 20 * time.Second

	p, err := lproxy.NewProxy(c, listenAddr, primaryAddr, cfg, logger, lrpc.KeyPathFn(lrpc.DefaultMerkleKeyPathFn()))
	if err != nil {
		return err
	}

	// IMPORTANT: add support for IAVL (cosmos) proofs (Tendermint is not aware of them by default)
	p.Client.RegisterOpDecoder(storetypes.ProofOpIAVLCommitment, storetypes.CommitmentOpDecoder)
	p.Client.RegisterOpDecoder(storetypes.ProofOpSimpleMerkleCommitment, storetypes.CommitmentOpDecoder)

	ctx, cancel := signal.NotifyContext(cmd.Context(), syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		p.Listener.Close()
	}()

	if tlsCertFile != "" && tlsKeyFile != "" {
		logger.Info("Starting Light Client Proxy over TLS/HTTPS...", "laddr", listenAddr)
		err = p.ListenAndServeTLS(tlsCertFile, tlsKeyFile)
	} else {
		logger.Info("Starting Light Client Proxy over HTTP...", "laddr", listenAddr)
		err = p.ListenAndServe()
	}

	if errors.Is(err, http.ErrServerClosed) {
		// Error starting or closing listener:
		logger.Error("proxy ListenAndServe", "err", err)
	}

	return nil
}

func checkForExistingProviders(db dbm.DB) (string, []string, error) {
	primaryBytes, err := db.Get(primaryKey)
	if err != nil {
		return "", []string{""}, err
	}
	witnessesBytes, err := db.Get(witnessesKey)
	if err != nil {
		return "", []string{""}, err
	}
	witnessesAddrs := strings.Split(string(witnessesBytes), ",")

	return string(primaryBytes), witnessesAddrs, nil
}

func saveProviders(db dbm.DB, primaryAddr, witnessesAddrs string) error {
	err := db.Set(primaryKey, []byte(primaryAddr))
	if err != nil {
		return fmt.Errorf("failed to save primary provider: %w", err)
	}
	err = db.Set(witnessesKey, []byte(witnessesAddrs))
	if err != nil {
		return fmt.Errorf("failed to save witness providers: %w", err)
	}

	return nil
}

func getTrustedHeightAndHash(primaryAddr string, witnessesAddrs []string) (int64, []byte, error) {
	config := rpcserver.DefaultConfig()

	// get height and hash from the primary
	primaryRPCClient, err := rpchttp.NewWithTimeout(primaryAddr, "/websocket", uint(config.WriteTimeout.Seconds()))
	if err != nil {
		return 0, []byte{}, fmt.Errorf("not able to obtain trusted height and hash: %w", err)
	}

	res, err := primaryRPCClient.Commit(context.Background(), nil)
	if err != nil {
		return 0, []byte{}, fmt.Errorf("not able to obtain trusted height and hash: %w", err)
	}
	primaryHeight := res.SignedHeader.Header.Height
	primaryHash := res.SignedHeader.Commit.BlockID.Hash

	time.Sleep(2 * time.Second) // sleep to make sure a new block is committed on all nodes

	// check that the hash for the given height is the same on all witnesses
	for _, witnessesAddr := range witnessesAddrs {
		witnessRPCClient, err := rpchttp.NewWithTimeout(witnessesAddr, "/websocket", uint(config.WriteTimeout.Seconds()))
		if err != nil {
			return 0, []byte{}, fmt.Errorf("not able to obtain trusted height and hash: %w", err)
		}

		res, err := witnessRPCClient.Commit(context.Background(), &primaryHeight)
		if err != nil {
			return 0, []byte{}, fmt.Errorf("not able to obtain trusted height and hash: %w", err)
		}
		h := res.SignedHeader.Header.Height
		hash := res.SignedHeader.Commit.BlockID.Hash

		if h != primaryHeight {
			return 0, []byte{}, fmt.Errorf(
				"not able to obtain trusted height and hash: primary height %d is not equal to witness height %d",
				primaryHeight, h,
			)
		}
		if !bytes.Equal(hash, primaryHash) {
			return 0, []byte{}, fmt.Errorf(
				"not able to obtain trusted height and hash: primary hash %s is not equal to witness hash %s atr height %d",
				primaryHash, hash, primaryHeight,
			)
		}
	}

	return primaryHeight, primaryHash, nil
}
