/*
Steampipe Cloud API

Steampipe Cloud API

API version: 1.0
Contact: support@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TypesUpdateTokenRequest struct for TypesUpdateTokenRequest
type TypesUpdateTokenRequest struct {
	Status string `json:"status"`
}

// NewTypesUpdateTokenRequest instantiates a new TypesUpdateTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesUpdateTokenRequest(status string) *TypesUpdateTokenRequest {
	this := TypesUpdateTokenRequest{}
	this.Status = status
	return &this
}

// NewTypesUpdateTokenRequestWithDefaults instantiates a new TypesUpdateTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesUpdateTokenRequestWithDefaults() *TypesUpdateTokenRequest {
	this := TypesUpdateTokenRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *TypesUpdateTokenRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TypesUpdateTokenRequest) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TypesUpdateTokenRequest) SetStatus(v string) {
	o.Status = v
}

func (o TypesUpdateTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableTypesUpdateTokenRequest struct {
	value *TypesUpdateTokenRequest
	isSet bool
}

func (v NullableTypesUpdateTokenRequest) Get() *TypesUpdateTokenRequest {
	return v.value
}

func (v *NullableTypesUpdateTokenRequest) Set(val *TypesUpdateTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesUpdateTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesUpdateTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesUpdateTokenRequest(val *TypesUpdateTokenRequest) *NullableTypesUpdateTokenRequest {
	return &NullableTypesUpdateTokenRequest{value: val, isSet: true}
}

func (v NullableTypesUpdateTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesUpdateTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


