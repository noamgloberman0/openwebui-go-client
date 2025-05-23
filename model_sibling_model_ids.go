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


// SiblingModelIds struct for SiblingModelIds
type SiblingModelIds struct {
	ArrayOfString *[]string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SiblingModelIds) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ArrayOfString
	err = json.Unmarshal(data, &dst.ArrayOfString);
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			return nil // data stored in dst.ArrayOfString, return on the first match
		}
	} else {
		dst.ArrayOfString = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SiblingModelIds)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SiblingModelIds) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableSiblingModelIds struct {
	value *SiblingModelIds
	isSet bool
}

func (v NullableSiblingModelIds) Get() *SiblingModelIds {
	return v.value
}

func (v *NullableSiblingModelIds) Set(val *SiblingModelIds) {
	v.value = val
	v.isSet = true
}

func (v NullableSiblingModelIds) IsSet() bool {
	return v.isSet
}

func (v *NullableSiblingModelIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiblingModelIds(val *SiblingModelIds) *NullableSiblingModelIds {
	return &NullableSiblingModelIds{value: val, isSet: true}
}

func (v NullableSiblingModelIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiblingModelIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


