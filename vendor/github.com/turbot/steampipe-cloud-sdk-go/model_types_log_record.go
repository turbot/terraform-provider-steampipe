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

// TypesLogRecord struct for TypesLogRecord
type TypesLogRecord struct {
	ActorAvatarUrl string `json:"actor_avatar_url"`
	ActorDisplayName string `json:"actor_display_name"`
	ActorHandle string `json:"actor_handle"`
	ActorId string `json:"actor_id"`
	CreatedAt TypesJSONTime `json:"created_at"`
	Duration *int32 `json:"duration,omitempty"`
	Id string `json:"id"`
	LogTimestamp *TypesJSONTime `json:"log_timestamp,omitempty"`
	Query *string `json:"query,omitempty"`
	WorkspaceHandle string `json:"workspace_handle"`
	WorkspaceId string `json:"workspace_id"`
}

// NewTypesLogRecord instantiates a new TypesLogRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesLogRecord(actorAvatarUrl string, actorDisplayName string, actorHandle string, actorId string, createdAt TypesJSONTime, id string, workspaceHandle string, workspaceId string) *TypesLogRecord {
	this := TypesLogRecord{}
	this.ActorAvatarUrl = actorAvatarUrl
	this.ActorDisplayName = actorDisplayName
	this.ActorHandle = actorHandle
	this.ActorId = actorId
	this.CreatedAt = createdAt
	this.Id = id
	this.WorkspaceHandle = workspaceHandle
	this.WorkspaceId = workspaceId
	return &this
}

// NewTypesLogRecordWithDefaults instantiates a new TypesLogRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesLogRecordWithDefaults() *TypesLogRecord {
	this := TypesLogRecord{}
	return &this
}

// GetActorAvatarUrl returns the ActorAvatarUrl field value
func (o *TypesLogRecord) GetActorAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActorAvatarUrl
}

// GetActorAvatarUrlOk returns a tuple with the ActorAvatarUrl field value
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetActorAvatarUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ActorAvatarUrl, true
}

// SetActorAvatarUrl sets field value
func (o *TypesLogRecord) SetActorAvatarUrl(v string) {
	o.ActorAvatarUrl = v
}

// GetActorDisplayName returns the ActorDisplayName field value
func (o *TypesLogRecord) GetActorDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActorDisplayName
}

// GetActorDisplayNameOk returns a tuple with the ActorDisplayName field value
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetActorDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ActorDisplayName, true
}

// SetActorDisplayName sets field value
func (o *TypesLogRecord) SetActorDisplayName(v string) {
	o.ActorDisplayName = v
}

// GetActorHandle returns the ActorHandle field value
func (o *TypesLogRecord) GetActorHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActorHandle
}

// GetActorHandleOk returns a tuple with the ActorHandle field value
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetActorHandleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ActorHandle, true
}

// SetActorHandle sets field value
func (o *TypesLogRecord) SetActorHandle(v string) {
	o.ActorHandle = v
}

// GetActorId returns the ActorId field value
func (o *TypesLogRecord) GetActorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActorId
}

// GetActorIdOk returns a tuple with the ActorId field value
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetActorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ActorId, true
}

// SetActorId sets field value
func (o *TypesLogRecord) SetActorId(v string) {
	o.ActorId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TypesLogRecord) GetCreatedAt() TypesJSONTime {
	if o == nil {
		var ret TypesJSONTime
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetCreatedAtOk() (*TypesJSONTime, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TypesLogRecord) SetCreatedAt(v TypesJSONTime) {
	o.CreatedAt = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *TypesLogRecord) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *TypesLogRecord) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *TypesLogRecord) SetDuration(v int32) {
	o.Duration = &v
}

// GetId returns the Id field value
func (o *TypesLogRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TypesLogRecord) SetId(v string) {
	o.Id = v
}

// GetLogTimestamp returns the LogTimestamp field value if set, zero value otherwise.
func (o *TypesLogRecord) GetLogTimestamp() TypesJSONTime {
	if o == nil || o.LogTimestamp == nil {
		var ret TypesJSONTime
		return ret
	}
	return *o.LogTimestamp
}

// GetLogTimestampOk returns a tuple with the LogTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetLogTimestampOk() (*TypesJSONTime, bool) {
	if o == nil || o.LogTimestamp == nil {
		return nil, false
	}
	return o.LogTimestamp, true
}

// HasLogTimestamp returns a boolean if a field has been set.
func (o *TypesLogRecord) HasLogTimestamp() bool {
	if o != nil && o.LogTimestamp != nil {
		return true
	}

	return false
}

// SetLogTimestamp gets a reference to the given TypesJSONTime and assigns it to the LogTimestamp field.
func (o *TypesLogRecord) SetLogTimestamp(v TypesJSONTime) {
	o.LogTimestamp = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TypesLogRecord) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TypesLogRecord) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *TypesLogRecord) SetQuery(v string) {
	o.Query = &v
}

// GetWorkspaceHandle returns the WorkspaceHandle field value
func (o *TypesLogRecord) GetWorkspaceHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceHandle
}

// GetWorkspaceHandleOk returns a tuple with the WorkspaceHandle field value
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetWorkspaceHandleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WorkspaceHandle, true
}

// SetWorkspaceHandle sets field value
func (o *TypesLogRecord) SetWorkspaceHandle(v string) {
	o.WorkspaceHandle = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *TypesLogRecord) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *TypesLogRecord) GetWorkspaceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *TypesLogRecord) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

func (o TypesLogRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["actor_avatar_url"] = o.ActorAvatarUrl
	}
	if true {
		toSerialize["actor_display_name"] = o.ActorDisplayName
	}
	if true {
		toSerialize["actor_handle"] = o.ActorHandle
	}
	if true {
		toSerialize["actor_id"] = o.ActorId
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LogTimestamp != nil {
		toSerialize["log_timestamp"] = o.LogTimestamp
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["workspace_handle"] = o.WorkspaceHandle
	}
	if true {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableTypesLogRecord struct {
	value *TypesLogRecord
	isSet bool
}

func (v NullableTypesLogRecord) Get() *TypesLogRecord {
	return v.value
}

func (v *NullableTypesLogRecord) Set(val *TypesLogRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesLogRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesLogRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesLogRecord(val *TypesLogRecord) *NullableTypesLogRecord {
	return &NullableTypesLogRecord{value: val, isSet: true}
}

func (v NullableTypesLogRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesLogRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


