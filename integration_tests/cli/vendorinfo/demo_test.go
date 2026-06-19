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
		out, err := QueryVendor(fmt.Sprintf("%d", vid))
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("QueryAllEmpty", func(t *testing.T) {
		// This test's specific vendor must not exist yet (other tests may have added different VIDs).
		out, err := QueryVendor(fmt.Sprintf("%d", vid))
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
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
		out, err := QueryVendor(fmt.Sprintf("%d", vid))
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vendorID":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"companyLegalName":"%s"`, companyLegalName))
		require.Contains(t, string(out), fmt.Sprintf(`"vendorName":"%s"`, vendorName))
		require.Contains(t, string(out), fmt.Sprintf(`"schemaVersion":%s`, schemaVersion0))
	})

	t.Run("QueryAllVendors", func(t *testing.T) {
		out, err := QueryAllVendors()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vendorID":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"companyLegalName":"%s"`, companyLegalName))
		require.Contains(t, string(out), fmt.Sprintf(`"vendorName":"%s"`, vendorName))
	})

	t.Run("UpdateVendorInfoRequiredFieldsOnly", func(t *testing.T) {
		txResult, err := UpdateVendor(vendorAccount, VendorOpts{VID: vid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Omitted optional fields should keep their previous values
		out, err := QueryVendor(fmt.Sprintf("%d", vid))
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vendorID":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"companyLegalName":"%s"`, companyLegalName))
		require.Contains(t, string(out), fmt.Sprintf(`"vendorName":"%s"`, vendorName))
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

		out, err := QueryVendor(fmt.Sprintf("%d", vid))
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vendorID":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"companyLegalName":"%s"`, updatedCompanyLegalName))
		require.Contains(t, string(out), fmt.Sprintf(`"vendorName":"%s"`, vendorName))
		require.Contains(t, string(out), fmt.Sprintf(`"vendorLandingPageURL":"%s"`, vendorLandingPageURL))
	})

	t.Run("AddVendorForWrongVID_Fails", func(t *testing.T) {
		vid1 := rand.Intn(60000) + 61000
		txResult, err := AddVendor(vendorAccount, VendorOpts{
			VID:              vid1,
			CompanyLegalName: updatedCompanyLegalName,
			VendorName:       vendorName,
		})
		// Either execution error or non-zero tx code
		if err == nil {
			require.NotEqual(t, uint32(0), txResult.Code)
		}
	})

	t.Run("UpdateVendorForWrongAccount_Fails", func(t *testing.T) {
		txResult, err := UpdateVendor(secondVendorAccount, VendorOpts{
			VID:              vid,
			CompanyLegalName: updatedCompanyLegalName,
			VendorName:       vendorName,
		})
		if err == nil {
			require.NotEqual(t, uint32(0), txResult.Code)
		}
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
