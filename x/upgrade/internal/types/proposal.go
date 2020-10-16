package types

import (
	"fmt"
)

// Software Upgrade Proposals
type SoftwareUpgradeProposal struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	Plan        Plan   `json:"plan" yaml:"plan"`
}

func NewSoftwareUpgradeProposal(title, description string, plan Plan) SoftwareUpgradeProposal {
	return SoftwareUpgradeProposal{title, description, plan}
}

// Implements Proposal Interface
var _ Content = SoftwareUpgradeProposal{}

// nolint
func (sup SoftwareUpgradeProposal) GetTitle() string       { return sup.Title }
func (sup SoftwareUpgradeProposal) GetDescription() string { return sup.Description }
func (sup SoftwareUpgradeProposal) ValidateBasic() error {
	if err := sup.Plan.ValidateBasic(); err != nil {
		return err
	}
	return ValidateAbstract(sup)
}

func (sup SoftwareUpgradeProposal) String() string {
	return fmt.Sprintf(`Software Upgrade Proposal:
  Title:       %s
  Description: %s
`, sup.Title, sup.Description)
}

// Cancel Software Upgrade Proposals
type CancelSoftwareUpgradeProposal struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
}

func NewCancelSoftwareUpgradeProposal(title, description string) CancelSoftwareUpgradeProposal {
	return CancelSoftwareUpgradeProposal{title, description}
}

// Implements Proposal Interface
var _ Content = CancelSoftwareUpgradeProposal{}

// nolint
func (sup CancelSoftwareUpgradeProposal) GetTitle() string       { return sup.Title }
func (sup CancelSoftwareUpgradeProposal) GetDescription() string { return sup.Description }
func (sup CancelSoftwareUpgradeProposal) ValidateBasic() error {
	return ValidateAbstract(sup)
}

func (sup CancelSoftwareUpgradeProposal) String() string {
	return fmt.Sprintf(`Cancel Software Upgrade Proposal:
  Title:       %s
  Description: %s
`, sup.Title, sup.Description)
}
