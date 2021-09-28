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
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeTestingResultsDoNotExist sdk.CodeType = 201
)

func ErrTestingResultDoesNotExist(vid interface{}, pid interface{}, softwareVersion uint32) sdk.Error {
	return sdk.NewError(Codespace, CodeTestingResultsDoNotExist,
		fmt.Sprintf("No testing results about the model with vid=%v pid=%v and softwareVersion=%v on the ledger", vid, pid, softwareVersion))
}
