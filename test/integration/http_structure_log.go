//go:build integration

package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

// redactingTransport logs JSON **structure only** (skeleton) for responses and minimal
// request metadata. It never logs Authorization, tokens, X-4me-Account, or raw query values.
type redactingTransport struct {
	base http.RoundTripper
	mu   sync.Mutex
	w    io.Writer
}

var digitsOnly = regexp.MustCompile(`^\d+$`)

// newRedactingTransport wraps base (typically http.DefaultTransport). Each request/response
// pair appends one NDJSON line to w.
func newRedactingTransport(base http.RoundTripper, w io.Writer) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &redactingTransport{base: base, w: w}
}

func normalizeOpenAPIPath(path string) string {
	segs := strings.Split(strings.Trim(path, "/"), "/")
	var out []string
	for _, s := range segs {
		if s == "" {
			continue
		}
		if digitsOnly.MatchString(s) {
			out = append(out, "{id}")
			continue
		}
		out = append(out, s)
	}
	return "/" + strings.Join(out, "/")
}

func queryParamNames(u *url.URL) []string {
	if u == nil || u.RawQuery == "" {
		return nil
	}
	q := u.Query()
	names := make([]string, 0, len(q))
	for k := range q {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

// jsonSkeleton returns a JSON-serializable tree with all leaf values replaced by
// empty / zero placeholders so nothing tenant-specific is retained.
func jsonSkeleton(v interface{}) interface{} {
	switch x := v.(type) {
	case map[string]interface{}:
		out := make(map[string]interface{}, len(x))
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			out[k] = jsonSkeleton(x[k])
		}
		return out
	case []interface{}:
		if len(x) == 0 {
			return []interface{}{}
		}
		// One representative element preserves field names without leaking cardinality beyond "non-empty".
		return []interface{}{jsonSkeleton(x[0])}
	case string:
		return ""
	case float64:
		return 0
	case json.Number:
		return 0
	case bool:
		return false
	case nil:
		return nil
	default:
		return map[string]string{"_type": fmt.Sprintf("%T", x)}
	}
}

const maxBodyLogBytes = 2 << 20 // 2 MiB

func (rt *redactingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now().UTC().Format(time.RFC3339Nano)
	path := ""
	if req.URL != nil {
		path = req.URL.Path
	}
	norm := normalizeOpenAPIPath(path)
	qnames := queryParamNames(req.URL)

	reqLine := map[string]interface{}{
		"ts":                start,
		"kind":              "request",
		"method":            req.Method,
		"path":              norm,
		"query_param_names": qnames,
	}
	rt.appendJSON(reqLine)

	resp, err := rt.base.RoundTrip(req)
	if err != nil {
		rt.appendJSON(map[string]interface{}{
			"ts":              time.Now().UTC().Format(time.RFC3339Nano),
			"kind":            "transport_error",
			"path":            norm,
			"error_structure": "RoundTrip failed (details not logged)",
		})
		return resp, err
	}
	if resp == nil {
		return resp, nil
	}

	body, readErr := io.ReadAll(io.LimitReader(resp.Body, maxBodyLogBytes))
	_ = resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewReader(body))

	ct := resp.Header.Get("Content-Type")
	entry := map[string]interface{}{
		"ts":           time.Now().UTC().Format(time.RFC3339Nano),
		"kind":         "response",
		"method":       req.Method,
		"path":         norm,
		"status":       resp.StatusCode,
		"content_type": ct,
	}

	if readErr != nil {
		entry["body_error"] = "read failed"
		rt.appendJSON(entry)
		return resp, nil
	}

	if !strings.Contains(strings.ToLower(ct), "application/json") {
		entry["json_skeleton"] = map[string]interface{}{
			"_non_json": true,
			"_bytes":    len(body),
		}
		rt.appendJSON(entry)
		return resp, nil
	}

	var parsed interface{}
	if err := json.Unmarshal(body, &parsed); err != nil {
		entry["json_skeleton"] = map[string]interface{}{
			"_parse_error": true,
			"_bytes":       len(body),
		}
	} else {
		entry["json_skeleton"] = jsonSkeleton(parsed)
	}
	rt.appendJSON(entry)
	return resp, nil
}

func (rt *redactingTransport) appendJSON(v interface{}) {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	_, _ = rt.w.Write(append(b, '\n'))
}
