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
	"bytes"
	"fmt"
)

// checks if the GetStatus200ResponseServerResets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatus200ResponseServerResets{}

// GetStatus200ResponseServerResets struct for GetStatus200ResponseServerResets
type GetStatus200ResponseServerResets struct {
	// The date and time when the game server will reset.
	Next string `json:"next"`
	// How often we intend to reset the game server.
	Frequency string `json:"frequency"`
}

type _GetStatus200ResponseServerResets GetStatus200ResponseServerResets

// NewGetStatus200ResponseServerResets instantiates a new GetStatus200ResponseServerResets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus200ResponseServerResets(next string, frequency string) *GetStatus200ResponseServerResets {
	this := GetStatus200ResponseServerResets{}
	this.Next = next
	this.Frequency = frequency
	return &this
}

// NewGetStatus200ResponseServerResetsWithDefaults instantiates a new GetStatus200ResponseServerResets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatus200ResponseServerResetsWithDefaults() *GetStatus200ResponseServerResets {
	this := GetStatus200ResponseServerResets{}
	return &this
}

// GetNext returns the Next field value
func (o *GetStatus200ResponseServerResets) GetNext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseServerResets) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *GetStatus200ResponseServerResets) SetNext(v string) {
	o.Next = v
}

// GetFrequency returns the Frequency field value
func (o *GetStatus200ResponseServerResets) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseServerResets) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *GetStatus200ResponseServerResets) SetFrequency(v string) {
	o.Frequency = v
}

func (o GetStatus200ResponseServerResets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatus200ResponseServerResets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["next"] = o.Next
	toSerialize["frequency"] = o.Frequency
	return toSerialize, nil
}

func (o *GetStatus200ResponseServerResets) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"next",
		"frequency",
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

	varGetStatus200ResponseServerResets := _GetStatus200ResponseServerResets{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStatus200ResponseServerResets)

	if err != nil {
		return err
	}

	*o = GetStatus200ResponseServerResets(varGetStatus200ResponseServerResets)

	return err
}

type NullableGetStatus200ResponseServerResets struct {
	value *GetStatus200ResponseServerResets
	isSet bool
}

func (v NullableGetStatus200ResponseServerResets) Get() *GetStatus200ResponseServerResets {
	return v.value
}

func (v *NullableGetStatus200ResponseServerResets) Set(val *GetStatus200ResponseServerResets) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus200ResponseServerResets) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus200ResponseServerResets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus200ResponseServerResets(val *GetStatus200ResponseServerResets) *NullableGetStatus200ResponseServerResets {
	return &NullableGetStatus200ResponseServerResets{value: val, isSet: true}
}

func (v NullableGetStatus200ResponseServerResets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus200ResponseServerResets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


