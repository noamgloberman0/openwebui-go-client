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

// checks if the CodeInterpreterConfigForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeInterpreterConfigForm{}

// CodeInterpreterConfigForm struct for CodeInterpreterConfigForm
type CodeInterpreterConfigForm struct {
	ENABLE_CODE_EXECUTION bool `json:"ENABLE_CODE_EXECUTION"`
	CODE_EXECUTION_ENGINE string `json:"CODE_EXECUTION_ENGINE"`
	CODE_EXECUTION_JUPYTER_URL NullableString `json:"CODE_EXECUTION_JUPYTER_URL"`
	CODE_EXECUTION_JUPYTER_AUTH NullableString `json:"CODE_EXECUTION_JUPYTER_AUTH"`
	CODE_EXECUTION_JUPYTER_AUTH_TOKEN NullableString `json:"CODE_EXECUTION_JUPYTER_AUTH_TOKEN"`
	CODE_EXECUTION_JUPYTER_AUTH_PASSWORD NullableString `json:"CODE_EXECUTION_JUPYTER_AUTH_PASSWORD"`
	CODE_EXECUTION_JUPYTER_TIMEOUT NullableInt32 `json:"CODE_EXECUTION_JUPYTER_TIMEOUT"`
	ENABLE_CODE_INTERPRETER bool `json:"ENABLE_CODE_INTERPRETER"`
	CODE_INTERPRETER_ENGINE string `json:"CODE_INTERPRETER_ENGINE"`
	CODE_INTERPRETER_PROMPT_TEMPLATE NullableString `json:"CODE_INTERPRETER_PROMPT_TEMPLATE"`
	CODE_INTERPRETER_JUPYTER_URL NullableString `json:"CODE_INTERPRETER_JUPYTER_URL"`
	CODE_INTERPRETER_JUPYTER_AUTH NullableString `json:"CODE_INTERPRETER_JUPYTER_AUTH"`
	CODE_INTERPRETER_JUPYTER_AUTH_TOKEN NullableString `json:"CODE_INTERPRETER_JUPYTER_AUTH_TOKEN"`
	CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD NullableString `json:"CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD"`
	CODE_INTERPRETER_JUPYTER_TIMEOUT NullableInt32 `json:"CODE_INTERPRETER_JUPYTER_TIMEOUT"`
}

type _CodeInterpreterConfigForm CodeInterpreterConfigForm

// NewCodeInterpreterConfigForm instantiates a new CodeInterpreterConfigForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeInterpreterConfigForm(eNABLECODEEXECUTION bool, cODEEXECUTIONENGINE string, cODEEXECUTIONJUPYTERURL NullableString, cODEEXECUTIONJUPYTERAUTH NullableString, cODEEXECUTIONJUPYTERAUTHTOKEN NullableString, cODEEXECUTIONJUPYTERAUTHPASSWORD NullableString, cODEEXECUTIONJUPYTERTIMEOUT NullableInt32, eNABLECODEINTERPRETER bool, cODEINTERPRETERENGINE string, cODEINTERPRETERPROMPTTEMPLATE NullableString, cODEINTERPRETERJUPYTERURL NullableString, cODEINTERPRETERJUPYTERAUTH NullableString, cODEINTERPRETERJUPYTERAUTHTOKEN NullableString, cODEINTERPRETERJUPYTERAUTHPASSWORD NullableString, cODEINTERPRETERJUPYTERTIMEOUT NullableInt32) *CodeInterpreterConfigForm {
	this := CodeInterpreterConfigForm{}
	this.ENABLE_CODE_EXECUTION = eNABLECODEEXECUTION
	this.CODE_EXECUTION_ENGINE = cODEEXECUTIONENGINE
	this.CODE_EXECUTION_JUPYTER_URL = cODEEXECUTIONJUPYTERURL
	this.CODE_EXECUTION_JUPYTER_AUTH = cODEEXECUTIONJUPYTERAUTH
	this.CODE_EXECUTION_JUPYTER_AUTH_TOKEN = cODEEXECUTIONJUPYTERAUTHTOKEN
	this.CODE_EXECUTION_JUPYTER_AUTH_PASSWORD = cODEEXECUTIONJUPYTERAUTHPASSWORD
	this.CODE_EXECUTION_JUPYTER_TIMEOUT = cODEEXECUTIONJUPYTERTIMEOUT
	this.ENABLE_CODE_INTERPRETER = eNABLECODEINTERPRETER
	this.CODE_INTERPRETER_ENGINE = cODEINTERPRETERENGINE
	this.CODE_INTERPRETER_PROMPT_TEMPLATE = cODEINTERPRETERPROMPTTEMPLATE
	this.CODE_INTERPRETER_JUPYTER_URL = cODEINTERPRETERJUPYTERURL
	this.CODE_INTERPRETER_JUPYTER_AUTH = cODEINTERPRETERJUPYTERAUTH
	this.CODE_INTERPRETER_JUPYTER_AUTH_TOKEN = cODEINTERPRETERJUPYTERAUTHTOKEN
	this.CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD = cODEINTERPRETERJUPYTERAUTHPASSWORD
	this.CODE_INTERPRETER_JUPYTER_TIMEOUT = cODEINTERPRETERJUPYTERTIMEOUT
	return &this
}

// NewCodeInterpreterConfigFormWithDefaults instantiates a new CodeInterpreterConfigForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeInterpreterConfigFormWithDefaults() *CodeInterpreterConfigForm {
	this := CodeInterpreterConfigForm{}
	return &this
}

// GetENABLE_CODE_EXECUTION returns the ENABLE_CODE_EXECUTION field value
func (o *CodeInterpreterConfigForm) GetENABLE_CODE_EXECUTION() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ENABLE_CODE_EXECUTION
}

// GetENABLE_CODE_EXECUTIONOk returns a tuple with the ENABLE_CODE_EXECUTION field value
// and a boolean to check if the value has been set.
func (o *CodeInterpreterConfigForm) GetENABLE_CODE_EXECUTIONOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ENABLE_CODE_EXECUTION, true
}

// SetENABLE_CODE_EXECUTION sets field value
func (o *CodeInterpreterConfigForm) SetENABLE_CODE_EXECUTION(v bool) {
	o.ENABLE_CODE_EXECUTION = v
}

// GetCODE_EXECUTION_ENGINE returns the CODE_EXECUTION_ENGINE field value
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_ENGINE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CODE_EXECUTION_ENGINE
}

// GetCODE_EXECUTION_ENGINEOk returns a tuple with the CODE_EXECUTION_ENGINE field value
// and a boolean to check if the value has been set.
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_ENGINEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CODE_EXECUTION_ENGINE, true
}

// SetCODE_EXECUTION_ENGINE sets field value
func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_ENGINE(v string) {
	o.CODE_EXECUTION_ENGINE = v
}

// GetCODE_EXECUTION_JUPYTER_URL returns the CODE_EXECUTION_JUPYTER_URL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_URL() string {
	if o == nil || o.CODE_EXECUTION_JUPYTER_URL.Get() == nil {
		var ret string
		return ret
	}

	return *o.CODE_EXECUTION_JUPYTER_URL.Get()
}

// GetCODE_EXECUTION_JUPYTER_URLOk returns a tuple with the CODE_EXECUTION_JUPYTER_URL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_EXECUTION_JUPYTER_URL.Get(), o.CODE_EXECUTION_JUPYTER_URL.IsSet()
}

// SetCODE_EXECUTION_JUPYTER_URL sets field value
func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_URL(v string) {
	o.CODE_EXECUTION_JUPYTER_URL.Set(&v)
}

// GetCODE_EXECUTION_JUPYTER_AUTH returns the CODE_EXECUTION_JUPYTER_AUTH field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH() string {
	if o == nil || o.CODE_EXECUTION_JUPYTER_AUTH.Get() == nil {
		var ret string
		return ret
	}

	return *o.CODE_EXECUTION_JUPYTER_AUTH.Get()
}

// GetCODE_EXECUTION_JUPYTER_AUTHOk returns a tuple with the CODE_EXECUTION_JUPYTER_AUTH field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTHOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_EXECUTION_JUPYTER_AUTH.Get(), o.CODE_EXECUTION_JUPYTER_AUTH.IsSet()
}

// SetCODE_EXECUTION_JUPYTER_AUTH sets field value
func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_AUTH(v string) {
	o.CODE_EXECUTION_JUPYTER_AUTH.Set(&v)
}

// GetCODE_EXECUTION_JUPYTER_AUTH_TOKEN returns the CODE_EXECUTION_JUPYTER_AUTH_TOKEN field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH_TOKEN() string {
	if o == nil || o.CODE_EXECUTION_JUPYTER_AUTH_TOKEN.Get() == nil {
		var ret string
		return ret
	}

	return *o.CODE_EXECUTION_JUPYTER_AUTH_TOKEN.Get()
}

// GetCODE_EXECUTION_JUPYTER_AUTH_TOKENOk returns a tuple with the CODE_EXECUTION_JUPYTER_AUTH_TOKEN field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH_TOKENOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_EXECUTION_JUPYTER_AUTH_TOKEN.Get(), o.CODE_EXECUTION_JUPYTER_AUTH_TOKEN.IsSet()
}

// SetCODE_EXECUTION_JUPYTER_AUTH_TOKEN sets field value
func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_AUTH_TOKEN(v string) {
	o.CODE_EXECUTION_JUPYTER_AUTH_TOKEN.Set(&v)
}

// GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD returns the CODE_EXECUTION_JUPYTER_AUTH_PASSWORD field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD() string {
	if o == nil || o.CODE_EXECUTION_JUPYTER_AUTH_PASSWORD.Get() == nil {
		var ret string
		return ret
	}

	return *o.CODE_EXECUTION_JUPYTER_AUTH_PASSWORD.Get()
}

// GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORDOk returns a tuple with the CODE_EXECUTION_JUPYTER_AUTH_PASSWORD field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_AUTH_PASSWORDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_EXECUTION_JUPYTER_AUTH_PASSWORD.Get(), o.CODE_EXECUTION_JUPYTER_AUTH_PASSWORD.IsSet()
}

// SetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD sets field value
func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_AUTH_PASSWORD(v string) {
	o.CODE_EXECUTION_JUPYTER_AUTH_PASSWORD.Set(&v)
}

// GetCODE_EXECUTION_JUPYTER_TIMEOUT returns the CODE_EXECUTION_JUPYTER_TIMEOUT field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_TIMEOUT() int32 {
	if o == nil || o.CODE_EXECUTION_JUPYTER_TIMEOUT.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CODE_EXECUTION_JUPYTER_TIMEOUT.Get()
}

// GetCODE_EXECUTION_JUPYTER_TIMEOUTOk returns a tuple with the CODE_EXECUTION_JUPYTER_TIMEOUT field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_EXECUTION_JUPYTER_TIMEOUTOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_EXECUTION_JUPYTER_TIMEOUT.Get(), o.CODE_EXECUTION_JUPYTER_TIMEOUT.IsSet()
}

// SetCODE_EXECUTION_JUPYTER_TIMEOUT sets field value
func (o *CodeInterpreterConfigForm) SetCODE_EXECUTION_JUPYTER_TIMEOUT(v int32) {
	o.CODE_EXECUTION_JUPYTER_TIMEOUT.Set(&v)
}

// GetENABLE_CODE_INTERPRETER returns the ENABLE_CODE_INTERPRETER field value
func (o *CodeInterpreterConfigForm) GetENABLE_CODE_INTERPRETER() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ENABLE_CODE_INTERPRETER
}

// GetENABLE_CODE_INTERPRETEROk returns a tuple with the ENABLE_CODE_INTERPRETER field value
// and a boolean to check if the value has been set.
func (o *CodeInterpreterConfigForm) GetENABLE_CODE_INTERPRETEROk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ENABLE_CODE_INTERPRETER, true
}

// SetENABLE_CODE_INTERPRETER sets field value
func (o *CodeInterpreterConfigForm) SetENABLE_CODE_INTERPRETER(v bool) {
	o.ENABLE_CODE_INTERPRETER = v
}

// GetCODE_INTERPRETER_ENGINE returns the CODE_INTERPRETER_ENGINE field value
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_ENGINE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CODE_INTERPRETER_ENGINE
}

// GetCODE_INTERPRETER_ENGINEOk returns a tuple with the CODE_INTERPRETER_ENGINE field value
// and a boolean to check if the value has been set.
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_ENGINEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CODE_INTERPRETER_ENGINE, true
}

// SetCODE_INTERPRETER_ENGINE sets field value
func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_ENGINE(v string) {
	o.CODE_INTERPRETER_ENGINE = v
}

// GetCODE_INTERPRETER_PROMPT_TEMPLATE returns the CODE_INTERPRETER_PROMPT_TEMPLATE field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_PROMPT_TEMPLATE() string {
	if o == nil || o.CODE_INTERPRETER_PROMPT_TEMPLATE.Get() == nil {
		var ret string
		return ret
	}

	return *o.CODE_INTERPRETER_PROMPT_TEMPLATE.Get()
}

// GetCODE_INTERPRETER_PROMPT_TEMPLATEOk returns a tuple with the CODE_INTERPRETER_PROMPT_TEMPLATE field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_PROMPT_TEMPLATEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_INTERPRETER_PROMPT_TEMPLATE.Get(), o.CODE_INTERPRETER_PROMPT_TEMPLATE.IsSet()
}

// SetCODE_INTERPRETER_PROMPT_TEMPLATE sets field value
func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_PROMPT_TEMPLATE(v string) {
	o.CODE_INTERPRETER_PROMPT_TEMPLATE.Set(&v)
}

// GetCODE_INTERPRETER_JUPYTER_URL returns the CODE_INTERPRETER_JUPYTER_URL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_URL() string {
	if o == nil || o.CODE_INTERPRETER_JUPYTER_URL.Get() == nil {
		var ret string
		return ret
	}

	return *o.CODE_INTERPRETER_JUPYTER_URL.Get()
}

// GetCODE_INTERPRETER_JUPYTER_URLOk returns a tuple with the CODE_INTERPRETER_JUPYTER_URL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_INTERPRETER_JUPYTER_URL.Get(), o.CODE_INTERPRETER_JUPYTER_URL.IsSet()
}

// SetCODE_INTERPRETER_JUPYTER_URL sets field value
func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_URL(v string) {
	o.CODE_INTERPRETER_JUPYTER_URL.Set(&v)
}

// GetCODE_INTERPRETER_JUPYTER_AUTH returns the CODE_INTERPRETER_JUPYTER_AUTH field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH() string {
	if o == nil || o.CODE_INTERPRETER_JUPYTER_AUTH.Get() == nil {
		var ret string
		return ret
	}

	return *o.CODE_INTERPRETER_JUPYTER_AUTH.Get()
}

// GetCODE_INTERPRETER_JUPYTER_AUTHOk returns a tuple with the CODE_INTERPRETER_JUPYTER_AUTH field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTHOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_INTERPRETER_JUPYTER_AUTH.Get(), o.CODE_INTERPRETER_JUPYTER_AUTH.IsSet()
}

// SetCODE_INTERPRETER_JUPYTER_AUTH sets field value
func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_AUTH(v string) {
	o.CODE_INTERPRETER_JUPYTER_AUTH.Set(&v)
}

// GetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN returns the CODE_INTERPRETER_JUPYTER_AUTH_TOKEN field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN() string {
	if o == nil || o.CODE_INTERPRETER_JUPYTER_AUTH_TOKEN.Get() == nil {
		var ret string
		return ret
	}

	return *o.CODE_INTERPRETER_JUPYTER_AUTH_TOKEN.Get()
}

// GetCODE_INTERPRETER_JUPYTER_AUTH_TOKENOk returns a tuple with the CODE_INTERPRETER_JUPYTER_AUTH_TOKEN field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH_TOKENOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_INTERPRETER_JUPYTER_AUTH_TOKEN.Get(), o.CODE_INTERPRETER_JUPYTER_AUTH_TOKEN.IsSet()
}

// SetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN sets field value
func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_AUTH_TOKEN(v string) {
	o.CODE_INTERPRETER_JUPYTER_AUTH_TOKEN.Set(&v)
}

// GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD returns the CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD() string {
	if o == nil || o.CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD.Get() == nil {
		var ret string
		return ret
	}

	return *o.CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD.Get()
}

// GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORDOk returns a tuple with the CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD.Get(), o.CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD.IsSet()
}

// SetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD sets field value
func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_AUTH_PASSWORD(v string) {
	o.CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD.Set(&v)
}

// GetCODE_INTERPRETER_JUPYTER_TIMEOUT returns the CODE_INTERPRETER_JUPYTER_TIMEOUT field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_TIMEOUT() int32 {
	if o == nil || o.CODE_INTERPRETER_JUPYTER_TIMEOUT.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CODE_INTERPRETER_JUPYTER_TIMEOUT.Get()
}

// GetCODE_INTERPRETER_JUPYTER_TIMEOUTOk returns a tuple with the CODE_INTERPRETER_JUPYTER_TIMEOUT field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeInterpreterConfigForm) GetCODE_INTERPRETER_JUPYTER_TIMEOUTOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CODE_INTERPRETER_JUPYTER_TIMEOUT.Get(), o.CODE_INTERPRETER_JUPYTER_TIMEOUT.IsSet()
}

// SetCODE_INTERPRETER_JUPYTER_TIMEOUT sets field value
func (o *CodeInterpreterConfigForm) SetCODE_INTERPRETER_JUPYTER_TIMEOUT(v int32) {
	o.CODE_INTERPRETER_JUPYTER_TIMEOUT.Set(&v)
}

func (o CodeInterpreterConfigForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeInterpreterConfigForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ENABLE_CODE_EXECUTION"] = o.ENABLE_CODE_EXECUTION
	toSerialize["CODE_EXECUTION_ENGINE"] = o.CODE_EXECUTION_ENGINE
	toSerialize["CODE_EXECUTION_JUPYTER_URL"] = o.CODE_EXECUTION_JUPYTER_URL.Get()
	toSerialize["CODE_EXECUTION_JUPYTER_AUTH"] = o.CODE_EXECUTION_JUPYTER_AUTH.Get()
	toSerialize["CODE_EXECUTION_JUPYTER_AUTH_TOKEN"] = o.CODE_EXECUTION_JUPYTER_AUTH_TOKEN.Get()
	toSerialize["CODE_EXECUTION_JUPYTER_AUTH_PASSWORD"] = o.CODE_EXECUTION_JUPYTER_AUTH_PASSWORD.Get()
	toSerialize["CODE_EXECUTION_JUPYTER_TIMEOUT"] = o.CODE_EXECUTION_JUPYTER_TIMEOUT.Get()
	toSerialize["ENABLE_CODE_INTERPRETER"] = o.ENABLE_CODE_INTERPRETER
	toSerialize["CODE_INTERPRETER_ENGINE"] = o.CODE_INTERPRETER_ENGINE
	toSerialize["CODE_INTERPRETER_PROMPT_TEMPLATE"] = o.CODE_INTERPRETER_PROMPT_TEMPLATE.Get()
	toSerialize["CODE_INTERPRETER_JUPYTER_URL"] = o.CODE_INTERPRETER_JUPYTER_URL.Get()
	toSerialize["CODE_INTERPRETER_JUPYTER_AUTH"] = o.CODE_INTERPRETER_JUPYTER_AUTH.Get()
	toSerialize["CODE_INTERPRETER_JUPYTER_AUTH_TOKEN"] = o.CODE_INTERPRETER_JUPYTER_AUTH_TOKEN.Get()
	toSerialize["CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD"] = o.CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD.Get()
	toSerialize["CODE_INTERPRETER_JUPYTER_TIMEOUT"] = o.CODE_INTERPRETER_JUPYTER_TIMEOUT.Get()
	return toSerialize, nil
}

func (o *CodeInterpreterConfigForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ENABLE_CODE_EXECUTION",
		"CODE_EXECUTION_ENGINE",
		"CODE_EXECUTION_JUPYTER_URL",
		"CODE_EXECUTION_JUPYTER_AUTH",
		"CODE_EXECUTION_JUPYTER_AUTH_TOKEN",
		"CODE_EXECUTION_JUPYTER_AUTH_PASSWORD",
		"CODE_EXECUTION_JUPYTER_TIMEOUT",
		"ENABLE_CODE_INTERPRETER",
		"CODE_INTERPRETER_ENGINE",
		"CODE_INTERPRETER_PROMPT_TEMPLATE",
		"CODE_INTERPRETER_JUPYTER_URL",
		"CODE_INTERPRETER_JUPYTER_AUTH",
		"CODE_INTERPRETER_JUPYTER_AUTH_TOKEN",
		"CODE_INTERPRETER_JUPYTER_AUTH_PASSWORD",
		"CODE_INTERPRETER_JUPYTER_TIMEOUT",
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

	varCodeInterpreterConfigForm := _CodeInterpreterConfigForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCodeInterpreterConfigForm)

	if err != nil {
		return err
	}

	*o = CodeInterpreterConfigForm(varCodeInterpreterConfigForm)

	return err
}

type NullableCodeInterpreterConfigForm struct {
	value *CodeInterpreterConfigForm
	isSet bool
}

func (v NullableCodeInterpreterConfigForm) Get() *CodeInterpreterConfigForm {
	return v.value
}

func (v *NullableCodeInterpreterConfigForm) Set(val *CodeInterpreterConfigForm) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeInterpreterConfigForm) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeInterpreterConfigForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeInterpreterConfigForm(val *CodeInterpreterConfigForm) *NullableCodeInterpreterConfigForm {
	return &NullableCodeInterpreterConfigForm{value: val, isSet: true}
}

func (v NullableCodeInterpreterConfigForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeInterpreterConfigForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


