package auth

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// SignatureVerificationGasConsumer is the type of function that is used to both consume gas when verifying signatures
// and also to accept or reject different types of PubKey's. This is where apps can define their own PubKey.
type SignatureVerificationGasConsumer = func(meter sdk.GasMeter, pubkey crypto.PubKey) sdk.Result

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers.
//nolint:funlen
func NewAnteHandler(ak Keeper, sigGasConsumer SignatureVerificationGasConsumer) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, simulate bool,
	) (newCtx sdk.Context, res sdk.Result, abort bool) {
		// all transactions must be of type auth.StdTx.
		stdTx, ok := tx.(auth.StdTx)
		if !ok {
			// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
			// during runTx.
			newCtx = SetGasMeter(simulate, ctx, 0)

			return newCtx, sdk.ErrInternal("tx must be StdTx").Result(), true
		}

		newCtx = SetGasMeter(simulate, ctx, stdTx.Fee.Gas)

		// AnteHandlers must have their own defer/recover in order for the BaseApp
		// to know how much gas was used! This is because the GasMeter is created in
		// the AnteHandler, but if it panics the context won't be set properly in
		// runTx's recover call.
		defer func() {
			if r := recover(); r != nil {
				switch rType := r.(type) {
				case sdk.ErrorOutOfGas:
					log := fmt.Sprintf(
						"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
						rType.Descriptor, stdTx.Fee.Gas, newCtx.GasMeter().GasConsumed(),
					)
					res = sdk.ErrOutOfGas(log).Result()

					res.GasWanted = stdTx.Fee.Gas
					res.GasUsed = newCtx.GasMeter().GasConsumed()
					abort = true
				default:
					panic(r)
				}
			}
		}()

		if err := tx.ValidateBasic(); err != nil {
			return newCtx, err.Result(), true
		}

		newCtx.GasMeter().ConsumeGas(types.TxSizeCostPerByte*sdk.Gas(len(newCtx.TxBytes())), "txSize")

		if res := ValidateMemo(stdTx); !res.IsOK() {
			return newCtx, res, true
		}

		// signatures contains the sequence number, account number, and signatures.
		signers := stdTx.GetSigners()
		signatures := stdTx.GetSignatures()

		isGenesis := ctx.BlockHeight() == 0

		for i, signature := range signatures {
			account, res := GetSignerAcc(newCtx, ak, signers[i])
			if !res.IsOK() {
				return newCtx, res, true
			}

			// check signature, return account with incremented nonce.
			signBytes := GetSignBytes(newCtx.ChainID(), stdTx, account, isGenesis)
			account, res = processSig(newCtx, account, signature, signBytes, simulate, sigGasConsumer)

			if !res.IsOK() {
				return newCtx, res, true
			}

			ak.SetAccount(newCtx, account)
		}

		return newCtx, sdk.Result{GasWanted: stdTx.Fee.Gas}, false // continue...
	}
}

// GetSignerAcc returns an account for a given address that is expected to sign a transaction.
func GetSignerAcc(ctx sdk.Context, keeper Keeper, address sdk.AccAddress) (acc types.Account, res sdk.Result) {
	if !keeper.IsAccountPresent(ctx, address) {
		return acc, types.ErrAccountDoesNotExist(address).Result()
	}

	acc = keeper.GetAccount(ctx, address)

	return acc, sdk.Result{}
}

// ValidateMemo validates the memo size.
func ValidateMemo(stdTx auth.StdTx) sdk.Result {
	memoLength := len(stdTx.GetMemo())
	if uint64(memoLength) > types.MaxMemoCharacters {
		return sdk.ErrMemoTooLarge(
			fmt.Sprintf(
				"maximum number of characters is %d but received %d characters",
				types.MaxMemoCharacters, memoLength,
			),
		).Result()
	}

	return sdk.Result{}
}

// verify the signature and increment the sequence.
func processSig(
	ctx sdk.Context, acc types.Account, sig auth.StdSignature, signBytes []byte, simulate bool,
	sigGasConsumer SignatureVerificationGasConsumer,
) (updatedAcc Account, res sdk.Result) {
	if res := sigGasConsumer(ctx.GasMeter(), acc.PubKey); !res.IsOK() {
		return acc, res
	}

	if !simulate && !acc.PubKey.VerifyBytes(signBytes, sig.Signature) {
		return acc, sdk.ErrUnauthorized(
			"Signature verification failed; verify correct account sequence and chain-id").Result()
	}

	acc.Sequence++

	return acc, res
}

// DefaultSigVerificationGasConsumer is the default implementation of SignatureVerificationGasConsumer. It consumes gas
// for signature verification based upon the public key type. The cost is fetched from the given params and is matched
// by the concrete type.
func DefaultSigVerificationGasConsumer(meter sdk.GasMeter, pubkey crypto.PubKey) sdk.Result {
	switch pubkey := pubkey.(type) {
	case ed25519.PubKeyEd25519:
		meter.ConsumeGas(types.DefaultSigVerifyCostED25519, "ante verify: ed25519")

		return sdk.ErrInvalidPubKey("ED25519 public keys are unsupported").Result()

	case secp256k1.PubKeySecp256k1:
		meter.ConsumeGas(types.DefaultSigVerifyCostSecp256k1, "ante verify: secp256k1")

		return sdk.Result{}

	default:
		return sdk.ErrInvalidPubKey(fmt.Sprintf("unrecognized public key type: %T", pubkey)).Result()
	}
}

// SetGasMeter returns a new context with a gas meter set from a given context.
func SetGasMeter(simulate bool, ctx sdk.Context, gasLimit uint64) sdk.Context {
	// In various cases such as simulation and during the genesis block, we do not
	// meter any gas utilization.
	if simulate || ctx.BlockHeight() == 0 {
		return ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	}

	return ctx.WithGasMeter(sdk.NewGasMeter(gasLimit))
}

// GetSignBytes returns a slice of bytes to sign over for a given transaction
// and an account.
func GetSignBytes(chainID string, stdTx auth.StdTx, acc types.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		accNum = acc.AccountNumber
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.Sequence, stdTx.Fee, stdTx.Msgs, stdTx.Memo,
	)
}
