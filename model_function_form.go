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

// checks if the FunctionForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionForm{}

// FunctionForm struct for FunctionForm
type FunctionForm struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Content string `json:"content"`
	Meta FunctionMeta `json:"meta"`
}

type _FunctionForm FunctionForm

// NewFunctionForm instantiates a new FunctionForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionForm(id string, name string, content string, meta FunctionMeta) *FunctionForm {
	this := FunctionForm{}
	this.Id = id
	this.Name = name
	this.Content = content
	this.Meta = meta
	return &this
}

// NewFunctionFormWithDefaults instantiates a new FunctionForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionFormWithDefaults() *FunctionForm {
	this := FunctionForm{}
	return &this
}

// GetId returns the Id field value
func (o *FunctionForm) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FunctionForm) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FunctionForm) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FunctionForm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FunctionForm) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FunctionForm) SetName(v string) {
	o.Name = v
}

// GetContent returns the Content field value
func (o *FunctionForm) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *FunctionForm) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *FunctionForm) SetContent(v string) {
	o.Content = v
}

// GetMeta returns the Meta field value
func (o *FunctionForm) GetMeta() FunctionMeta {
	if o == nil {
		var ret FunctionMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FunctionForm) GetMetaOk() (*FunctionMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FunctionForm) SetMeta(v FunctionMeta) {
	o.Meta = v
}

func (o FunctionForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["content"] = o.Content
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *FunctionForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"content",
		"meta",
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

	varFunctionForm := _FunctionForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFunctionForm)

	if err != nil {
		return err
	}

	*o = FunctionForm(varFunctionForm)

	return err
}

type NullableFunctionForm struct {
	value *FunctionForm
	isSet bool
}

func (v NullableFunctionForm) Get() *FunctionForm {
	return v.value
}

func (v *NullableFunctionForm) Set(val *FunctionForm) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionForm) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionForm(val *FunctionForm) *NullableFunctionForm {
	return &NullableFunctionForm{value: val, isSet: true}
}

func (v NullableFunctionForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


