package compliance_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		ComplianceInfoList: []dclcompltypes.ComplianceInfo{
			{
				Vid:               0,
				Pid:               0,
				SoftwareVersion:   0,
				CertificationType: "0",
			},
			{
				Vid:               1,
				Pid:               1,
				SoftwareVersion:   1,
				CertificationType: "1",
			},
		},
		CertifiedModelList: []types.CertifiedModel{
			{
				Vid:               0,
				Pid:               0,
				SoftwareVersion:   0,
				CertificationType: "0",
			},
			{
				Vid:               1,
				Pid:               1,
				SoftwareVersion:   1,
				CertificationType: "1",
			},
		},
		RevokedModelList: []types.RevokedModel{
			{
				Vid:               0,
				Pid:               0,
				SoftwareVersion:   0,
				CertificationType: "0",
			},
			{
				Vid:               1,
				Pid:               1,
				SoftwareVersion:   1,
				CertificationType: "1",
			},
		},
		ProvisionalModelList: []types.ProvisionalModel{
			{
				Vid:               0,
				Pid:               0,
				SoftwareVersion:   0,
				CertificationType: "0",
			},
			{
				Vid:               1,
				Pid:               1,
				SoftwareVersion:   1,
				CertificationType: "1",
			},
		},
		DeviceSoftwareComplianceList: []types.DeviceSoftwareCompliance{
			{
				CDCertificateId: "0",
			},
			{
				CDCertificateId: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	compliance.InitGenesis(ctx, *k, genesisState)
	got := compliance.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.ComplianceInfoList, got.ComplianceInfoList)
	require.ElementsMatch(t, genesisState.CertifiedModelList, got.CertifiedModelList)
	require.ElementsMatch(t, genesisState.RevokedModelList, got.RevokedModelList)
	require.ElementsMatch(t, genesisState.ProvisionalModelList, got.ProvisionalModelList)
	require.ElementsMatch(t, genesisState.DeviceSoftwareComplianceList, got.DeviceSoftwareComplianceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
