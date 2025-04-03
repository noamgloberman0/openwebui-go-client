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

// checks if the OpenWebuiRoutersRetrievalOpenAIConfigForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenWebuiRoutersRetrievalOpenAIConfigForm{}

// OpenWebuiRoutersRetrievalOpenAIConfigForm struct for OpenWebuiRoutersRetrievalOpenAIConfigForm
type OpenWebuiRoutersRetrievalOpenAIConfigForm struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type _OpenWebuiRoutersRetrievalOpenAIConfigForm OpenWebuiRoutersRetrievalOpenAIConfigForm

// NewOpenWebuiRoutersRetrievalOpenAIConfigForm instantiates a new OpenWebuiRoutersRetrievalOpenAIConfigForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenWebuiRoutersRetrievalOpenAIConfigForm(url string, key string) *OpenWebuiRoutersRetrievalOpenAIConfigForm {
	this := OpenWebuiRoutersRetrievalOpenAIConfigForm{}
	this.Url = url
	this.Key = key
	return &this
}

// NewOpenWebuiRoutersRetrievalOpenAIConfigFormWithDefaults instantiates a new OpenWebuiRoutersRetrievalOpenAIConfigForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenWebuiRoutersRetrievalOpenAIConfigFormWithDefaults() *OpenWebuiRoutersRetrievalOpenAIConfigForm {
	this := OpenWebuiRoutersRetrievalOpenAIConfigForm{}
	return &this
}

// GetUrl returns the Url field value
func (o *OpenWebuiRoutersRetrievalOpenAIConfigForm) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OpenWebuiRoutersRetrievalOpenAIConfigForm) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OpenWebuiRoutersRetrievalOpenAIConfigForm) SetUrl(v string) {
	o.Url = v
}

// GetKey returns the Key field value
func (o *OpenWebuiRoutersRetrievalOpenAIConfigForm) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *OpenWebuiRoutersRetrievalOpenAIConfigForm) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *OpenWebuiRoutersRetrievalOpenAIConfigForm) SetKey(v string) {
	o.Key = v
}

func (o OpenWebuiRoutersRetrievalOpenAIConfigForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenWebuiRoutersRetrievalOpenAIConfigForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

func (o *OpenWebuiRoutersRetrievalOpenAIConfigForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"key",
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

	varOpenWebuiRoutersRetrievalOpenAIConfigForm := _OpenWebuiRoutersRetrievalOpenAIConfigForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenWebuiRoutersRetrievalOpenAIConfigForm)

	if err != nil {
		return err
	}

	*o = OpenWebuiRoutersRetrievalOpenAIConfigForm(varOpenWebuiRoutersRetrievalOpenAIConfigForm)

	return err
}

type NullableOpenWebuiRoutersRetrievalOpenAIConfigForm struct {
	value *OpenWebuiRoutersRetrievalOpenAIConfigForm
	isSet bool
}

func (v NullableOpenWebuiRoutersRetrievalOpenAIConfigForm) Get() *OpenWebuiRoutersRetrievalOpenAIConfigForm {
	return v.value
}

func (v *NullableOpenWebuiRoutersRetrievalOpenAIConfigForm) Set(val *OpenWebuiRoutersRetrievalOpenAIConfigForm) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenWebuiRoutersRetrievalOpenAIConfigForm) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenWebuiRoutersRetrievalOpenAIConfigForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenWebuiRoutersRetrievalOpenAIConfigForm(val *OpenWebuiRoutersRetrievalOpenAIConfigForm) *NullableOpenWebuiRoutersRetrievalOpenAIConfigForm {
	return &NullableOpenWebuiRoutersRetrievalOpenAIConfigForm{value: val, isSet: true}
}

func (v NullableOpenWebuiRoutersRetrievalOpenAIConfigForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenWebuiRoutersRetrievalOpenAIConfigForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


