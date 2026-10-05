package xurrent

import (
	"net/url"
	"strconv"
	"strings"
)

// CollectionFilterValues builds url.Values for Xurrent collection filters.
// Keys are API field names; values are filter expressions as documented at
// https://developer.xurrent.com/v1/general/filtering/ (including operators like
// "!", ">", "<=", comma-separated IN lists, etc.).
func CollectionFilterValues(filters map[string]string) url.Values {
	v := make(url.Values)
	for k, val := range filters {
		v.Set(k, val)
	}
	return v
}

// EncodeQuery returns the encoded query string (without leading '?') suitable for
// attaching to a path when constructing requests outside the generated client.
func EncodeQuery(v url.Values) string {
	return v.Encode()
}

// MergeQuery combines two encoded query strings (e.g. generated client base query
// plus extra filter query from [CollectionFilterValues]).
func MergeQuery(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	return a + "&" + b
}

// FieldsParam returns a single "fields" query value from field names (comma-separated).
func FieldsParam(names ...string) string {
	return strings.Join(names, ",")
}

// ListCollectionQuery builds common collection GET query parameters: per_page, fields,
// and optional filters (e.g. name, state, serial_nr). Pass nil or empty filters for none.
func ListCollectionQuery(perPage int32, fields string, filters map[string]string) url.Values {
	v := CollectionFilterValues(filters)
	if perPage > 0 {
		v.Set("per_page", strconv.FormatInt(int64(perPage), 10))
	}
	if fields != "" {
		v.Set("fields", fields)
	}
	return v
}

// CloneURLValues returns a copy of q. A nil receiver yields an empty url.Values.
func CloneURLValues(q url.Values) url.Values {
	if len(q) == 0 {
		return url.Values{}
	}
	out := make(url.Values, len(q))
	for k, vals := range q {
		out[k] = append([]string(nil), vals...)
	}
	return out
}

// WithSearchAfter returns a copy of q with search_after set, or q unchanged if cursor is empty.
func WithSearchAfter(q url.Values, searchAfter string) url.Values {
	if searchAfter == "" {
		return q
	}
	out := CloneURLValues(q)
	out.Set("search_after", searchAfter)
	return out
}
