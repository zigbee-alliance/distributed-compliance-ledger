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
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	hexVid  = "0xA13" // = 2579
	hexVid2 = "0xA14" // = 2580
	hexVid3 = "0xA15" // = 2581

	hexVidDecimal  = 2579
	hexVid2Decimal = 2580
)

func TestVendorInfoDemoHex(t *testing.T) {
	vendorAccount := fmt.Sprintf("vendor_account_%s", hexVid)
	secondVendorAccount := fmt.Sprintf("vendor_account_%s", hexVid2)

	cliputils.CreateVendorAccount(t, vendorAccount, hexVidDecimal)
	cliputils.CreateVendorAccount(t, secondVendorAccount, hexVid2Decimal)

	const (
		companyLegalName = "XYZ IOT Devices Inc"
		vendorName       = "XYZ Devices"
	)

	t.Run("AddVendorInfoWithHexVID", func(t *testing.T) {
		txResult, err := AddVendor(vendorAccount, VendorOpts{
			VIDHex:           hexVid,
			CompanyLegalName: companyLegalName,
			VendorName:       vendorName,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryVendorInfoWithHexVID", func(t *testing.T) {
		out, err := QueryVendor(hexVid)
		require.NoError(t, err)
		// The ledger stores the integer value, not the hex string
		require.Contains(t, string(out), fmt.Sprintf(`"vendorID":%d`, hexVidDecimal))
		require.Contains(t, string(out), fmt.Sprintf(`"companyLegalName":"%s"`, companyLegalName))
		require.Contains(t, string(out), fmt.Sprintf(`"vendorName":"%s"`, vendorName))
	})

	updatedCompanyName := "ABC Subsidiary Corporation"
	vendorLandingPageURL := "https://www.w3.org/"

	t.Run("UpdateVendorInfoWithHexVID", func(t *testing.T) {
		txResult, err := UpdateVendor(vendorAccount, VendorOpts{
			VIDHex:               hexVid,
			CompanyLegalName:     updatedCompanyName,
			VendorLandingPageURL: vendorLandingPageURL,
			VendorName:           vendorName,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryVendor(hexVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vendorID":%d`, hexVidDecimal))
		require.Contains(t, string(out), fmt.Sprintf(`"companyLegalName":"%s"`, updatedCompanyName))
		require.Contains(t, string(out), fmt.Sprintf(`"vendorName":"%s"`, vendorName))
		require.Contains(t, string(out), fmt.Sprintf(`"vendorLandingPageURL":"%s"`, vendorLandingPageURL))
	})

	t.Run("AddVendorForWrongHexVID_Fails", func(t *testing.T) {
		txResult, err := AddVendor(vendorAccount, VendorOpts{
			VIDHex:           hexVid3,
			CompanyLegalName: updatedCompanyName,
			VendorName:       vendorName,
		})
		if err == nil {
			require.NotEqual(t, uint32(0), txResult.Code)
		}
	})

	t.Run("UpdateVendorForWrongHexAccount_Fails", func(t *testing.T) {
		txResult, err := UpdateVendor(secondVendorAccount, VendorOpts{
			VIDHex:           hexVid,
			CompanyLegalName: updatedCompanyName,
			VendorName:       vendorName,
		})
		if err == nil {
			require.NotEqual(t, uint32(0), txResult.Code)
		}
	})
}
