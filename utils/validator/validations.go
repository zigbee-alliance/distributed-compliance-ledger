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
	"net/url"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func requiredIfBit0Set(fl validator.FieldLevel) bool {
	param := fl.Param()
	if param == "" {
		return true
	}

	parentStruct := fl.Top()
	if parentStruct.Kind() == reflect.Ptr {
		parentStruct = parentStruct.Elem()
	}

	parentField := parentStruct.FieldByName(param)
	if !parentField.IsValid() {
		return false
	}

	parentValue := parentField.Interface()
	if parentValueInt, ok := parentValue.(int32); ok {
		field := fl.Field()

		// field must be specified IF AND ONLY IF parentValueInt&1 == 1
		return (parentValueInt&1 == 1) != field.IsZero()
	}

	return true
}

func isValidHttpOrHttpsUrl(fl validator.FieldLevel) bool {
	return _validURL(fl, "http", "https")
}

func isValidHttpsUrl(fl validator.FieldLevel) bool {
	return _validURL(fl, "https")
}

func _validURL(fl validator.FieldLevel, allowedSchemas ...string) bool {
	raw := fl.Field().String()
	// Field is empty or omitempty is set, skip checks
	if raw == "" {
		return true
	}

	u, _ := url.Parse(raw)
	if u.Host == "" {
		return false
	}

	for _, schema := range allowedSchemas {
		if u.Scheme == schema {
			return true
		}
	}

	return false
}
