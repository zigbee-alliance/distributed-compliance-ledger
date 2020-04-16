package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type TestingResult struct {
	VID        int16          `json:"vid"`
	PID        int16          `json:"pid"`
	Owner      sdk.AccAddress `json:"owner"`
	TestResult string         `json:"test_result"`
	TestDate   time.Time      `json:"test_date"` // rfc3339 encoded date
}

func NewTestingResult(vid int16, pid int16, owner sdk.AccAddress, testResult string, testDate time.Time) TestingResult {
	return TestingResult{
		VID:        vid,
		PID:        pid,
		Owner:      owner,
		TestResult: testResult,
		TestDate:   testDate,
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
			TestDate:   testingResult.TestDate,
		})
}

type TestingResultItem struct {
	Owner      sdk.AccAddress `json:"owner"`
	TestResult string         `json:"test_result"`
	TestDate   time.Time      `json:"test_date"` // rfc3339 encoded date
}

func (d TestingResultItem) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
