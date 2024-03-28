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

package validator

import (
	"cosmossdk.io/errors"
)

const (
	Codespace = "validation"
)

var (
	ErrRequiredFieldMissing     = errors.Register(Codespace, 900, "required field missing")
	ErrFieldMaxLengthExceeded   = errors.Register(Codespace, 901, "field max length exceeded")
	ErrFieldNotValid            = errors.Register(Codespace, 902, "field not valid")
	ErrFieldLowerBoundViolated  = errors.Register(Codespace, 903, "field lower bound violated")
	ErrFieldUpperBoundViolated  = errors.Register(Codespace, 904, "field upper bound violated")
	ErrFieldMinLengthNotReached = errors.Register(Codespace, 905, "field min length not reached")
)
