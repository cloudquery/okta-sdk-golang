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

// FeatureStageValue Current release stage of the feature
type FeatureStageValue string

// List of FeatureStageValue
const (
	FEATURESTAGEVALUE_BETA FeatureStageValue = "BETA"
	FEATURESTAGEVALUE_EA   FeatureStageValue = "EA"
)

// All allowed values of FeatureStageValue enum
var AllowedFeatureStageValueEnumValues = []FeatureStageValue{
	"BETA",
	"EA",
}

func (v *FeatureStageValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeatureStageValue(value)
	for _, existing := range AllowedFeatureStageValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeatureStageValue", value)
}

// NewFeatureStageValueFromValue returns a pointer to a valid FeatureStageValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeatureStageValueFromValue(v string) (*FeatureStageValue, error) {
	ev := FeatureStageValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeatureStageValue: valid values are %v", v, AllowedFeatureStageValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeatureStageValue) IsValid() bool {
	for _, existing := range AllowedFeatureStageValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureStageValue value
func (v FeatureStageValue) Ptr() *FeatureStageValue {
	return &v
}

type NullableFeatureStageValue struct {
	value *FeatureStageValue
	isSet bool
}

func (v NullableFeatureStageValue) Get() *FeatureStageValue {
	return v.value
}

func (v *NullableFeatureStageValue) Set(val *FeatureStageValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureStageValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureStageValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureStageValue(val *FeatureStageValue) *NullableFeatureStageValue {
	return &NullableFeatureStageValue{value: val, isSet: true}
}

func (v NullableFeatureStageValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureStageValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
