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

// checks if the OpenWebuiRoutersImagesOpenAIConfigForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenWebuiRoutersImagesOpenAIConfigForm{}

// OpenWebuiRoutersImagesOpenAIConfigForm struct for OpenWebuiRoutersImagesOpenAIConfigForm
type OpenWebuiRoutersImagesOpenAIConfigForm struct {
	OPENAI_API_BASE_URL string `json:"OPENAI_API_BASE_URL"`
	OPENAI_API_KEY string `json:"OPENAI_API_KEY"`
}

type _OpenWebuiRoutersImagesOpenAIConfigForm OpenWebuiRoutersImagesOpenAIConfigForm

// NewOpenWebuiRoutersImagesOpenAIConfigForm instantiates a new OpenWebuiRoutersImagesOpenAIConfigForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenWebuiRoutersImagesOpenAIConfigForm(oPENAIAPIBASEURL string, oPENAIAPIKEY string) *OpenWebuiRoutersImagesOpenAIConfigForm {
	this := OpenWebuiRoutersImagesOpenAIConfigForm{}
	this.OPENAI_API_BASE_URL = oPENAIAPIBASEURL
	this.OPENAI_API_KEY = oPENAIAPIKEY
	return &this
}

// NewOpenWebuiRoutersImagesOpenAIConfigFormWithDefaults instantiates a new OpenWebuiRoutersImagesOpenAIConfigForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenWebuiRoutersImagesOpenAIConfigFormWithDefaults() *OpenWebuiRoutersImagesOpenAIConfigForm {
	this := OpenWebuiRoutersImagesOpenAIConfigForm{}
	return &this
}

// GetOPENAI_API_BASE_URL returns the OPENAI_API_BASE_URL field value
func (o *OpenWebuiRoutersImagesOpenAIConfigForm) GetOPENAI_API_BASE_URL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OPENAI_API_BASE_URL
}

// GetOPENAI_API_BASE_URLOk returns a tuple with the OPENAI_API_BASE_URL field value
// and a boolean to check if the value has been set.
func (o *OpenWebuiRoutersImagesOpenAIConfigForm) GetOPENAI_API_BASE_URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OPENAI_API_BASE_URL, true
}

// SetOPENAI_API_BASE_URL sets field value
func (o *OpenWebuiRoutersImagesOpenAIConfigForm) SetOPENAI_API_BASE_URL(v string) {
	o.OPENAI_API_BASE_URL = v
}

// GetOPENAI_API_KEY returns the OPENAI_API_KEY field value
func (o *OpenWebuiRoutersImagesOpenAIConfigForm) GetOPENAI_API_KEY() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OPENAI_API_KEY
}

// GetOPENAI_API_KEYOk returns a tuple with the OPENAI_API_KEY field value
// and a boolean to check if the value has been set.
func (o *OpenWebuiRoutersImagesOpenAIConfigForm) GetOPENAI_API_KEYOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OPENAI_API_KEY, true
}

// SetOPENAI_API_KEY sets field value
func (o *OpenWebuiRoutersImagesOpenAIConfigForm) SetOPENAI_API_KEY(v string) {
	o.OPENAI_API_KEY = v
}

func (o OpenWebuiRoutersImagesOpenAIConfigForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenWebuiRoutersImagesOpenAIConfigForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["OPENAI_API_BASE_URL"] = o.OPENAI_API_BASE_URL
	toSerialize["OPENAI_API_KEY"] = o.OPENAI_API_KEY
	return toSerialize, nil
}

func (o *OpenWebuiRoutersImagesOpenAIConfigForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"OPENAI_API_BASE_URL",
		"OPENAI_API_KEY",
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

	varOpenWebuiRoutersImagesOpenAIConfigForm := _OpenWebuiRoutersImagesOpenAIConfigForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenWebuiRoutersImagesOpenAIConfigForm)

	if err != nil {
		return err
	}

	*o = OpenWebuiRoutersImagesOpenAIConfigForm(varOpenWebuiRoutersImagesOpenAIConfigForm)

	return err
}

type NullableOpenWebuiRoutersImagesOpenAIConfigForm struct {
	value *OpenWebuiRoutersImagesOpenAIConfigForm
	isSet bool
}

func (v NullableOpenWebuiRoutersImagesOpenAIConfigForm) Get() *OpenWebuiRoutersImagesOpenAIConfigForm {
	return v.value
}

func (v *NullableOpenWebuiRoutersImagesOpenAIConfigForm) Set(val *OpenWebuiRoutersImagesOpenAIConfigForm) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenWebuiRoutersImagesOpenAIConfigForm) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenWebuiRoutersImagesOpenAIConfigForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenWebuiRoutersImagesOpenAIConfigForm(val *OpenWebuiRoutersImagesOpenAIConfigForm) *NullableOpenWebuiRoutersImagesOpenAIConfigForm {
	return &NullableOpenWebuiRoutersImagesOpenAIConfigForm{value: val, isSet: true}
}

func (v NullableOpenWebuiRoutersImagesOpenAIConfigForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenWebuiRoutersImagesOpenAIConfigForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


