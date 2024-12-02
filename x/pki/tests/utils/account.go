package utils

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (m *DclauthKeeperMock) HasRole(
	ctx sdk.Context,
	addr sdk.AccAddress,
	roleToCheck dclauthtypes.AccountRole,
) bool {
	args := m.Called(ctx, addr, roleToCheck)

	return args.Bool(0)
}

func (m *DclauthKeeperMock) CountAccountsWithRole(ctx sdk.Context, roleToCount dclauthtypes.AccountRole) int {
	args := m.Called(ctx, roleToCount)

	return args.Int(0)
}

func (m *DclauthKeeperMock) GetAccountO(
	ctx sdk.Context,
	address sdk.AccAddress,
) (val dclauthtypes.Account, found bool) {
	args := m.Called(ctx, address)

	return args.Get(0).(dclauthtypes.Account), args.Bool(1)
}

var _ types.DclauthKeeper = &DclauthKeeperMock{}

func (setup *TestSetup) CreateVendorAccount(vid int32) sdk.AccAddress {
	accAddress := GenerateAccAddress()
	setup.AddAccount(accAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, vid)

	return accAddress
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
	vid int32,
) {
	dclauthKeeper := setup.DclauthKeeper
	currentTrusteeCount := 0
	// if the CountAccountsWithRole is Present get the value from the mock call
	for _, expectedCall := range dclauthKeeper.ExpectedCalls {
		if expectedCall.Method == "CountAccountsWithRole" {
			currentTrusteeCount = dclauthKeeper.CountAccountsWithRole(setup.Ctx, dclauthtypes.Trustee)
		}
	}

	for _, role := range roles {
		dclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
		if role == dclauthtypes.Trustee {
			currentTrusteeCount++
			// We remove the call to CountAccountsWithRole from the expected calls and add it back with the new value
			RemoveItemFromExpectedCalls(dclauthKeeper.ExpectedCalls, "CountAccountsWithRole")
			dclauthKeeper.On("CountAccountsWithRole", setup.Ctx, dclauthtypes.Trustee).Return(currentTrusteeCount)
		}
	}

	dclauthKeeper.On("GetAccountO", setup.Ctx, accAddress).Return(dclauthtypes.Account{VendorID: vid}, true)
	dclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()

	return accAddress
}

// Remove a item from ExpectedCalls Array and return it.
func RemoveItemFromExpectedCalls(expectedCalls []*mock.Call, methodName string) {
	for i, call := range expectedCalls {
		if call.Method == methodName {
			expectedCalls = append(expectedCalls[:i], expectedCalls[i+1:]...)
		}
	}
}

func (setup *TestSetup) CreateNTrusteeAccounts() ([]sdk.AccAddress, int) {
	// Create an array of trustee account from 1 to 50
	trusteeAccounts := make([]sdk.AccAddress, 50)
	for i := 0; i < 50; i++ {
		trusteeAccounts[i] = GenerateAccAddress()
	}

	totalAdditionalTrustees := rand.Intn(50)
	for i := 0; i < totalAdditionalTrustees; i++ {
		setup.AddAccount(trusteeAccounts[i], []dclauthtypes.AccountRole{dclauthtypes.Trustee}, 1)
	}

	return trusteeAccounts, totalAdditionalTrustees
}
