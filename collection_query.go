package xurrent

import (
	"net/url"
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

// EncodeQuery appends encoded query string (without leading '?') suitable for
// attaching to a path when constructing requests outside the generated client.
func EncodeQuery(v url.Values) string {
	return v.Encode()
}

// MergeQuery combines two encoded query strings (e.g. generated client base query
// plus extra filter query from CollectionFilterValues).
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
