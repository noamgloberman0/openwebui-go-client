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

// checks if the TagModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagModel{}

// TagModel struct for TagModel
type TagModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
	UserId string `json:"user_id"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

type _TagModel TagModel

// NewTagModel instantiates a new TagModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagModel(id string, name string, userId string) *TagModel {
	this := TagModel{}
	this.Id = id
	this.Name = name
	this.UserId = userId
	return &this
}

// NewTagModelWithDefaults instantiates a new TagModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagModelWithDefaults() *TagModel {
	this := TagModel{}
	return &this
}

// GetId returns the Id field value
func (o *TagModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TagModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagModel) SetName(v string) {
	o.Name = v
}

// GetUserId returns the UserId field value
func (o *TagModel) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TagModel) SetUserId(v string) {
	o.UserId = v
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagModel) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagModel) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TagModel) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *TagModel) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o TagModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["user_id"] = o.UserId
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *TagModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"user_id",
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

	varTagModel := _TagModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagModel)

	if err != nil {
		return err
	}

	*o = TagModel(varTagModel)

	return err
}

type NullableTagModel struct {
	value *TagModel
	isSet bool
}

func (v NullableTagModel) Get() *TagModel {
	return v.value
}

func (v *NullableTagModel) Set(val *TagModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagModel(val *TagModel) *NullableTagModel {
	return &NullableTagModel{value: val, isSet: true}
}

func (v NullableTagModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


