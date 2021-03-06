/*
 * Tenor GifAPI Client
 *
 * API client for https://tenor.com/gifapi/
 *
 * API version: 0.1
 * Contact: warren@warren-gray.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tenor

import (
	"encoding/json"
)

// Media struct for Media
type Media struct {
	Gif       *MediaObject `json:"gif,omitempty"`
	Mediumgif *MediaObject `json:"mediumgif,omitempty"`
	Tinygif   *MediaObject `json:"tinygif,omitempty"`
	Nanogif   *MediaObject `json:"nanogif,omitempty"`
	Mp4       *MediaObject `json:"mp4,omitempty"`
	Loopedmp4 *MediaObject `json:"loopedmp4,omitempty"`
	Tinymp4   *MediaObject `json:"tinymp4,omitempty"`
	Nanomp4   *MediaObject `json:"nanomp4,omitempty"`
	Webm      *MediaObject `json:"webm,omitempty"`
	Tinywebm  *MediaObject `json:"tinywebm,omitempty"`
	Nanowebm  *MediaObject `json:"nanowebm,omitempty"`
}

// NewMedia instantiates a new Media object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedia() *Media {
	this := Media{}
	return &this
}

// NewMediaWithDefaults instantiates a new Media object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaWithDefaults() *Media {
	this := Media{}
	return &this
}

// GetGif returns the Gif field value if set, zero value otherwise.
func (o *Media) GetGif() MediaObject {
	if o == nil || o.Gif == nil {
		var ret MediaObject
		return ret
	}
	return *o.Gif
}

// GetGifOk returns a tuple with the Gif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetGifOk() (*MediaObject, bool) {
	if o == nil || o.Gif == nil {
		return nil, false
	}
	return o.Gif, true
}

// HasGif returns a boolean if a field has been set.
func (o *Media) HasGif() bool {
	if o != nil && o.Gif != nil {
		return true
	}

	return false
}

// SetGif gets a reference to the given MediaObject and assigns it to the Gif field.
func (o *Media) SetGif(v MediaObject) {
	o.Gif = &v
}

// GetMediumgif returns the Mediumgif field value if set, zero value otherwise.
func (o *Media) GetMediumgif() MediaObject {
	if o == nil || o.Mediumgif == nil {
		var ret MediaObject
		return ret
	}
	return *o.Mediumgif
}

// GetMediumgifOk returns a tuple with the Mediumgif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetMediumgifOk() (*MediaObject, bool) {
	if o == nil || o.Mediumgif == nil {
		return nil, false
	}
	return o.Mediumgif, true
}

// HasMediumgif returns a boolean if a field has been set.
func (o *Media) HasMediumgif() bool {
	if o != nil && o.Mediumgif != nil {
		return true
	}

	return false
}

// SetMediumgif gets a reference to the given MediaObject and assigns it to the Mediumgif field.
func (o *Media) SetMediumgif(v MediaObject) {
	o.Mediumgif = &v
}

// GetTinygif returns the Tinygif field value if set, zero value otherwise.
func (o *Media) GetTinygif() MediaObject {
	if o == nil || o.Tinygif == nil {
		var ret MediaObject
		return ret
	}
	return *o.Tinygif
}

// GetTinygifOk returns a tuple with the Tinygif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetTinygifOk() (*MediaObject, bool) {
	if o == nil || o.Tinygif == nil {
		return nil, false
	}
	return o.Tinygif, true
}

// HasTinygif returns a boolean if a field has been set.
func (o *Media) HasTinygif() bool {
	if o != nil && o.Tinygif != nil {
		return true
	}

	return false
}

// SetTinygif gets a reference to the given MediaObject and assigns it to the Tinygif field.
func (o *Media) SetTinygif(v MediaObject) {
	o.Tinygif = &v
}

// GetNanogif returns the Nanogif field value if set, zero value otherwise.
func (o *Media) GetNanogif() MediaObject {
	if o == nil || o.Nanogif == nil {
		var ret MediaObject
		return ret
	}
	return *o.Nanogif
}

// GetNanogifOk returns a tuple with the Nanogif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetNanogifOk() (*MediaObject, bool) {
	if o == nil || o.Nanogif == nil {
		return nil, false
	}
	return o.Nanogif, true
}

// HasNanogif returns a boolean if a field has been set.
func (o *Media) HasNanogif() bool {
	if o != nil && o.Nanogif != nil {
		return true
	}

	return false
}

// SetNanogif gets a reference to the given MediaObject and assigns it to the Nanogif field.
func (o *Media) SetNanogif(v MediaObject) {
	o.Nanogif = &v
}

// GetMp4 returns the Mp4 field value if set, zero value otherwise.
func (o *Media) GetMp4() MediaObject {
	if o == nil || o.Mp4 == nil {
		var ret MediaObject
		return ret
	}
	return *o.Mp4
}

// GetMp4Ok returns a tuple with the Mp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetMp4Ok() (*MediaObject, bool) {
	if o == nil || o.Mp4 == nil {
		return nil, false
	}
	return o.Mp4, true
}

// HasMp4 returns a boolean if a field has been set.
func (o *Media) HasMp4() bool {
	if o != nil && o.Mp4 != nil {
		return true
	}

	return false
}

// SetMp4 gets a reference to the given MediaObject and assigns it to the Mp4 field.
func (o *Media) SetMp4(v MediaObject) {
	o.Mp4 = &v
}

// GetLoopedmp4 returns the Loopedmp4 field value if set, zero value otherwise.
func (o *Media) GetLoopedmp4() MediaObject {
	if o == nil || o.Loopedmp4 == nil {
		var ret MediaObject
		return ret
	}
	return *o.Loopedmp4
}

// GetLoopedmp4Ok returns a tuple with the Loopedmp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetLoopedmp4Ok() (*MediaObject, bool) {
	if o == nil || o.Loopedmp4 == nil {
		return nil, false
	}
	return o.Loopedmp4, true
}

// HasLoopedmp4 returns a boolean if a field has been set.
func (o *Media) HasLoopedmp4() bool {
	if o != nil && o.Loopedmp4 != nil {
		return true
	}

	return false
}

// SetLoopedmp4 gets a reference to the given MediaObject and assigns it to the Loopedmp4 field.
func (o *Media) SetLoopedmp4(v MediaObject) {
	o.Loopedmp4 = &v
}

// GetTinymp4 returns the Tinymp4 field value if set, zero value otherwise.
func (o *Media) GetTinymp4() MediaObject {
	if o == nil || o.Tinymp4 == nil {
		var ret MediaObject
		return ret
	}
	return *o.Tinymp4
}

// GetTinymp4Ok returns a tuple with the Tinymp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetTinymp4Ok() (*MediaObject, bool) {
	if o == nil || o.Tinymp4 == nil {
		return nil, false
	}
	return o.Tinymp4, true
}

// HasTinymp4 returns a boolean if a field has been set.
func (o *Media) HasTinymp4() bool {
	if o != nil && o.Tinymp4 != nil {
		return true
	}

	return false
}

// SetTinymp4 gets a reference to the given MediaObject and assigns it to the Tinymp4 field.
func (o *Media) SetTinymp4(v MediaObject) {
	o.Tinymp4 = &v
}

// GetNanomp4 returns the Nanomp4 field value if set, zero value otherwise.
func (o *Media) GetNanomp4() MediaObject {
	if o == nil || o.Nanomp4 == nil {
		var ret MediaObject
		return ret
	}
	return *o.Nanomp4
}

// GetNanomp4Ok returns a tuple with the Nanomp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetNanomp4Ok() (*MediaObject, bool) {
	if o == nil || o.Nanomp4 == nil {
		return nil, false
	}
	return o.Nanomp4, true
}

// HasNanomp4 returns a boolean if a field has been set.
func (o *Media) HasNanomp4() bool {
	if o != nil && o.Nanomp4 != nil {
		return true
	}

	return false
}

// SetNanomp4 gets a reference to the given MediaObject and assigns it to the Nanomp4 field.
func (o *Media) SetNanomp4(v MediaObject) {
	o.Nanomp4 = &v
}

// GetWebm returns the Webm field value if set, zero value otherwise.
func (o *Media) GetWebm() MediaObject {
	if o == nil || o.Webm == nil {
		var ret MediaObject
		return ret
	}
	return *o.Webm
}

// GetWebmOk returns a tuple with the Webm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetWebmOk() (*MediaObject, bool) {
	if o == nil || o.Webm == nil {
		return nil, false
	}
	return o.Webm, true
}

// HasWebm returns a boolean if a field has been set.
func (o *Media) HasWebm() bool {
	if o != nil && o.Webm != nil {
		return true
	}

	return false
}

// SetWebm gets a reference to the given MediaObject and assigns it to the Webm field.
func (o *Media) SetWebm(v MediaObject) {
	o.Webm = &v
}

// GetTinywebm returns the Tinywebm field value if set, zero value otherwise.
func (o *Media) GetTinywebm() MediaObject {
	if o == nil || o.Tinywebm == nil {
		var ret MediaObject
		return ret
	}
	return *o.Tinywebm
}

// GetTinywebmOk returns a tuple with the Tinywebm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetTinywebmOk() (*MediaObject, bool) {
	if o == nil || o.Tinywebm == nil {
		return nil, false
	}
	return o.Tinywebm, true
}

// HasTinywebm returns a boolean if a field has been set.
func (o *Media) HasTinywebm() bool {
	if o != nil && o.Tinywebm != nil {
		return true
	}

	return false
}

// SetTinywebm gets a reference to the given MediaObject and assigns it to the Tinywebm field.
func (o *Media) SetTinywebm(v MediaObject) {
	o.Tinywebm = &v
}

// GetNanowebm returns the Nanowebm field value if set, zero value otherwise.
func (o *Media) GetNanowebm() MediaObject {
	if o == nil || o.Nanowebm == nil {
		var ret MediaObject
		return ret
	}
	return *o.Nanowebm
}

// GetNanowebmOk returns a tuple with the Nanowebm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Media) GetNanowebmOk() (*MediaObject, bool) {
	if o == nil || o.Nanowebm == nil {
		return nil, false
	}
	return o.Nanowebm, true
}

// HasNanowebm returns a boolean if a field has been set.
func (o *Media) HasNanowebm() bool {
	if o != nil && o.Nanowebm != nil {
		return true
	}

	return false
}

// SetNanowebm gets a reference to the given MediaObject and assigns it to the Nanowebm field.
func (o *Media) SetNanowebm(v MediaObject) {
	o.Nanowebm = &v
}

func (o Media) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Gif != nil {
		toSerialize["gif"] = o.Gif
	}
	if o.Mediumgif != nil {
		toSerialize["mediumgif"] = o.Mediumgif
	}
	if o.Tinygif != nil {
		toSerialize["tinygif"] = o.Tinygif
	}
	if o.Nanogif != nil {
		toSerialize["nanogif"] = o.Nanogif
	}
	if o.Mp4 != nil {
		toSerialize["mp4"] = o.Mp4
	}
	if o.Loopedmp4 != nil {
		toSerialize["loopedmp4"] = o.Loopedmp4
	}
	if o.Tinymp4 != nil {
		toSerialize["tinymp4"] = o.Tinymp4
	}
	if o.Nanomp4 != nil {
		toSerialize["nanomp4"] = o.Nanomp4
	}
	if o.Webm != nil {
		toSerialize["webm"] = o.Webm
	}
	if o.Tinywebm != nil {
		toSerialize["tinywebm"] = o.Tinywebm
	}
	if o.Nanowebm != nil {
		toSerialize["nanowebm"] = o.Nanowebm
	}
	return json.Marshal(toSerialize)
}

type NullableMedia struct {
	value *Media
	isSet bool
}

func (v NullableMedia) Get() *Media {
	return v.value
}

func (v *NullableMedia) Set(val *Media) {
	v.value = val
	v.isSet = true
}

func (v NullableMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedia(val *Media) *NullableMedia {
	return &NullableMedia{value: val, isSet: true}
}

func (v NullableMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
