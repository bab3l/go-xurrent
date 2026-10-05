//go:build integration

// Platform endpoints not covered by GET sweep: export/import jobs, events, optional import POST.
// Opt-in with the same gates as write-shape (XURRENT_ALLOW_MUTATIONS + XURRENT_WRITE_SHAPE_COVERAGE).
//
// Export → poll GET /v1/import/{token} captures the linked read after write.
// PostImport is tenant- and file-format sensitive; enable with XURRENT_WRITE_CAPTURE_POST_IMPORT=1
// and XURRENT_WRITE_CAPTURE_IMPORT_CSV_PATH pointing at a minimal CSV the account accepts.
//
// PostEvents needs service instance + request template ids (same as lifecycle tests).

package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func exportJobTokenFromBody(raw []byte) string {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 {
		return ""
	}
	if raw[0] == '"' {
		var s string
		if json.Unmarshal(raw, &s) == nil {
			return strings.TrimSpace(s)
		}
	}
	var m map[string]interface{}
	if json.Unmarshal(raw, &m) != nil {
		return ""
	}
	for _, k := range []string{"token", "id", "job_id", "import_token"} {
		if v, ok := m[k]; ok && v != nil {
			s := strings.TrimSpace(fmt.Sprint(v))
			if s != "" {
				return s
			}
		}
	}
	return ""
}

// TestWriteCapture_PostExportAndPollImportToken starts a small export and polls the job once.
// Uses random-free identifiers only in export parameters (from date); capture NDJSON is still skeletonized.
func TestWriteCapture_PostExportAndPollImportToken(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()

	resp, err := client.ExportAPI.PostExport(ctx).
		Type_("sites").
		From(20200101).
		ExportFormat("csv").
		LineSeparator("lf").
		Execute()
	if err != nil || resp == nil {
		logWriteResult(t, "PostExport", resp, err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		t.Logf("PostExport read body: %v", err)
		return
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		t.Logf("PostExport: HTTP %d body=%s", resp.StatusCode, truncateForLog(body))
		return
	}
	token := exportJobTokenFromBody(body)
	if token == "" {
		t.Logf("PostExport: no token in body: %s", truncateForLog(body))
		return
	}
	t.Logf("PostExport: ok; polling import job token")

	poll, err := client.ExportAPI.GetImportToken(ctx, token).Execute()
	logWriteResult(t, "GetImportToken", poll, err)
	if err == nil && poll != nil && poll.StatusCode == 200 {
		_, _ = io.Copy(io.Discard, poll.Body)
		_ = poll.Body.Close()
	}
}

// TestWriteCapture_GetImportList lists export/import jobs (GET /v1/import).
func TestWriteCapture_GetImportList(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()
	resp, err := client.ExportAPI.GetImport(ctx).PerPage(5).Fields("id,token,state").Execute()
	logWriteResult(t, "GetImport", resp, err)
	if err == nil && resp != nil && resp.StatusCode == 200 {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}

// TestWriteCapture_PostImportOptional posts a JSON body only when explicitly enabled (see file comment).
func TestWriteCapture_PostImportOptional(t *testing.T) {
	if os.Getenv("XURRENT_WRITE_CAPTURE_POST_IMPORT") != "1" {
		t.Skip("set XURRENT_WRITE_CAPTURE_POST_IMPORT=1 and XURRENT_WRITE_CAPTURE_IMPORT_CSV_PATH to run PostImport")
	}
	csvPath := strings.TrimSpace(os.Getenv("XURRENT_WRITE_CAPTURE_IMPORT_CSV_PATH"))
	if csvPath == "" {
		t.Skip("XURRENT_WRITE_CAPTURE_IMPORT_CSV_PATH must point at a CSV file readable by this process")
	}
	client := writeShapeClient(t)
	ctx := context.Background()
	dataType := os.Getenv("XURRENT_WRITE_CAPTURE_IMPORT_DATA_TYPE")
	if dataType == "" {
		dataType = "people"
	}
	body := map[string]interface{}{
		"files": map[string]interface{}{
			"file": csvPath,
		},
		"data": map[string]interface{}{
			"type": dataType,
		},
	}
	resp, err := client.ImportAPI.PostImport(ctx).Body(body).Execute()
	logWriteResult(t, "PostImport", resp, err)
	if err == nil && resp != nil && resp.Body != nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}

// TestWriteCapture_PostEvents uses fixture ids when set (same env names as lifecycle request tests).
func TestWriteCapture_PostEvents(t *testing.T) {
	siRaw := strings.TrimSpace(os.Getenv("XURRENT_TEST_SERVICE_INSTANCE_ID"))
	tplRaw := strings.TrimSpace(os.Getenv("XURRENT_TEST_REQUEST_TEMPLATE_ID"))
	if siRaw == "" || tplRaw == "" {
		t.Skip("set XURRENT_TEST_SERVICE_INSTANCE_ID and XURRENT_TEST_REQUEST_TEMPLATE_ID for PostEvents capture")
	}
	siID, err1 := strconv.ParseInt(siRaw, 10, 32)
	tplID, err2 := strconv.ParseInt(tplRaw, 10, 32)
	if err1 != nil || err2 != nil {
		t.Fatalf("invalid XURRENT_TEST_SERVICE_INSTANCE_ID or XURRENT_TEST_REQUEST_TEMPLATE_ID")
	}

	client := writeShapeClient(t)
	ctx := context.Background()
	sfx := randomSuffix()
	payload := map[string]interface{}{
		"source":              "sdk-write-capture",
		"sourceID":            "evt-" + sfx,
		"service_instance_id": int32(siID),
		"template_id":         int32(tplID),
		"subject":             fmt.Sprintf("capture event %s", sfx),
		"impact":              "top",
		"category":            "other",
		"note":                fmt.Sprintf("NDJSON capture note %s", sfx),
	}
	_, resp, err := client.EventsAPI.PostEvents(ctx).Body(payload).Execute()
	logWriteResult(t, "PostEvents", resp, err)
}

func truncateForLog(b []byte) string {
	const max = 512
	if len(b) <= max {
		return string(b)
	}
	return string(b[:max]) + "…"
}
