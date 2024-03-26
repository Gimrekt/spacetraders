/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ExtractResourcesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtractResourcesRequest{}

// ExtractResourcesRequest struct for ExtractResourcesRequest
type ExtractResourcesRequest struct {
	Survey *Survey `json:"survey,omitempty"`
}

// NewExtractResourcesRequest instantiates a new ExtractResourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractResourcesRequest() *ExtractResourcesRequest {
	this := ExtractResourcesRequest{}
	return &this
}

// NewExtractResourcesRequestWithDefaults instantiates a new ExtractResourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractResourcesRequestWithDefaults() *ExtractResourcesRequest {
	this := ExtractResourcesRequest{}
	return &this
}

// GetSurvey returns the Survey field value if set, zero value otherwise.
func (o *ExtractResourcesRequest) GetSurvey() Survey {
	if o == nil || IsNil(o.Survey) {
		var ret Survey
		return ret
	}
	return *o.Survey
}

// GetSurveyOk returns a tuple with the Survey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractResourcesRequest) GetSurveyOk() (*Survey, bool) {
	if o == nil || IsNil(o.Survey) {
		return nil, false
	}
	return o.Survey, true
}

// HasSurvey returns a boolean if a field has been set.
func (o *ExtractResourcesRequest) HasSurvey() bool {
	if o != nil && !IsNil(o.Survey) {
		return true
	}

	return false
}

// SetSurvey gets a reference to the given Survey and assigns it to the Survey field.
func (o *ExtractResourcesRequest) SetSurvey(v Survey) {
	o.Survey = &v
}

func (o ExtractResourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtractResourcesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Survey) {
		toSerialize["survey"] = o.Survey
	}
	return toSerialize, nil
}

type NullableExtractResourcesRequest struct {
	value *ExtractResourcesRequest
	isSet bool
}

func (v NullableExtractResourcesRequest) Get() *ExtractResourcesRequest {
	return v.value
}

func (v *NullableExtractResourcesRequest) Set(val *ExtractResourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractResourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractResourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractResourcesRequest(val *ExtractResourcesRequest) *NullableExtractResourcesRequest {
	return &NullableExtractResourcesRequest{value: val, isSet: true}
}

func (v NullableExtractResourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractResourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


