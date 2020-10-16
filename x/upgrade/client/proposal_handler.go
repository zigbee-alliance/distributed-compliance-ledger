package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/upgrade/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
