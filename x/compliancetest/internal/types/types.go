package types

import (
	"encoding/json"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TestingResult struct {
	VID        int16          `json:"vid"`
	PID        int16          `json:"pid"`
	TestResult string         `json:"test_result"`
	Owner      sdk.AccAddress `json:"owner"`
}

func NewTestingResult(vid int16, pid int16, testResult string, owner sdk.AccAddress) TestingResult {
	return TestingResult{
		VID:        vid,
		PID:        pid,
		TestResult: testResult,
		Owner:      owner,
	}
}

func (d TestingResult) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

type TestingResults struct {
	VID     int16               `json:"vid"`
	PID     int16               `json:"pid"`
	Results []TestingResultItem `json:"results"`
}

func NewTestingResults(vid int16, pid int16) TestingResults {
	return TestingResults{
		VID:     vid,
		PID:     pid,
		Results: []TestingResultItem{},
	}
}

func (d TestingResults) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (d *TestingResults) AddTestingResult(testingResult TestingResult) {
	d.Results = append(d.Results,
		TestingResultItem{
			TestResult: testingResult.TestResult,
			Owner:      testingResult.Owner,
		})
}

func (d *TestingResults) ContainsTestingResult(owner sdk.AccAddress) bool {
	for _, item := range d.Results {
		if item.Owner.Equals(owner) {
			return true
		}
	}
	return false
}

type TestingResultItem struct {
	TestResult string         `json:"test_result"`
	Owner      sdk.AccAddress `json:"owner"`
}

func NewTestingResultItem(testResult string, owner sdk.AccAddress) TestingResultItem {
	return TestingResultItem{
		TestResult: testResult,
		Owner:      owner,
	}
}

func (d TestingResultItem) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func ParseVID(str string) (int16, error) {
	return conversions.ParseInt16FromString(str)
}

func ParsePID(str string) (int16, error) {
	return conversions.ParseInt16FromString(str)
}
