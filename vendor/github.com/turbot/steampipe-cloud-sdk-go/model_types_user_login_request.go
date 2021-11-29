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

// TypesUserLoginRequest struct for TypesUserLoginRequest
type TypesUserLoginRequest struct {
	Email string `json:"email"`
}

// NewTypesUserLoginRequest instantiates a new TypesUserLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesUserLoginRequest(email string) *TypesUserLoginRequest {
	this := TypesUserLoginRequest{}
	this.Email = email
	return &this
}

// NewTypesUserLoginRequestWithDefaults instantiates a new TypesUserLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesUserLoginRequestWithDefaults() *TypesUserLoginRequest {
	this := TypesUserLoginRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *TypesUserLoginRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *TypesUserLoginRequest) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *TypesUserLoginRequest) SetEmail(v string) {
	o.Email = v
}

func (o TypesUserLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableTypesUserLoginRequest struct {
	value *TypesUserLoginRequest
	isSet bool
}

func (v NullableTypesUserLoginRequest) Get() *TypesUserLoginRequest {
	return v.value
}

func (v *NullableTypesUserLoginRequest) Set(val *TypesUserLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesUserLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesUserLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesUserLoginRequest(val *TypesUserLoginRequest) *NullableTypesUserLoginRequest {
	return &NullableTypesUserLoginRequest{value: val, isSet: true}
}

func (v NullableTypesUserLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesUserLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


