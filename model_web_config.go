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

// checks if the WebConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebConfig{}

// WebConfig struct for WebConfig
type WebConfig struct {
	Search WebSearchConfig `json:"search"`
	ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION NullableBool `json:"ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION,omitempty"`
	BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL NullableBool `json:"BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL,omitempty"`
}

type _WebConfig WebConfig

// NewWebConfig instantiates a new WebConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebConfig(search WebSearchConfig) *WebConfig {
	this := WebConfig{}
	this.Search = search
	return &this
}

// NewWebConfigWithDefaults instantiates a new WebConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebConfigWithDefaults() *WebConfig {
	this := WebConfig{}
	return &this
}

// GetSearch returns the Search field value
func (o *WebConfig) GetSearch() WebSearchConfig {
	if o == nil {
		var ret WebSearchConfig
		return ret
	}

	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *WebConfig) GetSearchOk() (*WebSearchConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value
func (o *WebConfig) SetSearch(v WebSearchConfig) {
	o.Search = v
}

// GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION returns the ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebConfig) GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION() bool {
	if o == nil || IsNil(o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.Get()) {
		var ret bool
		return ret
	}
	return *o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.Get()
}

// GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONOk returns a tuple with the ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebConfig) GetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.Get(), o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.IsSet()
}

// HasENABLE_RAG_WEB_LOADER_SSL_VERIFICATION returns a boolean if a field has been set.
func (o *WebConfig) HasENABLE_RAG_WEB_LOADER_SSL_VERIFICATION() bool {
	if o != nil && o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.IsSet() {
		return true
	}

	return false
}

// SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION gets a reference to the given NullableBool and assigns it to the ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION field.
func (o *WebConfig) SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION(v bool) {
	o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.Set(&v)
}
// SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONNil sets the value for ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION to be an explicit nil
func (o *WebConfig) SetENABLE_RAG_WEB_LOADER_SSL_VERIFICATIONNil() {
	o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.Set(nil)
}

// UnsetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION ensures that no value is present for ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION, not even an explicit nil
func (o *WebConfig) UnsetENABLE_RAG_WEB_LOADER_SSL_VERIFICATION() {
	o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.Unset()
}

// GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL returns the BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebConfig) GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL() bool {
	if o == nil || IsNil(o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.Get()) {
		var ret bool
		return ret
	}
	return *o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.Get()
}

// GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALOk returns a tuple with the BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebConfig) GetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.Get(), o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.IsSet()
}

// HasBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL returns a boolean if a field has been set.
func (o *WebConfig) HasBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL() bool {
	if o != nil && o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.IsSet() {
		return true
	}

	return false
}

// SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL gets a reference to the given NullableBool and assigns it to the BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL field.
func (o *WebConfig) SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL(v bool) {
	o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.Set(&v)
}
// SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALNil sets the value for BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL to be an explicit nil
func (o *WebConfig) SetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVALNil() {
	o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.Set(nil)
}

// UnsetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL ensures that no value is present for BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL, not even an explicit nil
func (o *WebConfig) UnsetBYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL() {
	o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.Unset()
}

func (o WebConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["search"] = o.Search
	if o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.IsSet() {
		toSerialize["ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION"] = o.ENABLE_RAG_WEB_LOADER_SSL_VERIFICATION.Get()
	}
	if o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.IsSet() {
		toSerialize["BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL"] = o.BYPASS_WEB_SEARCH_EMBEDDING_AND_RETRIEVAL.Get()
	}
	return toSerialize, nil
}

func (o *WebConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"search",
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

	varWebConfig := _WebConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebConfig)

	if err != nil {
		return err
	}

	*o = WebConfig(varWebConfig)

	return err
}

type NullableWebConfig struct {
	value *WebConfig
	isSet bool
}

func (v NullableWebConfig) Get() *WebConfig {
	return v.value
}

func (v *NullableWebConfig) Set(val *WebConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableWebConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableWebConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebConfig(val *WebConfig) *NullableWebConfig {
	return &NullableWebConfig{value: val, isSet: true}
}

func (v NullableWebConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


