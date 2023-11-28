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

// ProvisioningConnectionRequest struct for ProvisioningConnectionRequest
type ProvisioningConnectionRequest struct {
	Profile              ProvisioningConnectionProfile `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionRequest ProvisioningConnectionRequest

// NewProvisioningConnectionRequest instantiates a new ProvisioningConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionRequest(profile ProvisioningConnectionProfile) *ProvisioningConnectionRequest {
	this := ProvisioningConnectionRequest{}
	this.Profile = profile
	return &this
}

// NewProvisioningConnectionRequestWithDefaults instantiates a new ProvisioningConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionRequestWithDefaults() *ProvisioningConnectionRequest {
	this := ProvisioningConnectionRequest{}
	return &this
}

// GetProfile returns the Profile field value
func (o *ProvisioningConnectionRequest) GetProfile() ProvisioningConnectionProfile {
	if o == nil {
		var ret ProvisioningConnectionProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionRequest) GetProfileOk() (*ProvisioningConnectionProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ProvisioningConnectionRequest) SetProfile(v ProvisioningConnectionProfile) {
	o.Profile = v
}

func (o ProvisioningConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnectionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnectionRequest := _ProvisioningConnectionRequest{}

	err = json.Unmarshal(bytes, &varProvisioningConnectionRequest)
	if err == nil {
		*o = ProvisioningConnectionRequest(varProvisioningConnectionRequest)
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

type NullableProvisioningConnectionRequest struct {
	value *ProvisioningConnectionRequest
	isSet bool
}

func (v NullableProvisioningConnectionRequest) Get() *ProvisioningConnectionRequest {
	return v.value
}

func (v *NullableProvisioningConnectionRequest) Set(val *ProvisioningConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionRequest(val *ProvisioningConnectionRequest) *NullableProvisioningConnectionRequest {
	return &NullableProvisioningConnectionRequest{value: val, isSet: true}
}

func (v NullableProvisioningConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
