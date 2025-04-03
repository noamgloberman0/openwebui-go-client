/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// ProfileImageUrl struct for ProfileImageUrl
type ProfileImageUrl struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProfileImageUrl) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ProfileImageUrl)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProfileImageUrl) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableProfileImageUrl struct {
	value *ProfileImageUrl
	isSet bool
}

func (v NullableProfileImageUrl) Get() *ProfileImageUrl {
	return v.value
}

func (v *NullableProfileImageUrl) Set(val *ProfileImageUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileImageUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileImageUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileImageUrl(val *ProfileImageUrl) *NullableProfileImageUrl {
	return &NullableProfileImageUrl{value: val, isSet: true}
}

func (v NullableProfileImageUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileImageUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


