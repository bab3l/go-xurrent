//go:build integration

package integration_test

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	openapiclient "github.com/xurrent/go-xurrent"
)

// XURRENT_STRUCTURE_LOG_FILE sets the path to an NDJSON log of redacted request metadata and JSON
// response skeletons (no secrets, no query values, no raw headers). The directory is created as needed.
const envStructureLogFile = "XURRENT_STRUCTURE_LOG_FILE"

// XURRENT_HTTP_MIN_INTERVAL is optional seconds between outbound HTTP requests (e.g. "0.35") plus
// 429 handling (see integration_rate_limit_transport.go). Live capture tooling sets this when unset.
const envHTTPMinInterval = "XURRENT_HTTP_MIN_INTERVAL"

func httpMinIntervalFromEnv() time.Duration {
	raw := strings.TrimSpace(os.Getenv(envHTTPMinInterval))
	if raw == "" {
		return 0
	}
	f, err := strconv.ParseFloat(raw, 64)
	if err != nil || f <= 0 {
		return 0
	}
	return time.Duration(f * float64(time.Second))
}

// newIntegrationConfig builds API client configuration. If XURRENT_STRUCTURE_LOG_FILE is set,
// all HTTP traffic uses a redacting transport that writes structure-only logs to that file.
func newIntegrationConfig(t *testing.T, withAuth bool) *openapiclient.Configuration {
	t.Helper()
	cfg := openapiclient.NewConfiguration()
	if withAuth {
		token := os.Getenv("XURRENT_TOKEN")
		account := os.Getenv("XURRENT_ACCOUNT")
		if token == "" || account == "" {
			t.Skip("set XURRENT_TOKEN and XURRENT_ACCOUNT")
		}
		cfg.AddDefaultHeader("Authorization", "Bearer "+token)
		cfg.AddDefaultHeader("X-4me-Account", account)
	}
	if logPath := os.Getenv(envStructureLogFile); logPath != "" {
		if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
			t.Fatalf("structure log dir: %v", err)
		}
		f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			t.Fatalf("structure log file: %v", err)
		}
		t.Cleanup(func() { _ = f.Close() })
		var base http.RoundTripper
		base = http.DefaultTransport
		if d := httpMinIntervalFromEnv(); d > 0 {
			base = newRateLimitRetryTransport(http.DefaultTransport, d)
		}
		cfg.HTTPClient = &http.Client{
			Transport: newRedactingTransport(base, f),
		}
	} else if d := httpMinIntervalFromEnv(); d > 0 {
		cfg.HTTPClient = &http.Client{
			Transport: newRateLimitRetryTransport(http.DefaultTransport, d),
		}
	}
	return cfg
}
