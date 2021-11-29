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

// TypesUpdateConnectionRequest struct for TypesUpdateConnectionRequest
type TypesUpdateConnectionRequest struct {
	Config *map[string]interface{} `json:"config,omitempty"`
	Handle *string `json:"handle,omitempty"`
}

// NewTypesUpdateConnectionRequest instantiates a new TypesUpdateConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesUpdateConnectionRequest() *TypesUpdateConnectionRequest {
	this := TypesUpdateConnectionRequest{}
	return &this
}

// NewTypesUpdateConnectionRequestWithDefaults instantiates a new TypesUpdateConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesUpdateConnectionRequestWithDefaults() *TypesUpdateConnectionRequest {
	this := TypesUpdateConnectionRequest{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *TypesUpdateConnectionRequest) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUpdateConnectionRequest) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *TypesUpdateConnectionRequest) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *TypesUpdateConnectionRequest) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *TypesUpdateConnectionRequest) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUpdateConnectionRequest) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *TypesUpdateConnectionRequest) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *TypesUpdateConnectionRequest) SetHandle(v string) {
	o.Handle = &v
}

func (o TypesUpdateConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	return json.Marshal(toSerialize)
}

type NullableTypesUpdateConnectionRequest struct {
	value *TypesUpdateConnectionRequest
	isSet bool
}

func (v NullableTypesUpdateConnectionRequest) Get() *TypesUpdateConnectionRequest {
	return v.value
}

func (v *NullableTypesUpdateConnectionRequest) Set(val *TypesUpdateConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesUpdateConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesUpdateConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesUpdateConnectionRequest(val *TypesUpdateConnectionRequest) *NullableTypesUpdateConnectionRequest {
	return &NullableTypesUpdateConnectionRequest{value: val, isSet: true}
}

func (v NullableTypesUpdateConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesUpdateConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


