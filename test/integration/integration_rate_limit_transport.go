//go:build integration

package integration_test

import (
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// rateLimitRetryTransport spaces out requests and retries GET/HEAD on 429 with Retry-After /
// exponential backoff (aligned with utils/probe_id_paths_with_real_ids.py). POST/PATCH/PUT/DELETE
// are not auto-retried to avoid duplicate side effects; callers still see 429 after sleeps.
type rateLimitRetryTransport struct {
	base        http.RoundTripper
	mu          sync.Mutex
	minInterval time.Duration
	maxRetries  int
	lastStart   time.Time
}

func newRateLimitRetryTransport(base http.RoundTripper, minInterval time.Duration) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &rateLimitRetryTransport{
		base:        base,
		minInterval: minInterval,
		maxRetries:  12,
	}
}

func (t *rateLimitRetryTransport) throttle() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.minInterval <= 0 {
		return
	}
	// Space requests; does not include response latency (same as a simple client-side QPS cap).
	if !t.lastStart.IsZero() {
		elapsed := time.Since(t.lastStart)
		if elapsed < t.minInterval {
			time.Sleep(t.minInterval - elapsed)
		}
	}
	t.lastStart = time.Now()
}

func (t *rateLimitRetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.throttle()
	return t.roundTripWith429Retry(req)
}

func (t *rateLimitRetryTransport) roundTripWith429Retry(req *http.Request) (*http.Response, error) {
	for attempt := 0; attempt < t.maxRetries; attempt++ {
		resp, err := t.base.RoundTrip(req)
		if err != nil || resp == nil {
			return resp, err
		}
		if resp.StatusCode != http.StatusTooManyRequests {
			return resp, nil
		}
		safe := req.Method == http.MethodGet || req.Method == http.MethodHead ||
			req.Method == http.MethodOptions || req.Method == http.MethodTrace
		if !safe {
			sleep429(resp.Header.Get("Retry-After"), attempt)
			return resp, nil
		}
		if attempt+1 >= t.maxRetries {
			return resp, nil
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
		sleep429(resp.Header.Get("Retry-After"), attempt)
		cloned := req.Clone(req.Context())
		if req.Body != nil && req.GetBody != nil {
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			cloned.Body = body
		}
		req = cloned
		t.throttle()
	}
	// Unreachable: every iteration returns or continues until maxRetries-1 exhausts safe retries.
	return nil, nil
}

func sleep429(retryAfter string, attempt int) {
	ra := parseRetryAfterSeconds(retryAfter)
	exp := exponentialBackoffSeconds(attempt)
	s := maxDurationSeconds(ra, exp)
	if s > 120*time.Second {
		s = 120 * time.Second
	}
	time.Sleep(s)
}

func parseRetryAfterSeconds(h string) time.Duration {
	h = strings.TrimSpace(h)
	if h == "" {
		return 0
	}
	if n, err := strconv.Atoi(h); err == nil && n >= 0 {
		return time.Duration(n) * time.Second
	}
	return 0
}

func exponentialBackoffSeconds(attempt int) time.Duration {
	if attempt > 8 {
		attempt = 8
	}
	sec := 3.0 * math.Pow(1.5, float64(attempt))
	return time.Duration(sec * float64(time.Second))
}

func maxDurationSeconds(a, b time.Duration) time.Duration {
	if a >= b {
		return a
	}
	return b
}
