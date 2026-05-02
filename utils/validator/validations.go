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
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/zigbee-alliance/distributed-compliance-ledger/internal/config"
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

func isValidHttpOrHttpsUrl(fl validator.FieldLevel) bool { //nolint:stylecheck
	return _validURL(fl, "http", "https")
}

func isValidHttpsUrl(fl validator.FieldLevel) bool { //nolint:stylecheck
	return _validURL(fl, "https")
}

var allowed4XXStatusCodes = []int{
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusUnavailableForLegalReasons,
}
var httpClient = &http.Client{Timeout: 10 * time.Second}

func _validURL(fl validator.FieldLevel, allowedSchemas ...string) bool {
	raw := fl.Field().String()
	// Field is empty, or omitempty is set, skip checks
	if raw == "" {
		return true
	}

	u, err := url.ParseRequestURI(raw)
	if err != nil || u.Host == "" {
		return false
	}

	isSchemaAllowed := false
	for _, schema := range allowedSchemas {
		if u.Scheme == schema {
			isSchemaAllowed = true
			break
		}
	}

	if _isLiveURL(u) || !isSchemaAllowed {
		return isSchemaAllowed
	}

	return false
}

func _isLiveURL(u *url.URL) bool {
	if config.DisableURLLivenessCheck {
		return true
	}

	// HEAD request only retrieves headers, not the body
	resp, err := httpClient.Head(u.String())
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true
	}

	for _, code := range allowed4XXStatusCodes {
		if code == resp.StatusCode {
			return true
		}
	}

	return false
}
