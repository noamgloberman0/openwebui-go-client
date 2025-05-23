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

// checks if the AudioConfigUpdateForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudioConfigUpdateForm{}

// AudioConfigUpdateForm struct for AudioConfigUpdateForm
type AudioConfigUpdateForm struct {
	Tts TTSConfigForm `json:"tts"`
	Stt STTConfigForm `json:"stt"`
}

type _AudioConfigUpdateForm AudioConfigUpdateForm

// NewAudioConfigUpdateForm instantiates a new AudioConfigUpdateForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudioConfigUpdateForm(tts TTSConfigForm, stt STTConfigForm) *AudioConfigUpdateForm {
	this := AudioConfigUpdateForm{}
	this.Tts = tts
	this.Stt = stt
	return &this
}

// NewAudioConfigUpdateFormWithDefaults instantiates a new AudioConfigUpdateForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudioConfigUpdateFormWithDefaults() *AudioConfigUpdateForm {
	this := AudioConfigUpdateForm{}
	return &this
}

// GetTts returns the Tts field value
func (o *AudioConfigUpdateForm) GetTts() TTSConfigForm {
	if o == nil {
		var ret TTSConfigForm
		return ret
	}

	return o.Tts
}

// GetTtsOk returns a tuple with the Tts field value
// and a boolean to check if the value has been set.
func (o *AudioConfigUpdateForm) GetTtsOk() (*TTSConfigForm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tts, true
}

// SetTts sets field value
func (o *AudioConfigUpdateForm) SetTts(v TTSConfigForm) {
	o.Tts = v
}

// GetStt returns the Stt field value
func (o *AudioConfigUpdateForm) GetStt() STTConfigForm {
	if o == nil {
		var ret STTConfigForm
		return ret
	}

	return o.Stt
}

// GetSttOk returns a tuple with the Stt field value
// and a boolean to check if the value has been set.
func (o *AudioConfigUpdateForm) GetSttOk() (*STTConfigForm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stt, true
}

// SetStt sets field value
func (o *AudioConfigUpdateForm) SetStt(v STTConfigForm) {
	o.Stt = v
}

func (o AudioConfigUpdateForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudioConfigUpdateForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tts"] = o.Tts
	toSerialize["stt"] = o.Stt
	return toSerialize, nil
}

func (o *AudioConfigUpdateForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tts",
		"stt",
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

	varAudioConfigUpdateForm := _AudioConfigUpdateForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAudioConfigUpdateForm)

	if err != nil {
		return err
	}

	*o = AudioConfigUpdateForm(varAudioConfigUpdateForm)

	return err
}

type NullableAudioConfigUpdateForm struct {
	value *AudioConfigUpdateForm
	isSet bool
}

func (v NullableAudioConfigUpdateForm) Get() *AudioConfigUpdateForm {
	return v.value
}

func (v *NullableAudioConfigUpdateForm) Set(val *AudioConfigUpdateForm) {
	v.value = val
	v.isSet = true
}

func (v NullableAudioConfigUpdateForm) IsSet() bool {
	return v.isSet
}

func (v *NullableAudioConfigUpdateForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudioConfigUpdateForm(val *AudioConfigUpdateForm) *NullableAudioConfigUpdateForm {
	return &NullableAudioConfigUpdateForm{value: val, isSet: true}
}

func (v NullableAudioConfigUpdateForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudioConfigUpdateForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


