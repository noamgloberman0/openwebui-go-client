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

// checks if the SigninForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigninForm{}

// SigninForm struct for SigninForm
type SigninForm struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type _SigninForm SigninForm

// NewSigninForm instantiates a new SigninForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigninForm(email string, password string) *SigninForm {
	this := SigninForm{}
	this.Email = email
	this.Password = password
	return &this
}

// NewSigninFormWithDefaults instantiates a new SigninForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigninFormWithDefaults() *SigninForm {
	this := SigninForm{}
	return &this
}

// GetEmail returns the Email field value
func (o *SigninForm) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SigninForm) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SigninForm) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *SigninForm) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SigninForm) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SigninForm) SetPassword(v string) {
	o.Password = v
}

func (o SigninForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigninForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *SigninForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"password",
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

	varSigninForm := _SigninForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSigninForm)

	if err != nil {
		return err
	}

	*o = SigninForm(varSigninForm)

	return err
}

type NullableSigninForm struct {
	value *SigninForm
	isSet bool
}

func (v NullableSigninForm) Get() *SigninForm {
	return v.value
}

func (v *NullableSigninForm) Set(val *SigninForm) {
	v.value = val
	v.isSet = true
}

func (v NullableSigninForm) IsSet() bool {
	return v.isSet
}

func (v *NullableSigninForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigninForm(val *SigninForm) *NullableSigninForm {
	return &NullableSigninForm{value: val, isSet: true}
}

func (v NullableSigninForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigninForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


