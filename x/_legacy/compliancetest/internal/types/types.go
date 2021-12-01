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

package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TestingResult struct {
	VID                   uint16         `json:"vid"`
	PID                   uint16         `json:"pid"`
	SoftwareVersion       uint32         `json:"softwareVersion"`
	SoftwareVersionString string         `json:"softwareVersionString,omitempty"`
	Owner                 sdk.AccAddress `json:"owner"`
	TestResult            string         `json:"test_result"`
	TestDate              time.Time      `json:"test_date"` // rfc3339 encoded date
}

func NewTestingResult(vid uint16, pid uint16,
	softwareVersion uint32, softwareVersionString string,
	owner sdk.AccAddress,
	testResult string, testDate time.Time) TestingResult {
	return TestingResult{
		VID:                   vid,
		PID:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		Owner:                 owner,
		TestResult:            testResult,
		TestDate:              testDate,
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
	VID                   uint16              `json:"vid"`
	PID                   uint16              `json:"pid"`
	SoftwareVersion       uint32              `json:"softwareVersion"`
	SoftwareVersionString string              `json:"softwareVersionString,omitempty"`
	Results               []TestingResultItem `json:"results"`
}

func NewTestingResults(vid uint16, pid uint16,
	softwareVersion uint32, softwareVersionString string) TestingResults {
	return TestingResults{
		VID:                   vid,
		PID:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		Results:               []TestingResultItem{},
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
