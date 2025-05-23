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

// checks if the GenerateEmbeddingsForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateEmbeddingsForm{}

// GenerateEmbeddingsForm struct for GenerateEmbeddingsForm
type GenerateEmbeddingsForm struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	Options map[string]interface{} `json:"options,omitempty"`
	KeepAlive NullableKeepAlive `json:"keep_alive,omitempty"`
}

type _GenerateEmbeddingsForm GenerateEmbeddingsForm

// NewGenerateEmbeddingsForm instantiates a new GenerateEmbeddingsForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateEmbeddingsForm(model string, prompt string) *GenerateEmbeddingsForm {
	this := GenerateEmbeddingsForm{}
	this.Model = model
	this.Prompt = prompt
	return &this
}

// NewGenerateEmbeddingsFormWithDefaults instantiates a new GenerateEmbeddingsForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateEmbeddingsFormWithDefaults() *GenerateEmbeddingsForm {
	this := GenerateEmbeddingsForm{}
	return &this
}

// GetModel returns the Model field value
func (o *GenerateEmbeddingsForm) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *GenerateEmbeddingsForm) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *GenerateEmbeddingsForm) SetModel(v string) {
	o.Model = v
}

// GetPrompt returns the Prompt field value
func (o *GenerateEmbeddingsForm) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *GenerateEmbeddingsForm) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *GenerateEmbeddingsForm) SetPrompt(v string) {
	o.Prompt = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateEmbeddingsForm) GetOptions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateEmbeddingsForm) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *GenerateEmbeddingsForm) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *GenerateEmbeddingsForm) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetKeepAlive returns the KeepAlive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateEmbeddingsForm) GetKeepAlive() KeepAlive {
	if o == nil || IsNil(o.KeepAlive.Get()) {
		var ret KeepAlive
		return ret
	}
	return *o.KeepAlive.Get()
}

// GetKeepAliveOk returns a tuple with the KeepAlive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateEmbeddingsForm) GetKeepAliveOk() (*KeepAlive, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeepAlive.Get(), o.KeepAlive.IsSet()
}

// HasKeepAlive returns a boolean if a field has been set.
func (o *GenerateEmbeddingsForm) HasKeepAlive() bool {
	if o != nil && o.KeepAlive.IsSet() {
		return true
	}

	return false
}

// SetKeepAlive gets a reference to the given NullableKeepAlive and assigns it to the KeepAlive field.
func (o *GenerateEmbeddingsForm) SetKeepAlive(v KeepAlive) {
	o.KeepAlive.Set(&v)
}
// SetKeepAliveNil sets the value for KeepAlive to be an explicit nil
func (o *GenerateEmbeddingsForm) SetKeepAliveNil() {
	o.KeepAlive.Set(nil)
}

// UnsetKeepAlive ensures that no value is present for KeepAlive, not even an explicit nil
func (o *GenerateEmbeddingsForm) UnsetKeepAlive() {
	o.KeepAlive.Unset()
}

func (o GenerateEmbeddingsForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateEmbeddingsForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	toSerialize["prompt"] = o.Prompt
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.KeepAlive.IsSet() {
		toSerialize["keep_alive"] = o.KeepAlive.Get()
	}
	return toSerialize, nil
}

func (o *GenerateEmbeddingsForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
		"prompt",
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

	varGenerateEmbeddingsForm := _GenerateEmbeddingsForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateEmbeddingsForm)

	if err != nil {
		return err
	}

	*o = GenerateEmbeddingsForm(varGenerateEmbeddingsForm)

	return err
}

type NullableGenerateEmbeddingsForm struct {
	value *GenerateEmbeddingsForm
	isSet bool
}

func (v NullableGenerateEmbeddingsForm) Get() *GenerateEmbeddingsForm {
	return v.value
}

func (v *NullableGenerateEmbeddingsForm) Set(val *GenerateEmbeddingsForm) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateEmbeddingsForm) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateEmbeddingsForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateEmbeddingsForm(val *GenerateEmbeddingsForm) *NullableGenerateEmbeddingsForm {
	return &NullableGenerateEmbeddingsForm{value: val, isSet: true}
}

func (v NullableGenerateEmbeddingsForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateEmbeddingsForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


