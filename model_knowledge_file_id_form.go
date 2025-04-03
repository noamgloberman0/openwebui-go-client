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

// checks if the KnowledgeFileIdForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KnowledgeFileIdForm{}

// KnowledgeFileIdForm struct for KnowledgeFileIdForm
type KnowledgeFileIdForm struct {
	FileId string `json:"file_id"`
}

type _KnowledgeFileIdForm KnowledgeFileIdForm

// NewKnowledgeFileIdForm instantiates a new KnowledgeFileIdForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKnowledgeFileIdForm(fileId string) *KnowledgeFileIdForm {
	this := KnowledgeFileIdForm{}
	this.FileId = fileId
	return &this
}

// NewKnowledgeFileIdFormWithDefaults instantiates a new KnowledgeFileIdForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKnowledgeFileIdFormWithDefaults() *KnowledgeFileIdForm {
	this := KnowledgeFileIdForm{}
	return &this
}

// GetFileId returns the FileId field value
func (o *KnowledgeFileIdForm) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *KnowledgeFileIdForm) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *KnowledgeFileIdForm) SetFileId(v string) {
	o.FileId = v
}

func (o KnowledgeFileIdForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KnowledgeFileIdForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_id"] = o.FileId
	return toSerialize, nil
}

func (o *KnowledgeFileIdForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file_id",
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

	varKnowledgeFileIdForm := _KnowledgeFileIdForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKnowledgeFileIdForm)

	if err != nil {
		return err
	}

	*o = KnowledgeFileIdForm(varKnowledgeFileIdForm)

	return err
}

type NullableKnowledgeFileIdForm struct {
	value *KnowledgeFileIdForm
	isSet bool
}

func (v NullableKnowledgeFileIdForm) Get() *KnowledgeFileIdForm {
	return v.value
}

func (v *NullableKnowledgeFileIdForm) Set(val *KnowledgeFileIdForm) {
	v.value = val
	v.isSet = true
}

func (v NullableKnowledgeFileIdForm) IsSet() bool {
	return v.isSet
}

func (v *NullableKnowledgeFileIdForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKnowledgeFileIdForm(val *KnowledgeFileIdForm) *NullableKnowledgeFileIdForm {
	return &NullableKnowledgeFileIdForm{value: val, isSet: true}
}

func (v NullableKnowledgeFileIdForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKnowledgeFileIdForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


