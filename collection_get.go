package xurrent

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// GetCollectionJSON performs GET on path (e.g. "/v1/teams") with query parameters.
// Use together with [CollectionFilterValues], [ListCollectionQuery], and [EncodeQuery]
// for filters from https://developer.xurrent.com/v1/general/filtering/ when those
// filters are not modeled as named parameters on generated request types.
//
// Hand-maintained (see .openapi-generator-ignore); do not add similar methods by editing
// generated client.go — extend collection_filter.go / collection_get.go or another ignored file.
func (c *APIClient) GetCollectionJSON(ctx context.Context, path string, query url.Values) ([]byte, *http.Response, error) {
	localBasePath, err := c.cfg.ServerURLWithContext(ctx, "APIClient.GetCollectionJSON")
	if err != nil {
		return nil, nil, err
	}
	localVarPath := localBasePath + path
	localVarHeaderParams := make(map[string]string)
	localVarHTTPHeaderAccepts := []string{"application/json"}
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(ctx, localVarPath, http.MethodGet, nil, localVarHeaderParams, query, url.Values{}, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := c.callAPI(req)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}
	return body, resp, nil
}
