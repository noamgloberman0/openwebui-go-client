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

// checks if the RerankingModelUpdateForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RerankingModelUpdateForm{}

// RerankingModelUpdateForm struct for RerankingModelUpdateForm
type RerankingModelUpdateForm struct {
	RerankingModel string `json:"reranking_model"`
}

type _RerankingModelUpdateForm RerankingModelUpdateForm

// NewRerankingModelUpdateForm instantiates a new RerankingModelUpdateForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRerankingModelUpdateForm(rerankingModel string) *RerankingModelUpdateForm {
	this := RerankingModelUpdateForm{}
	this.RerankingModel = rerankingModel
	return &this
}

// NewRerankingModelUpdateFormWithDefaults instantiates a new RerankingModelUpdateForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRerankingModelUpdateFormWithDefaults() *RerankingModelUpdateForm {
	this := RerankingModelUpdateForm{}
	return &this
}

// GetRerankingModel returns the RerankingModel field value
func (o *RerankingModelUpdateForm) GetRerankingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RerankingModel
}

// GetRerankingModelOk returns a tuple with the RerankingModel field value
// and a boolean to check if the value has been set.
func (o *RerankingModelUpdateForm) GetRerankingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RerankingModel, true
}

// SetRerankingModel sets field value
func (o *RerankingModelUpdateForm) SetRerankingModel(v string) {
	o.RerankingModel = v
}

func (o RerankingModelUpdateForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RerankingModelUpdateForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reranking_model"] = o.RerankingModel
	return toSerialize, nil
}

func (o *RerankingModelUpdateForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reranking_model",
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

	varRerankingModelUpdateForm := _RerankingModelUpdateForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRerankingModelUpdateForm)

	if err != nil {
		return err
	}

	*o = RerankingModelUpdateForm(varRerankingModelUpdateForm)

	return err
}

type NullableRerankingModelUpdateForm struct {
	value *RerankingModelUpdateForm
	isSet bool
}

func (v NullableRerankingModelUpdateForm) Get() *RerankingModelUpdateForm {
	return v.value
}

func (v *NullableRerankingModelUpdateForm) Set(val *RerankingModelUpdateForm) {
	v.value = val
	v.isSet = true
}

func (v NullableRerankingModelUpdateForm) IsSet() bool {
	return v.isSet
}

func (v *NullableRerankingModelUpdateForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRerankingModelUpdateForm(val *RerankingModelUpdateForm) *NullableRerankingModelUpdateForm {
	return &NullableRerankingModelUpdateForm{value: val, isSet: true}
}

func (v NullableRerankingModelUpdateForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRerankingModelUpdateForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


