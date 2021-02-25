/*
 * Tenor GifAPI Client
 *
 * API client for https://tenor.com/gifapi/
 *
 * API version: 0.1
 * Contact: warren@warren-gray.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SearchResponse struct for SearchResponse
type SearchResponse struct {
	// a position identifier to use with the next API query to retrieve the next set of results, or null if there are no further results.
	Next string `json:"next"`
	// an array of GifObjects containing the most relevant GIFs for the requested search term - Sorted by relevancy Rank
	Results []GifObject `json:"results"`
}

// NewSearchResponse instantiates a new SearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResponse(next string, results []GifObject) *SearchResponse {
	this := SearchResponse{}
	this.Next = next
	this.Results = results
	return &this
}

// NewSearchResponseWithDefaults instantiates a new SearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResponseWithDefaults() *SearchResponse {
	this := SearchResponse{}
	return &this
}

// GetNext returns the Next field value
func (o *SearchResponse) GetNext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *SearchResponse) SetNext(v string) {
	o.Next = v
}

// GetResults returns the Results field value
func (o *SearchResponse) GetResults() []GifObject {
	if o == nil {
		var ret []GifObject
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetResultsOk() (*[]GifObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *SearchResponse) SetResults(v []GifObject) {
	o.Results = v
}

func (o SearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["next"] = o.Next
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableSearchResponse struct {
	value *SearchResponse
	isSet bool
}

func (v NullableSearchResponse) Get() *SearchResponse {
	return v.value
}

func (v *NullableSearchResponse) Set(val *SearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResponse(val *SearchResponse) *NullableSearchResponse {
	return &NullableSearchResponse{value: val, isSet: true}
}

func (v NullableSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}