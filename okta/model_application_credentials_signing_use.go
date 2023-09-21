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

// ApplicationCredentialsSigningUse the model 'ApplicationCredentialsSigningUse'
type ApplicationCredentialsSigningUse string

// List of ApplicationCredentialsSigningUse
const (
	APPLICATIONCREDENTIALSSIGNINGUSE_SIG ApplicationCredentialsSigningUse = "sig"
)

// All allowed values of ApplicationCredentialsSigningUse enum
var AllowedApplicationCredentialsSigningUseEnumValues = []ApplicationCredentialsSigningUse{
	"sig",
}

func (v *ApplicationCredentialsSigningUse) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplicationCredentialsSigningUse(value)
	for _, existing := range AllowedApplicationCredentialsSigningUseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplicationCredentialsSigningUse", value)
}

// NewApplicationCredentialsSigningUseFromValue returns a pointer to a valid ApplicationCredentialsSigningUse
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplicationCredentialsSigningUseFromValue(v string) (*ApplicationCredentialsSigningUse, error) {
	ev := ApplicationCredentialsSigningUse(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplicationCredentialsSigningUse: valid values are %v", v, AllowedApplicationCredentialsSigningUseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplicationCredentialsSigningUse) IsValid() bool {
	for _, existing := range AllowedApplicationCredentialsSigningUseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationCredentialsSigningUse value
func (v ApplicationCredentialsSigningUse) Ptr() *ApplicationCredentialsSigningUse {
	return &v
}

type NullableApplicationCredentialsSigningUse struct {
	value *ApplicationCredentialsSigningUse
	isSet bool
}

func (v NullableApplicationCredentialsSigningUse) Get() *ApplicationCredentialsSigningUse {
	return v.value
}

func (v *NullableApplicationCredentialsSigningUse) Set(val *ApplicationCredentialsSigningUse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCredentialsSigningUse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCredentialsSigningUse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCredentialsSigningUse(val *ApplicationCredentialsSigningUse) *NullableApplicationCredentialsSigningUse {
	return &NullableApplicationCredentialsSigningUse{value: val, isSet: true}
}

func (v NullableApplicationCredentialsSigningUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCredentialsSigningUse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
