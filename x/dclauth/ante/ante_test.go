package ante_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/ante"
)

func TestDefaultSigVerificationGasConsumer(t *testing.T) {
	params := authtypes.DefaultParams()
	meter := storetypes.NewInfiniteGasMeter()

	// ED25519 keys are explicitly unsupported.
	err := ante.DefaultSigVerificationGasConsumer(meter,
		signing.SignatureV2{PubKey: ed25519.GenPrivKey().PubKey()}, params)
	require.Error(t, err)

	// secp256k1 is accepted.
	err = ante.DefaultSigVerificationGasConsumer(meter,
		signing.SignatureV2{PubKey: secp256k1.GenPrivKey().PubKey()}, params)
	require.NoError(t, err)

	// secp256r1 is accepted.
	r1, genErr := secp256r1.GenPrivKey()
	require.NoError(t, genErr)
	err = ante.DefaultSigVerificationGasConsumer(meter,
		signing.SignatureV2{PubKey: r1.PubKey()}, params)
	require.NoError(t, err)

	// Any other (here: nil) public key type is rejected.
	err = ante.DefaultSigVerificationGasConsumer(meter,
		signing.SignatureV2{PubKey: nil}, params)
	require.Error(t, err)
}

func TestNewAnteHandler_RequiresAccountKeeper(t *testing.T) {
	_, err := ante.NewAnteHandler(ante.HandlerOptions{})
	require.Error(t, err)
}
