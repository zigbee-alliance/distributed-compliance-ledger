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
	"github.com/go-playground/validator/v10"
)

func requiredIfBit0Set(fl validator.FieldLevel) bool {
	// Retrieve the tag parameter specifying the parent field name
	param := fl.Param()
	if param == "" {
		return true // If no parameter is provided, default to not requiring the field
	}

	// Use reflection to access the specified parent field
	parentField := fl.Top().FieldByName(param)
	if !parentField.IsZero() {
		return false // Return false if the parent field is not found
	}

	// Check if the parent field's 0th bit is set
	parentValue := parentField.Interface()
	if parentValueInt, ok := parentValue.(int32); ok && (parentValueInt&1 == 1) {
		return fl.Field().Len() > 0 // Field must not be empty if the bit is set
	}

	return true // Field is not required if the condition is not met
}
