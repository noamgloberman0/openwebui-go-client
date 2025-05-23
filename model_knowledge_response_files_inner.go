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


// KnowledgeResponseFilesInner struct for KnowledgeResponseFilesInner
type KnowledgeResponseFilesInner struct {
	FileMetadataResponse *FileMetadataResponse
	MapmapOfStringAny *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *KnowledgeResponseFilesInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FileMetadataResponse
	err = json.Unmarshal(data, &dst.FileMetadataResponse);
	if err == nil {
		jsonFileMetadataResponse, _ := json.Marshal(dst.FileMetadataResponse)
		if string(jsonFileMetadataResponse) == "{}" { // empty struct
			dst.FileMetadataResponse = nil
		} else {
			return nil // data stored in dst.FileMetadataResponse, return on the first match
		}
	} else {
		dst.FileMetadataResponse = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(KnowledgeResponseFilesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KnowledgeResponseFilesInner) MarshalJSON() ([]byte, error) {
	if src.FileMetadataResponse != nil {
		return json.Marshal(&src.FileMetadataResponse)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableKnowledgeResponseFilesInner struct {
	value *KnowledgeResponseFilesInner
	isSet bool
}

func (v NullableKnowledgeResponseFilesInner) Get() *KnowledgeResponseFilesInner {
	return v.value
}

func (v *NullableKnowledgeResponseFilesInner) Set(val *KnowledgeResponseFilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableKnowledgeResponseFilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableKnowledgeResponseFilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKnowledgeResponseFilesInner(val *KnowledgeResponseFilesInner) *NullableKnowledgeResponseFilesInner {
	return &NullableKnowledgeResponseFilesInner{value: val, isSet: true}
}

func (v NullableKnowledgeResponseFilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKnowledgeResponseFilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


