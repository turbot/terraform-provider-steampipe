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

// TypesCreateWorkspaceRequest struct for TypesCreateWorkspaceRequest
type TypesCreateWorkspaceRequest struct {
	Handle string `json:"handle"`
}

// NewTypesCreateWorkspaceRequest instantiates a new TypesCreateWorkspaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesCreateWorkspaceRequest(handle string) *TypesCreateWorkspaceRequest {
	this := TypesCreateWorkspaceRequest{}
	this.Handle = handle
	return &this
}

// NewTypesCreateWorkspaceRequestWithDefaults instantiates a new TypesCreateWorkspaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesCreateWorkspaceRequestWithDefaults() *TypesCreateWorkspaceRequest {
	this := TypesCreateWorkspaceRequest{}
	return &this
}

// GetHandle returns the Handle field value
func (o *TypesCreateWorkspaceRequest) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TypesCreateWorkspaceRequest) GetHandleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *TypesCreateWorkspaceRequest) SetHandle(v string) {
	o.Handle = v
}

func (o TypesCreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handle"] = o.Handle
	}
	return json.Marshal(toSerialize)
}

type NullableTypesCreateWorkspaceRequest struct {
	value *TypesCreateWorkspaceRequest
	isSet bool
}

func (v NullableTypesCreateWorkspaceRequest) Get() *TypesCreateWorkspaceRequest {
	return v.value
}

func (v *NullableTypesCreateWorkspaceRequest) Set(val *TypesCreateWorkspaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesCreateWorkspaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesCreateWorkspaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesCreateWorkspaceRequest(val *TypesCreateWorkspaceRequest) *NullableTypesCreateWorkspaceRequest {
	return &NullableTypesCreateWorkspaceRequest{value: val, isSet: true}
}

func (v NullableTypesCreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesCreateWorkspaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


