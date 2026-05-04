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
	"context"
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
	return validURL(fl, "http", "https")
}

func isValidHttpsUrl(fl validator.FieldLevel) bool { //nolint:stylecheck
	return validURL(fl, "https")
}

var allowed4XXStatusCodes = []int{
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusUnavailableForLegalReasons,
}

const (
	livenessCheckTimeout = 10 * time.Second
)

var httpClient = &http.Client{Timeout: livenessCheckTimeout}

func validURL(fl validator.FieldLevel, allowedSchemes ...string) bool {
	raw := fl.Field().String()
	if raw == "" {
		return true
	}

	u, err := url.ParseRequestURI(raw)
	if err != nil || u.Host == "" {
		return false
	}

	if !isSchemeAllowed(u.Scheme, allowedSchemes) {
		return false
	}

	return isLiveURL(u)
}

func isSchemeAllowed(scheme string, allowed []string) bool {
	for _, s := range allowed {
		if scheme == s {
			return true
		}
	}

	return false
}

func isLiveURL(u *url.URL) bool {
	if config.DisableURLLivenessCheck {
		return true
	}

	ctx, cancel := context.WithTimeout(context.Background(), livenessCheckTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, u.String(), nil)
	if err != nil {
		return false
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		return true
	}

	for _, code := range allowed4XXStatusCodes {
		if code == resp.StatusCode {
			return true
		}
	}

	return false
}
