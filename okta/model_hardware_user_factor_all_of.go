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
)

// HardwareUserFactorAllOf struct for HardwareUserFactorAllOf
type HardwareUserFactorAllOf struct {
	Profile              *HardwareUserFactorProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HardwareUserFactorAllOf HardwareUserFactorAllOf

// NewHardwareUserFactorAllOf instantiates a new HardwareUserFactorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHardwareUserFactorAllOf() *HardwareUserFactorAllOf {
	this := HardwareUserFactorAllOf{}
	return &this
}

// NewHardwareUserFactorAllOfWithDefaults instantiates a new HardwareUserFactorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHardwareUserFactorAllOfWithDefaults() *HardwareUserFactorAllOf {
	this := HardwareUserFactorAllOf{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *HardwareUserFactorAllOf) GetProfile() HardwareUserFactorProfile {
	if o == nil || o.Profile == nil {
		var ret HardwareUserFactorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareUserFactorAllOf) GetProfileOk() (*HardwareUserFactorProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *HardwareUserFactorAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given HardwareUserFactorProfile and assigns it to the Profile field.
func (o *HardwareUserFactorAllOf) SetProfile(v HardwareUserFactorProfile) {
	o.Profile = &v
}

func (o HardwareUserFactorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HardwareUserFactorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHardwareUserFactorAllOf := _HardwareUserFactorAllOf{}

	err = json.Unmarshal(bytes, &varHardwareUserFactorAllOf)
	if err == nil {
		*o = HardwareUserFactorAllOf(varHardwareUserFactorAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableHardwareUserFactorAllOf struct {
	value *HardwareUserFactorAllOf
	isSet bool
}

func (v NullableHardwareUserFactorAllOf) Get() *HardwareUserFactorAllOf {
	return v.value
}

func (v *NullableHardwareUserFactorAllOf) Set(val *HardwareUserFactorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHardwareUserFactorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHardwareUserFactorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHardwareUserFactorAllOf(val *HardwareUserFactorAllOf) *NullableHardwareUserFactorAllOf {
	return &NullableHardwareUserFactorAllOf{value: val, isSet: true}
}

func (v NullableHardwareUserFactorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHardwareUserFactorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
