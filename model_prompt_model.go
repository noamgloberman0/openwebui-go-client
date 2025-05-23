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

// checks if the PromptModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptModel{}

// PromptModel struct for PromptModel
type PromptModel struct {
	Command string `json:"command"`
	UserId string `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Timestamp int32 `json:"timestamp"`
	AccessControl map[string]interface{} `json:"access_control,omitempty"`
}

type _PromptModel PromptModel

// NewPromptModel instantiates a new PromptModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptModel(command string, userId string, title string, content string, timestamp int32) *PromptModel {
	this := PromptModel{}
	this.Command = command
	this.UserId = userId
	this.Title = title
	this.Content = content
	this.Timestamp = timestamp
	return &this
}

// NewPromptModelWithDefaults instantiates a new PromptModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptModelWithDefaults() *PromptModel {
	this := PromptModel{}
	return &this
}

// GetCommand returns the Command field value
func (o *PromptModel) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *PromptModel) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *PromptModel) SetCommand(v string) {
	o.Command = v
}

// GetUserId returns the UserId field value
func (o *PromptModel) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PromptModel) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PromptModel) SetUserId(v string) {
	o.UserId = v
}

// GetTitle returns the Title field value
func (o *PromptModel) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PromptModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PromptModel) SetTitle(v string) {
	o.Title = v
}

// GetContent returns the Content field value
func (o *PromptModel) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *PromptModel) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *PromptModel) SetContent(v string) {
	o.Content = v
}

// GetTimestamp returns the Timestamp field value
func (o *PromptModel) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *PromptModel) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *PromptModel) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PromptModel) GetAccessControl() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PromptModel) GetAccessControlOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return map[string]interface{}{}, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *PromptModel) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given map[string]interface{} and assigns it to the AccessControl field.
func (o *PromptModel) SetAccessControl(v map[string]interface{}) {
	o.AccessControl = v
}

func (o PromptModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["command"] = o.Command
	toSerialize["user_id"] = o.UserId
	toSerialize["title"] = o.Title
	toSerialize["content"] = o.Content
	toSerialize["timestamp"] = o.Timestamp
	if o.AccessControl != nil {
		toSerialize["access_control"] = o.AccessControl
	}
	return toSerialize, nil
}

func (o *PromptModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"command",
		"user_id",
		"title",
		"content",
		"timestamp",
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

	varPromptModel := _PromptModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPromptModel)

	if err != nil {
		return err
	}

	*o = PromptModel(varPromptModel)

	return err
}

type NullablePromptModel struct {
	value *PromptModel
	isSet bool
}

func (v NullablePromptModel) Get() *PromptModel {
	return v.value
}

func (v *NullablePromptModel) Set(val *PromptModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptModel(val *PromptModel) *NullablePromptModel {
	return &NullablePromptModel{value: val, isSet: true}
}

func (v NullablePromptModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


