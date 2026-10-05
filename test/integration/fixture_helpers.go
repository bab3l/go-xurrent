//go:build integration

package integration_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	openapiclient "github.com/xurrent/go-xurrent"
)

// parentFixtureEnvPath returns the path to parent_fixture.env (cwd may be repo root or test/integration).
func parentFixtureEnvPath(t *testing.T) string {
	t.Helper()
	return parentFixtureEnvPathOnDisk()
}

// parentFixtureEnvPathOnDisk resolves parent_fixture.env relative to the current working directory.
func parentFixtureEnvPathOnDisk() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	if filepath.Base(wd) == "integration" {
		return filepath.Join(wd, "_live_shapes", "parent_fixture.env")
	}
	return filepath.Join(wd, "test", "integration", "_live_shapes", "parent_fixture.env")
}

// readParentFixtureKey returns a KEY=value from parent_fixture.env on disk (same run as TestFixture_*).
func readParentFixtureKey(key string) string {
	path := parentFixtureEnvPathOnDisk()
	if path == "" {
		return ""
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		k, val, ok := strings.Cut(line, "=")
		if !ok || strings.TrimSpace(k) != key {
			continue
		}
		return strings.TrimSpace(val)
	}
	return ""
}

// int32FromEnvOrFixture parses key from the process environment, then parent_fixture.env.
func int32FromEnvOrFixture(key string) (int32, bool) {
	for _, s := range []string{os.Getenv(key), readParentFixtureKey(key)} {
		if s == "" {
			continue
		}
		n, err := strconv.ParseInt(s, 10, 32)
		if err == nil && n > 0 {
			return int32(n), true
		}
	}
	return 0, false
}

// requestIDFromFirstRow parses id from the first row of a collection of JSON objects (requests, problems, workflows, etc.).
func requestIDFromFirstRow(rows []map[string]interface{}) (int32, bool) {
	if len(rows) == 0 {
		return 0, false
	}
	raw, ok := rows[0]["id"]
	if !ok {
		return 0, false
	}
	switch v := raw.(type) {
	case float64:
		return int32(v), true
	case int:
		return int32(v), true
	case int32:
		return v, true
	case int64:
		return int32(v), true
	default:
		return 0, false
	}
}

// int32IDFromHTTPJSONArray tries to read the first row's id from a list endpoint that returns only *http.Response.
func int32IDFromHTTPJSONArray(resp *http.Response, err error) (int32, bool) {
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
		return requestIDFromFirstRow(rows)
	}
	var wrap map[string]interface{}
	if json.Unmarshal(body, &wrap) != nil {
		return 0, false
	}
	if raw, ok := wrap["data"]; ok {
		switch v := raw.(type) {
		case []interface{}:
			if len(v) == 0 {
				return 0, false
			}
			if m, ok := v[0].(map[string]interface{}); ok {
				return requestIDFromFirstRow([]map[string]interface{}{m})
			}
		}
	}
	return 0, false
}

// firstProblemIDFromAPI returns a problem id from predefined problem lists, then GET /v1/problems (response-body only).
func firstProblemIDFromAPI(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	type tryFn func() ([]map[string]interface{}, *http.Response, error)
	tries := []tryFn{
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.ProblemsAPI.GetProblemsActive(ctx).PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.ProblemsAPI.GetProblemsAssignedToMe(ctx).PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.ProblemsAPI.GetProblemsAssignedToMyTeam(ctx).PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.ProblemsAPI.GetProblemsManagedByMe(ctx).PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.ProblemsAPI.GetProblemsKnownErrors(ctx).PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.ProblemsAPI.GetProblemsSolved(ctx).PerPage(1).Fields("id").Execute()
		},
	}
	for _, fn := range tries {
		rows, resp, err := fn()
		if err != nil || resp == nil || resp.StatusCode != 200 {
			continue
		}
		if id, ok := requestIDFromFirstRow(rows); ok {
			return id, true
		}
	}
	return int32IDFromHTTPJSONArray(client.ProblemsAPI.GetProblems(ctx).PerPage(1).Fields("id").Execute())
}

// problemIDFromEnvOrAPI uses XURRENT_TEST_PROBLEM_ID from env, then parent_fixture.env, then problem list discovery (see firstProblemIDFromAPI).
func problemIDFromEnvOrAPI(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	if id, ok := int32FromEnvOrFixture("XURRENT_TEST_PROBLEM_ID"); ok {
		return id, true
	}
	return firstProblemIDFromAPI(t, ctx, client)
}

// firstOpenRequestID returns a request id the token can read, preferring "open" lists.
// The API exposes open requests via GET /v1/requests/open (and optional state=open on other predefined filters);
// there is no generic GET /v1/requests collection in the generated client. If /open is empty, we try other
// predefined filters (waiting for me, assigned to me, etc.) with state=open, then without state.
func firstOpenRequestID(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	type tryFn func() ([]map[string]interface{}, *http.Response, error)
	tries := []tryFn{
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.RequestsAPI.GetRequestsOpen(ctx).PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.RequestsAPI.GetRequestsOpen(ctx).State("open").PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.RequestsAPI.GetRequestsWaitingForMe(ctx).State("open").PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.RequestsAPI.GetRequestsAssignedToMe(ctx).State("open").PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.RequestsAPI.GetRequestsAssignedToMyTeam(ctx).State("open").PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.RequestsAPI.GetRequestsRequestedByOrForMe(ctx).State("open").PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.RequestsAPI.GetRequestsWaitingForMe(ctx).PerPage(1).Fields("id").Execute()
		},
		func() ([]map[string]interface{}, *http.Response, error) {
			return client.RequestsAPI.GetRequestsAssignedToMe(ctx).PerPage(1).Fields("id").Execute()
		},
	}
	for _, fn := range tries {
		rows, resp, err := fn()
		if err != nil || resp == nil || resp.StatusCode != 200 {
			continue
		}
		if id, ok := requestIDFromFirstRow(rows); ok {
			return id, true
		}
	}
	return 0, false
}

// openRequestIDFromEnvOrAPI uses XURRENT_TEST_OPEN_REQUEST_ID from env, then parent_fixture.env, then open-request discovery (see firstOpenRequestID).
func openRequestIDFromEnvOrAPI(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	if id, ok := int32FromEnvOrFixture("XURRENT_TEST_OPEN_REQUEST_ID"); ok {
		return id, true
	}
	return firstOpenRequestID(t, ctx, client)
}

// int32FromOptionalFloat32ID converts list row ids from generated models (*float32) to int32.
func int32FromOptionalFloat32ID(id *float32) (int32, bool) {
	if id == nil || *id == 0 {
		return 0, false
	}
	return int32(*id), true
}

// firstTeamIDFromAPI returns a team id from GET /v1/teams.
func firstTeamIDFromAPI(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.TeamsAPI.GetTeams(ctx).PerPage(1).Fields("id").Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 || len(rows) == 0 {
		return 0, false
	}
	return int32FromOptionalFloat32ID(rows[0].Id)
}

// firstSiteIDFromAPI returns a site id from GET /v1/sites.
func firstSiteIDFromAPI(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.SitesAPI.GetSites(ctx).PerPage(1).Fields("id").Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 || len(rows) == 0 {
		return 0, false
	}
	return int32FromOptionalFloat32ID(rows[0].Id)
}

// firstSLAIDFromAPI returns an SLA id from active/all/inactive lists.
func firstSLAIDFromAPI(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.ServiceLevelAgreementsAPI.GetSlasActive(ctx).PerPage(1).Fields("id").Execute()
	if err == nil && resp != nil && resp.StatusCode == 200 && len(rows) > 0 {
		if id, ok := requestIDFromFirstRow(rows); ok {
			return id, true
		}
	}
	rowsM, resp2, err2 := client.ServiceLevelAgreementsAPI.GetSlas(ctx).PerPage(1).Fields("id").Execute()
	if err2 == nil && resp2 != nil && resp2.StatusCode == 200 {
		if id, ok := requestIDFromFirstRow(rowsM); ok {
			return id, true
		}
	}
	rowsI, resp3, err3 := client.ServiceLevelAgreementsAPI.GetSlasInactive(ctx).PerPage(1).Fields("id").Execute()
	if err3 == nil && resp3 != nil && resp3.StatusCode == 200 {
		if id, ok := requestIDFromFirstRow(rowsI); ok {
			return id, true
		}
	}
	return 0, false
}

// firstTaskIDFromAPI returns a task id from list endpoints (finished and approval lists decode to typed rows; others use raw JSON body).
func firstTaskIDFromAPI(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.TasksAPI.GetTasksFinished(ctx).PerPage(1).Fields("id").Execute()
	if err == nil && resp != nil && resp.StatusCode == 200 && len(rows) > 0 {
		if id, ok := requestIDFromFirstRow(rows); ok {
			return id, true
		}
	}
	appr, resp, err := client.TasksAPI.GetTasksApprovalByMe(ctx).PerPage(1).Fields("id").Execute()
	if err == nil && resp != nil && resp.StatusCode == 200 && len(appr) > 0 {
		if id, ok := int32FromOptionalFloat32ID(appr[0].Id); ok {
			return id, true
		}
	}
	httpTries := []func() (*http.Response, error){
		func() (*http.Response, error) {
			return client.TasksAPI.GetTasksOpen(ctx).PerPage(1).Fields("id").Execute()
		},
		func() (*http.Response, error) {
			return client.TasksAPI.GetTasksAssignedToMe(ctx).PerPage(1).Fields("id").Execute()
		},
		func() (*http.Response, error) {
			return client.TasksAPI.GetTasksManagedByMe(ctx).PerPage(1).Fields("id").Execute()
		},
		func() (*http.Response, error) {
			return client.TasksAPI.GetTasksAssignedToMyTeam(ctx).PerPage(1).Fields("id").Execute()
		},
		func() (*http.Response, error) {
			return client.TasksAPI.GetTasksAssignedByMe(ctx).PerPage(1).Fields("id").Execute()
		},
	}
	for _, fn := range httpTries {
		if id, ok := int32IDFromHTTPJSONArray(fn()); ok {
			return id, true
		}
	}
	return 0, false
}
