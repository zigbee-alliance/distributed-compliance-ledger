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

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni *ut.UniversalTranslator
	vl  *validator.Validate
)

func validate(s interface{}, performAddValidation bool) sdk.Error {

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	vl = validator.New()

	vl.RegisterValidation("requiredForAdd", onlyRequiredForAdd)
	en_translations.RegisterDefaultTranslations(vl, trans)

	_ = vl.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = vl.RegisterTranslation("requiredForAdd", trans, func(ut ut.Translator) error {
		return ut.Add("requiredForAdd", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("requiredForAdd", fe.Field())
		return t
	})

	_ = vl.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "maximum length for {0} allowed is {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())
		return t
	})

	_ = vl.RegisterTranslation("url", trans, func(ut ut.Translator) error {
		return ut.Add("url", "Field {0} : {1} is not a valid url", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("url", fe.Field(), fmt.Sprintf("%v", fe.Value()))
		return t
	})

	errs := vl.Struct(s)

	if errs != nil {
		for _, e := range errs.(validator.ValidationErrors) {
			if e.Tag() == "max" {
				return sdk.NewError(Codespace, CodeFieldMaxLengthExceeded, e.Translate(trans))
			}
			if e.Tag() == "required" {
				return sdk.NewError(Codespace, CodeRequiredFieldMissing, e.Translate(trans))
			}
			if e.Tag() == "requiredForAdd" && performAddValidation {
				return sdk.NewError(Codespace, CodeRequiredFieldMissing, e.Translate(trans))
			}
			if e.Tag() == "url" {
				return sdk.NewError(Codespace, CodeFieldNotValid, e.Translate(trans))
			}

		}
	}

	return nil
}

func ValidateUpdate(s interface{}) sdk.Error {
	return validate(s, false)
}

func ValidateAdd(s interface{}) sdk.Error {
	return validate(s, true)

}

func onlyRequiredForAdd(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return false
	} else {
		return true
	}
}
