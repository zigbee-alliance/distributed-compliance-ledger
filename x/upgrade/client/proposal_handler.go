package client

import (
	govclient "github.com/zigbee-alliance/distributed-compliance-ledger/x/gov_old/client"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/upgrade/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
