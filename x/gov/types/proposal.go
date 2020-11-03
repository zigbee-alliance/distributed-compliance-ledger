package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Proposal defines a struct used by the governance module to allow for voting.
type Proposal struct {
	Content `json:"content" yaml:"content"` // Proposal content interface

	ProposalID uint64         `json:"id" yaml:"id"`                           //  ID of the proposal
	Status     ProposalStatus `json:"proposal_status" yaml:"proposal_status"` // Status of the Proposal {Active, Passed, Rejected}
}

func NewProposal(content Content, id uint64) Proposal {
	return Proposal{
		Content:    content,
		ProposalID: id,
		Status:     StatusActive,
	}
}

// nolint
func (p Proposal) String() string {
	return fmt.Sprintf(`Proposal %d:
  Title:              %s
  Type:               %s
  Status:             %s
  Description:        %s`,
		p.ProposalID, p.GetTitle(), p.ProposalType(),
		p.Status, p.GetDescription(),
	)
}

// Proposals is an array of proposal
type Proposals []Proposal

func (p Proposals) String() string {
	out := "ID - (Status) [Type] Title\n"
	for _, prop := range p {
		out += fmt.Sprintf("%d - (%s) [%s] %s\n",
			prop.ProposalID, prop.Status,
			prop.ProposalType(), prop.GetTitle())
	}
	return strings.TrimSpace(out)
}

type (
	// ProposalStatus is a type alias that represents a proposal status as a byte
	ProposalStatus byte
)

const (
	StatusActive   ProposalStatus = 0x00
	StatusPassed   ProposalStatus = 0x01
	StatusRejected ProposalStatus = 0x02
)

// ProposalStatusToString turns a string into a ProposalStatus
func ProposalStatusFromString(str string) (ProposalStatus, error) {
	switch str {
	case "Active":
		return StatusActive, nil

	case "Passed":
		return StatusPassed, nil

	case "Rejected":
		return StatusRejected, nil

	default:
		return ProposalStatus(0xff), fmt.Errorf("'%s' is not a valid proposal status", str)
	}
}

// ValidProposalStatus returns true if the proposal status is valid and false
// otherwise.
func ValidProposalStatus(status ProposalStatus) bool {
	if status == StatusActive ||
		status == StatusPassed ||
		status == StatusRejected {
		return true
	}
	return false
}

// Marshal needed for protobuf compatibility
func (status ProposalStatus) Marshal() ([]byte, error) {
	return []byte{byte(status)}, nil
}

// Unmarshal needed for protobuf compatibility
func (status *ProposalStatus) Unmarshal(data []byte) error {
	*status = ProposalStatus(data[0])
	return nil
}

// Marshals to JSON using string
func (status ProposalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.String())
}

// Unmarshals from JSON assuming Bech32 encoding
func (status *ProposalStatus) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	bz2, err := ProposalStatusFromString(s)
	if err != nil {
		return err
	}

	*status = bz2
	return nil
}

// String implements the Stringer interface.
func (status ProposalStatus) String() string {
	switch status {
	case StatusActive:
		return "Active"

	case StatusPassed:
		return "Passed"

	case StatusRejected:
		return "Rejected"

	default:
		return ""
	}
}

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (status ProposalStatus) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(status.String()))
	default:
		// TODO: Do this conversion more directly
		s.Write([]byte(fmt.Sprintf("%v", byte(status))))
	}
}
