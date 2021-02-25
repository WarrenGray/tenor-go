# GifObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **float32** | a unix timestamp representing when this post was created. | 
**Hasaudio** | **bool** | true if this post contains audio (only video formats support audio, the gif image file format can not contain audio information).  | 
**Id** | **string** | Tenor result identifier | 
**Media** | [**[]Media**](Media.md) | An array of dictionaries with GIF_FORMAT as the key and MEDIA_OBJECT as the value | 
**Tags** | **[]string** | an array of tags for the post | 
**Title** | **string** |  | 

## Methods

### NewGifObject

`func NewGifObject(created float32, hasaudio bool, id string, media []Media, tags []string, title string, ) *GifObject`

NewGifObject instantiates a new GifObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGifObjectWithDefaults

`func NewGifObjectWithDefaults() *GifObject`

NewGifObjectWithDefaults instantiates a new GifObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GifObject) GetCreated() float32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GifObject) GetCreatedOk() (*float32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GifObject) SetCreated(v float32)`

SetCreated sets Created field to given value.


### GetHasaudio

`func (o *GifObject) GetHasaudio() bool`

GetHasaudio returns the Hasaudio field if non-nil, zero value otherwise.

### GetHasaudioOk

`func (o *GifObject) GetHasaudioOk() (*bool, bool)`

GetHasaudioOk returns a tuple with the Hasaudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasaudio

`func (o *GifObject) SetHasaudio(v bool)`

SetHasaudio sets Hasaudio field to given value.


### GetId

`func (o *GifObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GifObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GifObject) SetId(v string)`

SetId sets Id field to given value.


### GetMedia

`func (o *GifObject) GetMedia() []Media`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *GifObject) GetMediaOk() (*[]Media, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *GifObject) SetMedia(v []Media)`

SetMedia sets Media field to given value.


### GetTags

`func (o *GifObject) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GifObject) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GifObject) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTitle

`func (o *GifObject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GifObject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GifObject) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


