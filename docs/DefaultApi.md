# \DefaultApi

All URIs are relative to *https://g.tenor.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Random**](DefaultApi.md#Random) | **Get** /v1/random | 
[**Search**](DefaultApi.md#Search) | **Get** /v1/search | 



## Random

> SearchResponse Random(ctx).Q(q).Key(key).Locale(locale).Contentfilter(contentfilter).MediaFilter(mediaFilter).ArRange(arRange).Limit(limit).Pos(pos).AnonId(anonId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    q := "q_example" // string | a search string
    key := "key_example" // string | client key for privileged API access
    locale := "locale_example" // string | specify default language to interpret search string; xx is ISO 639-1 language code, _YY (optional) is 2-letter ISO 3166-1 country code  (optional)
    contentfilter := "contentfilter_example" // string | (values: off | low | medium | high) specify the content safety filter level  (optional)
    mediaFilter := "mediaFilter_example" // string | (values: basic | minimal) Reduce the Number of GIF formats returned in the GIF_OBJECT list.  minimal - tinygif, gif, and mp4.  basic - nanomp4, tinygif, tinymp4, gif, mp4, and nanogif  (optional)
    arRange := "arRange_example" // string | (values: all | wide | standard ) Filter the response GIF_OBJECT list to only include GIFs with aspect ratios that fit with in the selected range.  all - no constraints  wide - 0.42 <= aspect ratio <= 2.36  standard - .56 <= aspect ratio <= 1.78  (optional)
    limit := int32(56) // int32 | fetch up to a specified number of results (max: 50).  (optional)
    pos := "pos_example" // string | get results starting at position \"value\". Use a non-zero \"next\" value returned by API results to get the next set of results. pos is not an index and may be an integer, float, or string  (optional)
    anonId := "anonId_example" // string | specify the anonymous_id tied to the given user (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Random(context.Background()).Q(q).Key(key).Locale(locale).Contentfilter(contentfilter).MediaFilter(mediaFilter).ArRange(arRange).Limit(limit).Pos(pos).AnonId(anonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Random``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Random`: SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Random`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRandomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | a search string | 
 **key** | **string** | client key for privileged API access | 
 **locale** | **string** | specify default language to interpret search string; xx is ISO 639-1 language code, _YY (optional) is 2-letter ISO 3166-1 country code  | 
 **contentfilter** | **string** | (values: off | low | medium | high) specify the content safety filter level  | 
 **mediaFilter** | **string** | (values: basic | minimal) Reduce the Number of GIF formats returned in the GIF_OBJECT list.  minimal - tinygif, gif, and mp4.  basic - nanomp4, tinygif, tinymp4, gif, mp4, and nanogif  | 
 **arRange** | **string** | (values: all | wide | standard ) Filter the response GIF_OBJECT list to only include GIFs with aspect ratios that fit with in the selected range.  all - no constraints  wide - 0.42 &lt;&#x3D; aspect ratio &lt;&#x3D; 2.36  standard - .56 &lt;&#x3D; aspect ratio &lt;&#x3D; 1.78  | 
 **limit** | **int32** | fetch up to a specified number of results (max: 50).  | 
 **pos** | **string** | get results starting at position \&quot;value\&quot;. Use a non-zero \&quot;next\&quot; value returned by API results to get the next set of results. pos is not an index and may be an integer, float, or string  | 
 **anonId** | **string** | specify the anonymous_id tied to the given user | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SearchResponse Search(ctx).Q(q).Key(key).Locale(locale).Contentfilter(contentfilter).MediaFilter(mediaFilter).ArRange(arRange).Limit(limit).Pos(pos).AnonId(anonId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    q := "q_example" // string | a search string
    key := "key_example" // string | client key for privileged API access
    locale := "locale_example" // string | specify default language to interpret search string; xx is ISO 639-1 language code, _YY (optional) is 2-letter ISO 3166-1 country code  (optional)
    contentfilter := "contentfilter_example" // string | (values: off | low | medium | high) specify the content safety filter level  (optional)
    mediaFilter := "mediaFilter_example" // string | (values: basic | minimal) Reduce the Number of GIF formats returned in the GIF_OBJECT list.  minimal - tinygif, gif, and mp4.  basic - nanomp4, tinygif, tinymp4, gif, mp4, and nanogif  (optional)
    arRange := "arRange_example" // string | (values: all | wide | standard ) Filter the response GIF_OBJECT list to only include GIFs with aspect ratios that fit with in the selected range.  all - no constraints  wide - 0.42 <= aspect ratio <= 2.36  standard - .56 <= aspect ratio <= 1.78  (optional)
    limit := int32(56) // int32 | fetch up to a specified number of results (max: 50).  (optional)
    pos := "pos_example" // string | get results starting at position \"value\". Use a non-zero \"next\" value returned by API results to get the next set of results. pos is not an index and may be an integer, float, or string  (optional)
    anonId := "anonId_example" // string | specify the anonymous_id tied to the given user (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Search(context.Background()).Q(q).Key(key).Locale(locale).Contentfilter(contentfilter).MediaFilter(mediaFilter).ArRange(arRange).Limit(limit).Pos(pos).AnonId(anonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | a search string | 
 **key** | **string** | client key for privileged API access | 
 **locale** | **string** | specify default language to interpret search string; xx is ISO 639-1 language code, _YY (optional) is 2-letter ISO 3166-1 country code  | 
 **contentfilter** | **string** | (values: off | low | medium | high) specify the content safety filter level  | 
 **mediaFilter** | **string** | (values: basic | minimal) Reduce the Number of GIF formats returned in the GIF_OBJECT list.  minimal - tinygif, gif, and mp4.  basic - nanomp4, tinygif, tinymp4, gif, mp4, and nanogif  | 
 **arRange** | **string** | (values: all | wide | standard ) Filter the response GIF_OBJECT list to only include GIFs with aspect ratios that fit with in the selected range.  all - no constraints  wide - 0.42 &lt;&#x3D; aspect ratio &lt;&#x3D; 2.36  standard - .56 &lt;&#x3D; aspect ratio &lt;&#x3D; 1.78  | 
 **limit** | **int32** | fetch up to a specified number of results (max: 50).  | 
 **pos** | **string** | get results starting at position \&quot;value\&quot;. Use a non-zero \&quot;next\&quot; value returned by API results to get the next set of results. pos is not an index and may be an integer, float, or string  | 
 **anonId** | **string** | specify the anonymous_id tied to the given user | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

