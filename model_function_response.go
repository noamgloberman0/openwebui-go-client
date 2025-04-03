/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FunctionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionResponse{}

// FunctionResponse struct for FunctionResponse
type FunctionResponse struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Meta FunctionMeta `json:"meta"`
	IsActive bool `json:"is_active"`
	IsGlobal bool `json:"is_global"`
	UpdatedAt int32 `json:"updated_at"`
	CreatedAt int32 `json:"created_at"`
}

type _FunctionResponse FunctionResponse

// NewFunctionResponse instantiates a new FunctionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionResponse(id string, userId string, type_ string, name string, meta FunctionMeta, isActive bool, isGlobal bool, updatedAt int32, createdAt int32) *FunctionResponse {
	this := FunctionResponse{}
	this.Id = id
	this.UserId = userId
	this.Type = type_
	this.Name = name
	this.Meta = meta
	this.IsActive = isActive
	this.IsGlobal = isGlobal
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewFunctionResponseWithDefaults instantiates a new FunctionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionResponseWithDefaults() *FunctionResponse {
	this := FunctionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *FunctionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FunctionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FunctionResponse) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *FunctionResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FunctionResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FunctionResponse) SetUserId(v string) {
	o.UserId = v
}

// GetType returns the Type field value
func (o *FunctionResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FunctionResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FunctionResponse) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *FunctionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FunctionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FunctionResponse) SetName(v string) {
	o.Name = v
}

// GetMeta returns the Meta field value
func (o *FunctionResponse) GetMeta() FunctionMeta {
	if o == nil {
		var ret FunctionMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FunctionResponse) GetMetaOk() (*FunctionMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FunctionResponse) SetMeta(v FunctionMeta) {
	o.Meta = v
}

// GetIsActive returns the IsActive field value
func (o *FunctionResponse) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *FunctionResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *FunctionResponse) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsGlobal returns the IsGlobal field value
func (o *FunctionResponse) GetIsGlobal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value
// and a boolean to check if the value has been set.
func (o *FunctionResponse) GetIsGlobalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGlobal, true
}

// SetIsGlobal sets field value
func (o *FunctionResponse) SetIsGlobal(v bool) {
	o.IsGlobal = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *FunctionResponse) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *FunctionResponse) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *FunctionResponse) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FunctionResponse) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FunctionResponse) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FunctionResponse) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

func (o FunctionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["meta"] = o.Meta
	toSerialize["is_active"] = o.IsActive
	toSerialize["is_global"] = o.IsGlobal
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *FunctionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"type",
		"name",
		"meta",
		"is_active",
		"is_global",
		"updated_at",
		"created_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFunctionResponse := _FunctionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFunctionResponse)

	if err != nil {
		return err
	}

	*o = FunctionResponse(varFunctionResponse)

	return err
}

type NullableFunctionResponse struct {
	value *FunctionResponse
	isSet bool
}

func (v NullableFunctionResponse) Get() *FunctionResponse {
	return v.value
}

func (v *NullableFunctionResponse) Set(val *FunctionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionResponse(val *FunctionResponse) *NullableFunctionResponse {
	return &NullableFunctionResponse{value: val, isSet: true}
}

func (v NullableFunctionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


