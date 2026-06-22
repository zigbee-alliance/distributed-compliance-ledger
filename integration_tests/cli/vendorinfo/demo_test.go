// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vendorinfo

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	vendorinfotypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func TestVendorInfoDemo(t *testing.T) {
	vid := rand.Intn(60000) + 1
	vid2 := rand.Intn(60000) + 1

	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	secondVendorAccount := fmt.Sprintf("vendor_account_%d", vid2)

	cliputils.CreateVendorAccount(t, vendorAccount, vid)
	cliputils.CreateVendorAccount(t, secondVendorAccount, vid2)

	// Create a VendorAdmin account
	vendorAdminAccount := cliputils.CreateAccount(t, "VendorAdmin")

	t.Run("QueryNonExistent", func(t *testing.T) {
		v, err := GetVendor(vid)
		require.NoError(t, err)
		require.Nil(t, v)
	})

	t.Run("QueryAllEmpty", func(t *testing.T) {
		// This test's specific vendor must not exist yet (other tests may have
		// added different VIDs, so the global all-vendors list is not literally
		// empty on the shared ledger — assert only that our VID is absent).
		v, err := GetVendor(vid)
		require.NoError(t, err)
		require.Nil(t, v)

		all, err := GetAllVendors()
		require.NoError(t, err)
		require.False(t, containsVendorByID(all, int32(vid)))
	})

	const (
		companyLegalName = "XYZ IOT Devices Inc"
		vendorName       = "XYZ Devices"
		schemaVersion0   = "0"
	)

	t.Run("AddVendorInfo", func(t *testing.T) {
		txResult, err := AddVendor(vendorAccount, VendorOpts{
			VID:              vid,
			CompanyLegalName: companyLegalName,
			VendorName:       vendorName,
			SchemaVersion:    schemaVersion0,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryVendorInfo", func(t *testing.T) {
		v, err := GetVendor(vid)
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, int32(vid), v.VendorID)
		require.Equal(t, companyLegalName, v.CompanyLegalName)
		require.Equal(t, vendorName, v.VendorName)
		require.Equal(t, uint32(0), v.SchemaVersion)
	})

	t.Run("QueryAllVendors", func(t *testing.T) {
		all, err := GetAllVendors()
		require.NoError(t, err)
		require.True(t, containsVendorByID(all, int32(vid)))
		var got *vendorinfotypes.VendorInfo
		for i := range all {
			if all[i].VendorID == int32(vid) {
				got = &all[i]

				break
			}
		}
		require.NotNil(t, got)
		require.Equal(t, companyLegalName, got.CompanyLegalName)
		require.Equal(t, vendorName, got.VendorName)
	})

	t.Run("UpdateVendorInfoRequiredFieldsOnly", func(t *testing.T) {
		txResult, err := UpdateVendor(vendorAccount, VendorOpts{VID: vid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Omitted optional fields should keep their previous values
		v, err := GetVendor(vid)
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, int32(vid), v.VendorID)
		require.Equal(t, companyLegalName, v.CompanyLegalName)
		require.Equal(t, vendorName, v.VendorName)
	})

	updatedCompanyLegalName := "ABC Subsidiary Corporation"
	vendorLandingPageURL := "https://www.w3.org/"

	t.Run("UpdateVendorInfoAllFields", func(t *testing.T) {
		txResult, err := UpdateVendor(vendorAccount, VendorOpts{
			VID:                  vid,
			CompanyLegalName:     updatedCompanyLegalName,
			VendorLandingPageURL: vendorLandingPageURL,
			VendorName:           vendorName,
			SchemaVersion:        schemaVersion0,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		v, err := GetVendor(vid)
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, int32(vid), v.VendorID)
		require.Equal(t, updatedCompanyLegalName, v.CompanyLegalName)
		require.Equal(t, vendorName, v.VendorName)
		require.Equal(t, vendorLandingPageURL, v.VendorLandingPageURL)
		require.Equal(t, uint32(0), v.SchemaVersion)
	})

	t.Run("AddVendorForWrongVID_Fails", func(t *testing.T) {
		// vid1 must be a *valid* VID (<= 65535) that differs from the vendor
		// account's VID, so the add reaches the vendor-association check rather
		// than the VID upper-bound validation. vid/vid2 are both <= 60000, so
		// [60001, 65535] is always distinct.
		vid1 := rand.Intn(5535) + 60001
		txResult, err := AddVendor(vendorAccount, VendorOpts{
			VID:              vid1,
			CompanyLegalName: updatedCompanyLegalName,
			VendorName:       vendorName,
		})
		require.Contains(t, txFailureText(txResult, err),
			fmt.Sprintf("transaction should be signed by a vendor account associated with the vendorID %d", vid1))
	})

	t.Run("UpdateVendorForWrongAccount_Fails", func(t *testing.T) {
		// secondVendorAccount (vid2) cannot update vid's record.
		txResult, err := UpdateVendor(secondVendorAccount, VendorOpts{
			VID:              vid,
			CompanyLegalName: updatedCompanyLegalName,
			VendorName:       vendorName,
		})
		require.Contains(t, txFailureText(txResult, err),
			fmt.Sprintf("transaction should be signed by a vendor account associated with the vendorID %d", vid))
	})

	t.Run("AddVendorByVendorAdmin", func(t *testing.T) {
		adminVid := rand.Intn(60000) + 1
		txResult, err := AddVendor(vendorAdminAccount, VendorOpts{
			VID:              adminVid,
			CompanyLegalName: updatedCompanyLegalName,
			VendorName:       vendorName,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Update the same record by vendor admin
		newCompanyName := "New Corp"
		newVendorName := "New Vendor Name"
		txResult, err = UpdateVendor(vendorAdminAccount, VendorOpts{
			VID:              adminVid,
			CompanyLegalName: newCompanyName,
			VendorName:       newVendorName,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})
}
