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

// LifecycleStatus the model 'LifecycleStatus'
type LifecycleStatus string

// List of LifecycleStatus
const (
	LIFECYCLESTATUS_ACTIVE   LifecycleStatus = "ACTIVE"
	LIFECYCLESTATUS_INACTIVE LifecycleStatus = "INACTIVE"
)

// All allowed values of LifecycleStatus enum
var AllowedLifecycleStatusEnumValues = []LifecycleStatus{
	"ACTIVE",
	"INACTIVE",
}

func (v *LifecycleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LifecycleStatus(value)
	for _, existing := range AllowedLifecycleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LifecycleStatus", value)
}

// NewLifecycleStatusFromValue returns a pointer to a valid LifecycleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLifecycleStatusFromValue(v string) (*LifecycleStatus, error) {
	ev := LifecycleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LifecycleStatus: valid values are %v", v, AllowedLifecycleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LifecycleStatus) IsValid() bool {
	for _, existing := range AllowedLifecycleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LifecycleStatus value
func (v LifecycleStatus) Ptr() *LifecycleStatus {
	return &v
}

type NullableLifecycleStatus struct {
	value *LifecycleStatus
	isSet bool
}

func (v NullableLifecycleStatus) Get() *LifecycleStatus {
	return v.value
}

func (v *NullableLifecycleStatus) Set(val *LifecycleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleStatus(val *LifecycleStatus) *NullableLifecycleStatus {
	return &NullableLifecycleStatus{value: val, isSet: true}
}

func (v NullableLifecycleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
