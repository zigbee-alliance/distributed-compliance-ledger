package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type TestingResult struct {
	VID        int16          `json:"vid"`
	PID        int16          `json:"pid"`
	TestResult string         `json:"test_result"`
	Owner      sdk.AccAddress `json:"owner"`
	CreatedAt  time.Time      `json:"created_at"` // creation time - sets automatically
}

func NewTestingResult(vid int16, pid int16, testResult string, owner sdk.AccAddress) TestingResult {
	return TestingResult{
		VID:        vid,
		PID:        pid,
		TestResult: testResult,
		Owner:      owner,
		CreatedAt:  time.Now(),
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
			CreatedAt:  testingResult.CreatedAt,
		})
}

type TestingResultItem struct {
	TestResult string         `json:"test_result"`
	Owner      sdk.AccAddress `json:"owner"`
	CreatedAt  time.Time      `json:"created_at"`
}

func (d TestingResultItem) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
