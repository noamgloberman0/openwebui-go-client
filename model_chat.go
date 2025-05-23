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


// Chat struct for Chat
type Chat struct {
	MapmapOfStringAny *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Chat) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into MapmapOfStringAny
	err = json.Unmarshal(data, &dst.MapmapOfStringAny);
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			return nil // data stored in dst.MapmapOfStringAny, return on the first match
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Chat)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Chat) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableChat struct {
	value *Chat
	isSet bool
}

func (v NullableChat) Get() *Chat {
	return v.value
}

func (v *NullableChat) Set(val *Chat) {
	v.value = val
	v.isSet = true
}

func (v NullableChat) IsSet() bool {
	return v.isSet
}

func (v *NullableChat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChat(val *Chat) *NullableChat {
	return &NullableChat{value: val, isSet: true}
}

func (v NullableChat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


