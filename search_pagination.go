package xurrent

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchNextPageHeader is the response header that carries the continuation token for
// [APIClient.GetSearchJSON] / GET /v1/search (see https://developer.xurrent.com/v1/search/).
const SearchNextPageHeader = "X-Pagination-Next-Page"

// DefaultSearchMaxPages is the default cap for [APIClient.ForEachSearchPage] when maxPages <= 0.
const DefaultSearchMaxPages = 50

// HardMaxSearchPages is the upper bound for maxPages passed to [APIClient.ForEachSearchPage].
const HardMaxSearchPages = 500

// SearchNextPageToken returns the opaque value for the **page** query parameter on the next
// GET /v1/search request. It reads [SearchNextPageHeader]. If the header looks like an
// absolute URL, a path-only URL, or contains a **page=** query parameter, that token is
// extracted; otherwise the trimmed header value is returned as-is.
func SearchNextPageToken(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	raw := strings.TrimSpace(resp.Header.Get(SearchNextPageHeader))
	if raw == "" {
		return ""
	}
	if tok := searchPageTokenFromURLLike(raw); tok != "" {
		return tok
	}
	return raw
}

func searchPageTokenFromURLLike(raw string) string {
	s := raw
	switch {
	case strings.HasPrefix(s, "/"):
		s = "https://api.xurrent.invalid" + s
	case !strings.Contains(s, "://"):
		return ""
	}
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}
	if p := strings.TrimSpace(u.Query().Get("page")); p != "" {
		return p
	}
	return ""
}

// WithSearchPage returns a copy of q with **page** set to pageToken, or **page** removed
// when pageToken is empty (same pattern as [WithSearchAfter] for collection pagination).
func WithSearchPage(q url.Values, pageToken string) url.Values {
	out := CloneURLValues(q)
	if pageToken == "" {
		out.Del("page")
		return out
	}
	out.Set("page", pageToken)
	return out
}

// SearchPageFunc is invoked for each successful GET /v1/search response (HTTP 2xx) inside
// [APIClient.ForEachSearchPage]. Return a non-nil error to stop iteration.
type SearchPageFunc func(body []byte, resp *http.Response) error

// ForEachSearchPage calls fn once per page of search results, following
// [SearchNextPageToken] until the header is absent or maxPages is reached.
// If maxPages <= 0, [DefaultSearchMaxPages] is used; values above [HardMaxSearchPages] are clamped.
// If fn returns an error, iteration stops and that error is returned.
// If more pages exist after the last fetched page, returns a non-nil error describing truncation.
func (c *APIClient) ForEachSearchPage(ctx context.Context, initial url.Values, maxPages int, fn SearchPageFunc) error {
	if maxPages <= 0 {
		maxPages = DefaultSearchMaxPages
	}
	if maxPages > HardMaxSearchPages {
		maxPages = HardMaxSearchPages
	}
	q := CloneURLValues(initial)
	for page := 0; page < maxPages; page++ {
		body, resp, err := c.GetSearchJSON(ctx, q)
		if err != nil {
			return err
		}
		if resp != nil && (resp.StatusCode < 200 || resp.StatusCode >= 300) {
			return fmt.Errorf("search: HTTP %d", resp.StatusCode)
		}
		if err := fn(body, resp); err != nil {
			return err
		}
		tok := SearchNextPageToken(resp)
		if tok == "" {
			return nil
		}
		if page == maxPages-1 {
			return fmt.Errorf("search: max pages (%d) reached; more results available", maxPages)
		}
		q = WithSearchPage(q, tok)
	}
	// The loop always returns inside; this satisfies the compiler.
	return nil
}
