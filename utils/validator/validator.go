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
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni *ut.UniversalTranslator
	vl  *validator.Validate
)

//nolint:wrapcheck,errcheck
func Validate(s interface{}) error {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	vl = validator.New()

	_ = en_translations.RegisterDefaultTranslations(vl, trans)

	_ = vl.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	_ = vl.RegisterTranslation("required_with", trans, func(ut ut.Translator) error {
		return ut.Add("required_with", "{0} is required if {1} is set", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_with", fe.Field(), fe.Param())

		return t
	})

	_ = vl.RegisterTranslation("required_if", trans, func(ut ut.Translator) error {
		return ut.Add("required_if", "{0} is required if {1}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_if", fe.Field(), strings.Replace(fe.Param(), " ", "=", 1))

		return t
	})

	vl.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "{0} must not be less than {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gte", fe.Field(), fe.Param())

		return t
	})

	vl.RegisterTranslation("lte", trans, func(ut ut.Translator) error {
		return ut.Add("lte", "{0} must not be greater than {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("lte", fe.Field(), fe.Param())

		return t
	})

	// Please note that we use `max` tag for fields of `string` type only
	vl.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "maximum length for {0} allowed is {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())

		return t
	})

	// Please note that we use `min` tag for fields of `string` type only
	vl.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "minimum length for {0} allowed is {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())

		return t
	})

	vl.RegisterTranslation("url", trans, func(ut ut.Translator) error {
		return ut.Add("url", "Field {0} : {1} is not a valid url", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("url", fe.Field(), fmt.Sprintf("%v", fe.Value()))

		return t
	})

	vl.RegisterTranslation("startsnotwith", trans, func(ut ut.Translator) error {
		return ut.Add("startsnotwith", "Field {0} : {1} is not a valid url", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("startsnotwith", fe.Field(), fmt.Sprintf("%v", fe.Value()))
		return t
	})

	vl.RegisterTranslation("gtecsfield", trans, func(ut ut.Translator) error {
		return ut.Add("gtecsfield", "{0} must not be less than {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gtecsfield", fe.Field(), fe.Param())
		return t
	})

	//nolint:nestif
	if errs := vl.Struct(s); errs != nil {
		//nolint:errorlint
		for _, e := range errs.(validator.ValidationErrors) {
			if e.Tag() == "required" || e.Tag() == "required_with" || e.Tag() == "required_if" {
				return sdkerrors.Wrap(ErrRequiredFieldMissing, e.Translate(trans))
			}

			if e.Tag() == "max" {
				return sdkerrors.Wrap(ErrFieldMaxLengthExceeded, e.Translate(trans))
			}

			if e.Tag() == "min" {
				return sdkerrors.Wrap(ErrFieldMinLengthExceeded, e.Translate(trans))
			}

			if e.Tag() == "url" || e.Tag() == "startsnotwith" || e.Tag() == "gtecsfield" {
				return sdkerrors.Wrap(ErrFieldNotValid, e.Translate(trans))
			}

			if e.Tag() == "gte" {
				return sdkerrors.Wrap(ErrFieldLowerBoundViolated, e.Translate(trans))
			}

			if e.Tag() == "lte" {
				return sdkerrors.Wrap(ErrFieldUpperBoundViolated, e.Translate(trans))
			}
		}
	}

	return nil
}
