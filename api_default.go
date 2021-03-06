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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiRandomRequest struct {
	ctx           _context.Context
	ApiService    *DefaultApiService
	q             *string
	key           *string
	locale        *string
	contentfilter *string
	mediaFilter   *string
	arRange       *string
	limit         *int32
	pos           *string
	anonId        *string
}

func (r ApiRandomRequest) Q(q string) ApiRandomRequest {
	r.q = &q
	return r
}
func (r ApiRandomRequest) Key(key string) ApiRandomRequest {
	r.key = &key
	return r
}
func (r ApiRandomRequest) Locale(locale string) ApiRandomRequest {
	r.locale = &locale
	return r
}
func (r ApiRandomRequest) Contentfilter(contentfilter string) ApiRandomRequest {
	r.contentfilter = &contentfilter
	return r
}
func (r ApiRandomRequest) MediaFilter(mediaFilter string) ApiRandomRequest {
	r.mediaFilter = &mediaFilter
	return r
}
func (r ApiRandomRequest) ArRange(arRange string) ApiRandomRequest {
	r.arRange = &arRange
	return r
}
func (r ApiRandomRequest) Limit(limit int32) ApiRandomRequest {
	r.limit = &limit
	return r
}
func (r ApiRandomRequest) Pos(pos string) ApiRandomRequest {
	r.pos = &pos
	return r
}
func (r ApiRandomRequest) AnonId(anonId string) ApiRandomRequest {
	r.anonId = &anonId
	return r
}

func (r ApiRandomRequest) Execute() (SearchResponse, *_nethttp.Response, error) {
	return r.ApiService.RandomExecute(r)
}

/*
 * Random Method for Random
 * Get a randomized list of GIFs for a given search term. This differs from the search endpoint which returns a rank ordered list of GIFs for a given search term.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRandomRequest
 */
func (a *DefaultApiService) Random(ctx _context.Context) ApiRandomRequest {
	return ApiRandomRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SearchResponse
 */
func (a *DefaultApiService) RandomExecute(r ApiRandomRequest) (SearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.Random")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/random"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.q == nil {
		return localVarReturnValue, nil, reportError("q is required and must be specified")
	}
	if r.key == nil {
		return localVarReturnValue, nil, reportError("key is required and must be specified")
	}

	localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	localVarQueryParams.Add("key", parameterToString(*r.key, ""))
	if r.locale != nil {
		localVarQueryParams.Add("locale", parameterToString(*r.locale, ""))
	}
	if r.contentfilter != nil {
		localVarQueryParams.Add("contentfilter", parameterToString(*r.contentfilter, ""))
	}
	if r.mediaFilter != nil {
		localVarQueryParams.Add("media_filter", parameterToString(*r.mediaFilter, ""))
	}
	if r.arRange != nil {
		localVarQueryParams.Add("ar_range", parameterToString(*r.arRange, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.pos != nil {
		localVarQueryParams.Add("pos", parameterToString(*r.pos, ""))
	}
	if r.anonId != nil {
		localVarQueryParams.Add("anon_id", parameterToString(*r.anonId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchRequest struct {
	ctx           _context.Context
	ApiService    *DefaultApiService
	q             *string
	key           *string
	locale        *string
	contentfilter *string
	mediaFilter   *string
	arRange       *string
	limit         *int32
	pos           *string
	anonId        *string
}

func (r ApiSearchRequest) Q(q string) ApiSearchRequest {
	r.q = &q
	return r
}
func (r ApiSearchRequest) Key(key string) ApiSearchRequest {
	r.key = &key
	return r
}
func (r ApiSearchRequest) Locale(locale string) ApiSearchRequest {
	r.locale = &locale
	return r
}
func (r ApiSearchRequest) Contentfilter(contentfilter string) ApiSearchRequest {
	r.contentfilter = &contentfilter
	return r
}
func (r ApiSearchRequest) MediaFilter(mediaFilter string) ApiSearchRequest {
	r.mediaFilter = &mediaFilter
	return r
}
func (r ApiSearchRequest) ArRange(arRange string) ApiSearchRequest {
	r.arRange = &arRange
	return r
}
func (r ApiSearchRequest) Limit(limit int32) ApiSearchRequest {
	r.limit = &limit
	return r
}
func (r ApiSearchRequest) Pos(pos string) ApiSearchRequest {
	r.pos = &pos
	return r
}
func (r ApiSearchRequest) AnonId(anonId string) ApiSearchRequest {
	r.anonId = &anonId
	return r
}

func (r ApiSearchRequest) Execute() (SearchResponse, *_nethttp.Response, error) {
	return r.ApiService.SearchExecute(r)
}

/*
 * Search Method for Search
 * Get a json object containing a list of the most relevant GIFs for a given search term(s), category(ies), emoji(s), or any combination thereof.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSearchRequest
 */
func (a *DefaultApiService) Search(ctx _context.Context) ApiSearchRequest {
	return ApiSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SearchResponse
 */
func (a *DefaultApiService) SearchExecute(r ApiSearchRequest) (SearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.Search")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.q == nil {
		return localVarReturnValue, nil, reportError("q is required and must be specified")
	}
	if r.key == nil {
		return localVarReturnValue, nil, reportError("key is required and must be specified")
	}

	localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	localVarQueryParams.Add("key", parameterToString(*r.key, ""))
	if r.locale != nil {
		localVarQueryParams.Add("locale", parameterToString(*r.locale, ""))
	}
	if r.contentfilter != nil {
		localVarQueryParams.Add("contentfilter", parameterToString(*r.contentfilter, ""))
	}
	if r.mediaFilter != nil {
		localVarQueryParams.Add("media_filter", parameterToString(*r.mediaFilter, ""))
	}
	if r.arRange != nil {
		localVarQueryParams.Add("ar_range", parameterToString(*r.arRange, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.pos != nil {
		localVarQueryParams.Add("pos", parameterToString(*r.pos, ""))
	}
	if r.anonId != nil {
		localVarQueryParams.Add("anon_id", parameterToString(*r.anonId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
