package xurrent

import (
	"context"
	"net/http"
	"net/url"
)

// GetSearchJSON performs GET /v1/search with the given query parameters.
//
// Per https://developer.xurrent.com/v1/search/, result continuation uses the **page**
// query parameter (token from the **x-pagination-next-page** response header — see
// [SearchNextPageToken] and [WithSearchPage]), not **search_after** / **search_before** like
// most collection endpoints. For multi-page walks, use [APIClient.ForEachSearchPage].
// The generated [SearchAPIService.GetSearch] request builder still exposes **SearchAfter**
// for parity with the OpenAPI file; prefer this helper (or [APIClient.GetCollectionJSON]
// with path "/v1/search") when you need **page=** explicitly.
func (c *APIClient) GetSearchJSON(ctx context.Context, query url.Values) ([]byte, *http.Response, error) {
	return c.GetCollectionJSON(ctx, "/v1/search", query)
}
