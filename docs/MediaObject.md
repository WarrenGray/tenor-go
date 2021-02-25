# MediaObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preview** | **string** | a url to a preview image of the media source | 
**Url** | **string** | a url to the media source | 
**Dims** | **[]int32** | width and height in pixels | 
**Size** | **int32** | size of file in bytes | 

## Methods

### NewMediaObject

`func NewMediaObject(preview string, url string, dims []int32, size int32, ) *MediaObject`

NewMediaObject instantiates a new MediaObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaObjectWithDefaults

`func NewMediaObjectWithDefaults() *MediaObject`

NewMediaObjectWithDefaults instantiates a new MediaObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreview

`func (o *MediaObject) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *MediaObject) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *MediaObject) SetPreview(v string)`

SetPreview sets Preview field to given value.


### GetUrl

`func (o *MediaObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MediaObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MediaObject) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDims

`func (o *MediaObject) GetDims() []int32`

GetDims returns the Dims field if non-nil, zero value otherwise.

### GetDimsOk

`func (o *MediaObject) GetDimsOk() (*[]int32, bool)`

GetDimsOk returns a tuple with the Dims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDims

`func (o *MediaObject) SetDims(v []int32)`

SetDims sets Dims field to given value.


### GetSize

`func (o *MediaObject) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MediaObject) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MediaObject) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


