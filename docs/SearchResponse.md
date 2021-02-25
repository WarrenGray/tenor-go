# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | **string** | a position identifier to use with the next API query to retrieve the next set of results, or null if there are no further results.  | 
**Results** | [**[]GifObject**](GifObject.md) | an array of GifObjects containing the most relevant GIFs for the requested search term - Sorted by relevancy Rank  | 

## Methods

### NewSearchResponse

`func NewSearchResponse(next string, results []GifObject, ) *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *SearchResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SearchResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SearchResponse) SetNext(v string)`

SetNext sets Next field to given value.


### GetResults

`func (o *SearchResponse) GetResults() []GifObject`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchResponse) GetResultsOk() (*[]GifObject, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchResponse) SetResults(v []GifObject)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


