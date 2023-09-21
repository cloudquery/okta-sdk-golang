/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// OtpTotpEncoding the model 'OtpTotpEncoding'
type OtpTotpEncoding string

// List of OtpTotpEncoding
const (
	OTPTOTPENCODING_BASE32      OtpTotpEncoding = "base32"
	OTPTOTPENCODING_BASE64      OtpTotpEncoding = "base64"
	OTPTOTPENCODING_HEXADECIMAL OtpTotpEncoding = "hexadecimal"
)

// All allowed values of OtpTotpEncoding enum
var AllowedOtpTotpEncodingEnumValues = []OtpTotpEncoding{
	"base32",
	"base64",
	"hexadecimal",
}

func (v *OtpTotpEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OtpTotpEncoding(value)
	for _, existing := range AllowedOtpTotpEncodingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OtpTotpEncoding", value)
}

// NewOtpTotpEncodingFromValue returns a pointer to a valid OtpTotpEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOtpTotpEncodingFromValue(v string) (*OtpTotpEncoding, error) {
	ev := OtpTotpEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OtpTotpEncoding: valid values are %v", v, AllowedOtpTotpEncodingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OtpTotpEncoding) IsValid() bool {
	for _, existing := range AllowedOtpTotpEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OtpTotpEncoding value
func (v OtpTotpEncoding) Ptr() *OtpTotpEncoding {
	return &v
}

type NullableOtpTotpEncoding struct {
	value *OtpTotpEncoding
	isSet bool
}

func (v NullableOtpTotpEncoding) Get() *OtpTotpEncoding {
	return v.value
}

func (v *NullableOtpTotpEncoding) Set(val *OtpTotpEncoding) {
	v.value = val
	v.isSet = true
}

func (v NullableOtpTotpEncoding) IsSet() bool {
	return v.isSet
}

func (v *NullableOtpTotpEncoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtpTotpEncoding(val *OtpTotpEncoding) *NullableOtpTotpEncoding {
	return &NullableOtpTotpEncoding{value: val, isSet: true}
}

func (v NullableOtpTotpEncoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtpTotpEncoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
