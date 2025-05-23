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

// checks if the ComfyUIConfigForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComfyUIConfigForm{}

// ComfyUIConfigForm struct for ComfyUIConfigForm
type ComfyUIConfigForm struct {
	COMFYUI_BASE_URL string `json:"COMFYUI_BASE_URL"`
	COMFYUI_API_KEY string `json:"COMFYUI_API_KEY"`
	COMFYUI_WORKFLOW string `json:"COMFYUI_WORKFLOW"`
	COMFYUI_WORKFLOW_NODES []map[string]interface{} `json:"COMFYUI_WORKFLOW_NODES"`
}

type _ComfyUIConfigForm ComfyUIConfigForm

// NewComfyUIConfigForm instantiates a new ComfyUIConfigForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComfyUIConfigForm(cOMFYUIBASEURL string, cOMFYUIAPIKEY string, cOMFYUIWORKFLOW string, cOMFYUIWORKFLOWNODES []map[string]interface{}) *ComfyUIConfigForm {
	this := ComfyUIConfigForm{}
	this.COMFYUI_BASE_URL = cOMFYUIBASEURL
	this.COMFYUI_API_KEY = cOMFYUIAPIKEY
	this.COMFYUI_WORKFLOW = cOMFYUIWORKFLOW
	this.COMFYUI_WORKFLOW_NODES = cOMFYUIWORKFLOWNODES
	return &this
}

// NewComfyUIConfigFormWithDefaults instantiates a new ComfyUIConfigForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComfyUIConfigFormWithDefaults() *ComfyUIConfigForm {
	this := ComfyUIConfigForm{}
	return &this
}

// GetCOMFYUI_BASE_URL returns the COMFYUI_BASE_URL field value
func (o *ComfyUIConfigForm) GetCOMFYUI_BASE_URL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.COMFYUI_BASE_URL
}

// GetCOMFYUI_BASE_URLOk returns a tuple with the COMFYUI_BASE_URL field value
// and a boolean to check if the value has been set.
func (o *ComfyUIConfigForm) GetCOMFYUI_BASE_URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.COMFYUI_BASE_URL, true
}

// SetCOMFYUI_BASE_URL sets field value
func (o *ComfyUIConfigForm) SetCOMFYUI_BASE_URL(v string) {
	o.COMFYUI_BASE_URL = v
}

// GetCOMFYUI_API_KEY returns the COMFYUI_API_KEY field value
func (o *ComfyUIConfigForm) GetCOMFYUI_API_KEY() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.COMFYUI_API_KEY
}

// GetCOMFYUI_API_KEYOk returns a tuple with the COMFYUI_API_KEY field value
// and a boolean to check if the value has been set.
func (o *ComfyUIConfigForm) GetCOMFYUI_API_KEYOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.COMFYUI_API_KEY, true
}

// SetCOMFYUI_API_KEY sets field value
func (o *ComfyUIConfigForm) SetCOMFYUI_API_KEY(v string) {
	o.COMFYUI_API_KEY = v
}

// GetCOMFYUI_WORKFLOW returns the COMFYUI_WORKFLOW field value
func (o *ComfyUIConfigForm) GetCOMFYUI_WORKFLOW() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.COMFYUI_WORKFLOW
}

// GetCOMFYUI_WORKFLOWOk returns a tuple with the COMFYUI_WORKFLOW field value
// and a boolean to check if the value has been set.
func (o *ComfyUIConfigForm) GetCOMFYUI_WORKFLOWOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.COMFYUI_WORKFLOW, true
}

// SetCOMFYUI_WORKFLOW sets field value
func (o *ComfyUIConfigForm) SetCOMFYUI_WORKFLOW(v string) {
	o.COMFYUI_WORKFLOW = v
}

// GetCOMFYUI_WORKFLOW_NODES returns the COMFYUI_WORKFLOW_NODES field value
func (o *ComfyUIConfigForm) GetCOMFYUI_WORKFLOW_NODES() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.COMFYUI_WORKFLOW_NODES
}

// GetCOMFYUI_WORKFLOW_NODESOk returns a tuple with the COMFYUI_WORKFLOW_NODES field value
// and a boolean to check if the value has been set.
func (o *ComfyUIConfigForm) GetCOMFYUI_WORKFLOW_NODESOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.COMFYUI_WORKFLOW_NODES, true
}

// SetCOMFYUI_WORKFLOW_NODES sets field value
func (o *ComfyUIConfigForm) SetCOMFYUI_WORKFLOW_NODES(v []map[string]interface{}) {
	o.COMFYUI_WORKFLOW_NODES = v
}

func (o ComfyUIConfigForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComfyUIConfigForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["COMFYUI_BASE_URL"] = o.COMFYUI_BASE_URL
	toSerialize["COMFYUI_API_KEY"] = o.COMFYUI_API_KEY
	toSerialize["COMFYUI_WORKFLOW"] = o.COMFYUI_WORKFLOW
	toSerialize["COMFYUI_WORKFLOW_NODES"] = o.COMFYUI_WORKFLOW_NODES
	return toSerialize, nil
}

func (o *ComfyUIConfigForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"COMFYUI_BASE_URL",
		"COMFYUI_API_KEY",
		"COMFYUI_WORKFLOW",
		"COMFYUI_WORKFLOW_NODES",
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

	varComfyUIConfigForm := _ComfyUIConfigForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varComfyUIConfigForm)

	if err != nil {
		return err
	}

	*o = ComfyUIConfigForm(varComfyUIConfigForm)

	return err
}

type NullableComfyUIConfigForm struct {
	value *ComfyUIConfigForm
	isSet bool
}

func (v NullableComfyUIConfigForm) Get() *ComfyUIConfigForm {
	return v.value
}

func (v *NullableComfyUIConfigForm) Set(val *ComfyUIConfigForm) {
	v.value = val
	v.isSet = true
}

func (v NullableComfyUIConfigForm) IsSet() bool {
	return v.isSet
}

func (v *NullableComfyUIConfigForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComfyUIConfigForm(val *ComfyUIConfigForm) *NullableComfyUIConfigForm {
	return &NullableComfyUIConfigForm{value: val, isSet: true}
}

func (v NullableComfyUIConfigForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComfyUIConfigForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


