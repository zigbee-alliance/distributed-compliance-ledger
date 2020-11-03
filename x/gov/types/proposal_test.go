package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func init() {
	RegisterProposalTypeCodec(TestContent{}, "testContent")
}

type TestContent struct {
	title       string
	description string
	route       string
	type_       string
}

func (t TestContent) GetTitle() string {
	return t.title
}

func (t TestContent) GetDescription() string {
	return t.description
}

func (t TestContent) ProposalRoute() string {
	return t.route
}

func (t TestContent) ProposalType() string {
	return t.type_
}

func (t TestContent) ValidateBasic() sdk.Error {
	return nil
}

func (t TestContent) String() string {
	return fmt.Sprintf("Title: %v, sedcription: %v", t.title, t.description)
}

func TestJsonCodec(t *testing.T) {
	proposal := Proposal{
		Content: TestContent{
			title:       "title",
			description: "description",
			route:       "route",
			type_:       "type",
		},
		ProposalID: 34,
		Status:     StatusPassed,
	}

	json := ModuleCdc.MustMarshalJSON(proposal)
	decoded := Proposal{}
	ModuleCdc.MustUnmarshalJSON(json, &decoded)

	assert.ObjectsAreEqual(proposal, decoded)
}
