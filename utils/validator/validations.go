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
		parentStruct = parentStruct.Elem() // Dereference the pointer
	}

	parentField := parentStruct.FieldByName(param)
	if !parentField.IsValid() {
		return false
	}

	parentValue := parentField.Interface()
	if parentValueInt, ok := parentValue.(int32); ok && (parentValueInt&1 == 1) {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			return field.Len() > 0
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return field.Int() != 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return field.Uint() != 0
		default:
			return false
		}
	}

	return true
}
