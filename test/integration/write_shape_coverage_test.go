//go:build integration

// Extra POST/PATCH/PUT calls to widen NDJSON capture for merge_skeleton_into_openapi.py.
// Best-effort: failures (401/403/404/422) are logged; tests still pass.
//
// Requires:
//   XURRENT_ALLOW_MUTATIONS=1
//   XURRENT_WRITE_SHAPE_COVERAGE=1
//
//   go test -tags=integration ./test/integration/ -run TestWriteShape_ -v -count=1
//
// Run with XURRENT_STRUCTURE_LOG_FILE set (e.g. via live_capture_pipeline.py) to append NDJSON.

package integration_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/xurrent/go-xurrent"
)

func writeShapeClient(t *testing.T) *openapiclient.APIClient {
	t.Helper()
	if os.Getenv("XURRENT_WRITE_SHAPE_COVERAGE") != "1" {
		t.Skip("set XURRENT_WRITE_SHAPE_COVERAGE=1 for extra POST/PATCH/PUT shape capture")
	}
	return mutationClient(t)
}

func firstMapRowID(rows []map[string]interface{}) (int32, bool) {
	if len(rows) == 0 {
		return 0, false
	}
	return mapIDInt32(rows[0])
}

// firstIDFromDecodedList extracts the first row id from list payloads decoded as
// []map, a single map, map with a "data" array, or []interface{}.
func firstIDFromDecodedList(payload interface{}) (int32, bool) {
	switch v := payload.(type) {
	case nil:
		return 0, false
	case []map[string]interface{}:
		return firstMapRowID(v)
	case map[string]interface{}:
		if id, ok := mapIDInt32(v); ok {
			return id, true
		}
		if raw, ok := v["data"]; ok {
			return firstIDFromDecodedList(raw)
		}
	case []interface{}:
		if len(v) == 0 {
			return 0, false
		}
		if m, ok := v[0].(map[string]interface{}); ok {
			return mapIDInt32(m)
		}
	}
	return 0, false
}

// firstIDFromJSONArrayResponse parses a JSON body from list endpoints whose generated
// client returns (*http.Response, error) instead of ([]map, *http.Response, error).
func firstIDFromJSONArrayResponse(resp *http.Response, err error) (int32, bool) {
	if err != nil || resp == nil || resp.StatusCode != 200 {
		return 0, false
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, false
	}
	var rows []map[string]interface{}
	if json.Unmarshal(body, &rows) == nil && len(rows) > 0 {
		return mapIDInt32(rows[0])
	}
	var wrap map[string]interface{}
	if json.Unmarshal(body, &wrap) != nil {
		return 0, false
	}
	return firstIDFromDecodedList(wrap)
}

func logWriteResult(t *testing.T, label string, resp *http.Response, err error) {
	t.Helper()
	if err != nil {
		t.Logf("%s: %v", label, err)
		return
	}
	if resp == nil {
		t.Logf("%s: nil response", label)
		return
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		t.Logf("%s: HTTP %d", label, resp.StatusCode)
	}
}

// TestWriteShape_SearchGET hits GET /v1/search (gap: get_search).
// Official docs use multi-word queries (e.g. q=windows); very short q may return 400 — see docs/internal/GET_SEARCH_AND_NDJSON_GAPS.md.
func TestWriteShape_SearchGET(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()
	// Match documented examples: https://developer.xurrent.com/v1/search/
	resp, err := client.SearchAPI.GetSearch(ctx).Q("windows").PerPage(25).Execute()
	if err != nil {
		var ge *openapiclient.GenericOpenAPIError
		if errors.As(err, &ge) && len(ge.Body()) > 0 {
			t.Logf("GetSearch: %v — body: %s", err, string(ge.Body()))
		} else {
			t.Logf("GetSearch: %v", err)
		}
		return
	}
	logWriteResult(t, "GetSearch", resp, nil)
}

// TestWriteShape_SearchGETForEachPage exercises GetSearchJSON + ForEachSearchPage + SearchNextPageToken
// against the live API (best-effort; logs errors, does not fail the run).
func TestWriteShape_SearchGETForEachPage(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()
	q := url.Values{}
	q.Set("q", "windows")
	q.Set("per_page", "25")
	var pages int
	err := client.ForEachSearchPage(ctx, q, 5, func(body []byte, resp *http.Response) error {
		pages++
		if resp != nil {
			t.Logf("search page %d: HTTP %d, %d bytes, next=%q", pages, resp.StatusCode, len(body), openapiclient.SearchNextPageToken(resp))
		}
		return nil
	})
	if err != nil {
		t.Logf("ForEachSearchPage: %v", err)
		return
	}
	t.Logf("ForEachSearchPage: completed %d page(s)", pages)
}

// TestWriteShape_PatchOrganizationsId PATCHes remarks on the first organization (gap: patch_organizations_id).
func TestWriteShape_PatchOrganizationsId(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()
	orgID, ok := firstOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("no organization")
	}
	sfx := randomSuffix()
	remarks := fmt.Sprintf("write-shape-org-%s", sfx)
	ns := openapiclient.NullableString{}
	ns.Set(&remarks)
	org := openapiclient.Organization{Remarks: ns}
	_, resp, err := client.OrganizationsAPI.PatchOrganizationsId(ctx, orgID).Organization(org).Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 {
		logWriteResult(t, "PatchOrganizationsId", resp, err)
		return
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	t.Logf("PatchOrganizationsId: 200")
}

// TestWriteShape_PatchMapResources PATCHes remarks (or note) on first row of each collection when present.
func TestWriteShape_PatchMapResources(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()
	sfx := randomSuffix()
	tag := fmt.Sprintf("write-shape-%s", sfx)

	patch := func(label string, fn func() (*http.Response, error)) {
		t.Helper()
		resp, err := fn()
		logWriteResult(t, label, resp, err)
	}

	// Services
	if rows, resp, err := client.ServicesAPI.GetServices(ctx).PerPage(1).Fields("id").Execute(); err == nil && resp != nil && resp.StatusCode == 200 {
		if id, ok := firstMapRowID(rows); ok {
			patch("PatchServicesId", func() (*http.Response, error) {
				_, r, e := client.ServicesAPI.PatchServicesId(ctx, id).Body(map[string]interface{}{
					"remarks": tag + " svc",
				}).Execute()
				return r, e
			})
		}
	}

	// SLAs (list decodes to map or wrapped list)
	if rows, resp, err := client.ServiceLevelAgreementsAPI.GetSlas(ctx).PerPage(1).Fields("id").Execute(); err == nil && resp != nil && resp.StatusCode == 200 {
		if id, ok := firstIDFromDecodedList(rows); ok {
			patch("PatchSlasId", func() (*http.Response, error) {
				_, r, e := client.ServiceLevelAgreementsAPI.PatchSlasId(ctx, id).Body(map[string]interface{}{
					"remarks": tag + " sla",
				}).Execute()
				return r, e
			})
		}
	}

	// Teams
	if rows, resp, err := client.TeamsAPI.GetTeams(ctx).PerPage(1).Fields("id").Execute(); err == nil && resp != nil && resp.StatusCode == 200 {
		if len(rows) == 0 {
			// skip
		} else if id, ok := int32IDFromSliceFirst(&rows[0]); ok {
			patch("PatchTeamsId", func() (*http.Response, error) {
				_, r, e := client.TeamsAPI.PatchTeamsId(ctx, id).Body(map[string]interface{}{
					"remarks": tag + " team",
				}).Execute()
				return r, e
			})
		}
	}

	// Sites
	if rows, resp, err := client.SitesAPI.GetSites(ctx).PerPage(1).Fields("id").Execute(); err == nil && resp != nil && resp.StatusCode == 200 {
		if len(rows) == 0 {
			// skip
		} else if id, ok := int32IDFromSliceFirst(&rows[0]); ok {
			patch("PatchSitesId", func() (*http.Response, error) {
				_, r, e := client.SitesAPI.PatchSitesId(ctx, id).Body(map[string]interface{}{
					"remarks": tag + " site",
				}).Execute()
				return r, e
			})
		}
	}

	// Service offerings (list API returns *http.Response only)
	if id, ok := firstIDFromJSONArrayResponse(client.ServiceOfferingsAPI.GetServiceOfferings(ctx).PerPage(1).Fields("id").Execute()); ok {
		patch("PatchServiceOfferingsId", func() (*http.Response, error) {
			_, r, e := client.ServiceOfferingsAPI.PatchServiceOfferingsId(ctx, id).Body(map[string]interface{}{
				"remarks": tag + " off",
			}).Execute()
			return r, e
		})
	}

	// Request templates (list decodes to map or wrapped list)
	if rows, resp, err := client.RequestTemplatesAPI.GetRequestTemplates(ctx).PerPage(1).Fields("id").Execute(); err == nil && resp != nil && resp.StatusCode == 200 {
		if id, ok := firstIDFromDecodedList(rows); ok {
			patch("PatchRequestTemplatesId", func() (*http.Response, error) {
				_, r, e := client.RequestTemplatesAPI.PatchRequestTemplatesId(ctx, id).Body(map[string]interface{}{
					"remarks": tag + " rt",
				}).Execute()
				return r, e
			})
		}
	}

	// Calendars
	if rows, resp, err := client.CalendarsAPI.GetCalendars(ctx).PerPage(1).Fields("id").Execute(); err == nil && resp != nil && resp.StatusCode == 200 {
		if len(rows) == 0 {
			// skip
		} else if id, ok := int32IDFromSliceFirst(&rows[0]); ok {
			patch("PatchCalendarsId", func() (*http.Response, error) {
				_, r, e := client.CalendarsAPI.PatchCalendarsId(ctx, id).Body(map[string]interface{}{
					"remarks": tag + " cal",
				}).Execute()
				return r, e
			})
		}
	}

	// Problems (list API returns *http.Response only)
	if id, ok := firstIDFromJSONArrayResponse(client.ProblemsAPI.GetProblems(ctx).PerPage(1).Fields("id").Execute()); ok {
		patch("PatchProblemsId", func() (*http.Response, error) {
			return client.ProblemsAPI.PatchProblemsId(ctx, id).Body(map[string]interface{}{
				"note": tag + " note",
			}).Execute()
		})
	}

	// Requests (open) — PATCH /v1/requests/{id} is generated on AttachmentsAPI (same path; body may include note_attachments).
	if rid, ok := openRequestIDFromEnvOrAPI(t, ctx, client); ok {
		patch("PatchRequestsId", func() (*http.Response, error) {
			_, r, e := client.AttachmentsAPI.PatchRequestsId(ctx, rid).Body(map[string]interface{}{
				"note": tag + " req",
			}).Execute()
			return r, e
		})
	}

	// Reservations (list API returns *http.Response only; PATCH returns *http.Response only)
	if id, ok := firstIDFromJSONArrayResponse(client.ReservationsAPI.GetReservations(ctx).PerPage(1).Fields("id").Execute()); ok {
		patch("PatchReservationsId", func() (*http.Response, error) {
			return client.ReservationsAPI.PatchReservationsId(ctx, id).Body(map[string]interface{}{
				"remarks": tag + " res",
			}).Execute()
		})
	}

	// Workflows (env or first in list)
	if widStr := os.Getenv("XURRENT_TEST_WORKFLOW_ID"); widStr != "" {
		var wid int32
		_, _ = fmt.Sscanf(widStr, "%d", &wid)
		if wid != 0 {
			patch("PatchWorkflowsId", func() (*http.Response, error) {
				_, r, e := client.WorkflowsAPI.PatchWorkflowsId(ctx, wid).Body(map[string]interface{}{
					"remarks": tag + " wf",
				}).Execute()
				return r, e
			})
		}
	}
}

// TestWriteShape_PutRequestsId sends PUT /v1/requests/{id} with a minimal body when an open request exists (gap: put_requests_id).
func TestWriteShape_PutRequestsId(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()
	rid, ok := openRequestIDFromEnvOrAPI(t, ctx, client)
	if !ok {
		t.Skip("no open request (set XURRENT_TEST_OPEN_REQUEST_ID or ensure GET /v1/requests/open returns rows)")
	}
	curRaw, getResp, err := client.RequestsAPI.GetRequestsId(ctx, rid).Execute()
	require.NoError(t, err)
	require.NotNil(t, getResp)
	require.Equal(t, 200, getResp.StatusCode)
	_, _ = io.Copy(io.Discard, getResp.Body)
	_ = getResp.Body.Close()
	require.NotNil(t, curRaw)
	cur := jsonMap(curRaw)
	// PUT with same shape + touched subject to satisfy API if required
	if _, has := cur["subject"]; has {
		cur["subject"] = fmt.Sprintf("%v (write-shape put)", cur["subject"])
	}
	resp, err := client.RequestsAPI.PutRequestsId(ctx, rid).Body(cur).Execute()
	logWriteResult(t, "PutRequestsId", resp, err)
}

// TestWriteShape_PostSubresourceNotes adds a public note on an open request and a task note when ids exist.
func TestWriteShape_PostSubresourceNotes(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()
	sfx := randomSuffix()

	if rid, ok := openRequestIDFromEnvOrAPI(t, ctx, client); ok {
		_, resp, err := client.RequestsAPI.PostRequestsIdNotes(ctx, rid).Body(map[string]interface{}{
			"text":     fmt.Sprintf("write-shape note %s", sfx),
			"internal": false,
		}).Execute()
		logWriteResult(t, "PostRequestsIdNotes", resp, err)
	}

	if tid, ok := firstIDFromJSONArrayResponse(client.TasksAPI.GetTasksOpen(ctx).PerPage(1).Fields("id").Execute()); ok {
		_, r2, e2 := client.TasksAPI.PostTasksIdNotes(ctx, tid).Body(map[string]interface{}{
			"text": fmt.Sprintf("write-shape task note %s", sfx),
		}).Execute()
		logWriteResult(t, "PostTasksIdNotes", r2, e2)
	}
}

// TestWriteShape_PostPeopleContact adds a throwaway contact line for /v1/me (gap: post_people_id_contacts pattern).
func TestWriteShape_PostPeopleContact(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()
	pid := mePersonID(t, ctx, client)
	_, resp, err := client.PeopleAPI.PostPeopleIdContacts(ctx, pid).Body(map[string]interface{}{
		"label":     "mobile",
		"telephone": fmt.Sprintf("0999%s", randomSuffix()[:8]),
	}).Execute()
	logWriteResult(t, "PostPeopleIdContacts", resp, err)
}

// TestWriteShape_PostCalendarMinimal creates a calendar and disables it (gaps: post path coverage for calendars).
func TestWriteShape_PostCalendarMinimal(t *testing.T) {
	client := writeShapeClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")
	sfx := randomSuffix()
	name := fmt.Sprintf("write-shape-cal-%s", sfx)
	created, resp, err := client.CalendarsAPI.PostCalendars(ctx).Body(map[string]interface{}{
		"name": name,
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		logWriteResult(t, "PostCalendars", resp, err)
		return
	}
	requireResourceOwnedByEnvAccountAny(t, acc, created)
	cid, ok := mapIDInt32Any(created)
	require.True(t, ok)
	defer func() {
		_, _, _ = client.CalendarsAPI.PatchCalendarsId(context.Background(), cid).Body(map[string]interface{}{
			"disabled": true,
			"remarks":  "disabled after write-shape calendar test",
		}).Execute()
	}()
	t.Logf("PostCalendars: 201 id=%d", cid)
}
