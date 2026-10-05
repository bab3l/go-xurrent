//go:build integration

// Opt-in archive / trash / restore lifecycle tests plus POST→GET fixtures for disposable records.
//
//   set XURRENT_ALLOW_MUTATIONS=1
//   set XURRENT_ALLOW_LIFECYCLE_MUTATIONS=1
//   go test -tags=integration ./test/integration/ -run TestLifecycle_ -v -count=1
//
// Workflows (and workflow tasks) need XURRENT_TEST_WORKFLOW_TEMPLATE_ID (and working PostServices).
// Requests need XURRENT_TEST_SERVICE_INSTANCE_ID and XURRENT_TEST_REQUEST_TEMPLATE_ID (numeric ids).

package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/xurrent/go-xurrent"
)

func lifecycleClient(t *testing.T) *openapiclient.APIClient {
	t.Helper()
	if os.Getenv("XURRENT_ALLOW_MUTATIONS") != "1" {
		t.Skip("set XURRENT_ALLOW_MUTATIONS=1 to run mutation integration tests")
	}
	if os.Getenv("XURRENT_ALLOW_LIFECYCLE_MUTATIONS") != "1" {
		t.Skip("set XURRENT_ALLOW_LIFECYCLE_MUTATIONS=1 for archive/trash/restore lifecycle tests")
	}
	cfg := newIntegrationConfig(t, true)
	return openapiclient.NewAPIClient(cfg)
}

func firstTeamID(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.TeamsAPI.GetTeams(ctx).PerPage(1).Fields("id").Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 || len(rows) == 0 {
		return 0, false
	}
	if !rows[0].HasId() {
		return 0, false
	}
	return int32(rows[0].GetId()), true
}

func decodeMapFromHTTPResponse(t *testing.T, resp *http.Response) map[string]interface{} {
	t.Helper()
	if resp == nil || resp.Body == nil {
		return nil
	}
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	_ = resp.Body.Close()
	b = []byte(strings.TrimSpace(string(b)))
	if len(b) == 0 {
		return nil
	}
	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(b, &m))
	return m
}

func createWorkflowWithService(t *testing.T, ctx context.Context, client *openapiclient.APIClient, acc string, templateID int32, wfType string) (workflowID, svcID int32, sfx string) {
	t.Helper()
	sfx = randomSuffix()
	orgID, ok := firstOrganizationID(t, ctx, client)
	require.True(t, ok, "organization for service")

	svcName := fmt.Sprintf("sdk-lc-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":     svcName,
		"provider": map[string]interface{}{"id": orgID},
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode == 401 || resp.StatusCode == 403 {
		t.Skipf("PostServices not available: %v HTTP %v", err, statusOrZero(resp))
	}
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, svcCreated)
	svcID, ok = mapIDInt32Any(svcCreated)
	require.True(t, ok, "service id")

	managerID := mePersonID(t, ctx, client)
	wfBody := map[string]interface{}{
		"subject":           fmt.Sprintf("sdk-lc-wf-%s", sfx),
		"service":           map[string]interface{}{"id": svcID},
		"category":          "rfc",
		"manager":           map[string]interface{}{"id": managerID},
		"workflow_type":     wfType,
		"workflow_template": map[string]interface{}{"id": templateID},
	}
	wfCreated, resp, err := client.WorkflowsAPI.PostWorkflows(ctx).Body(wfBody).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, wfCreated)
	workflowID, ok = mapIDInt32Any(wfCreated)
	require.True(t, ok, "workflow id")
	return workflowID, svcID, sfx
}

func statusOrZero(resp *http.Response) int {
	if resp == nil {
		return 0
	}
	return resp.StatusCode
}

// TestLifecycle_WorkflowArchiveRestoreTrashAndTask exercises GET /v1/workflows/{id}, PostWorkflowsIdTasks,
// PostWorkflowsIdArchive, PostWorkflowsIdRestore, PostWorkflowsIdTrash.
func TestLifecycle_WorkflowArchiveRestoreTrashAndTask(t *testing.T) {
	tplStr := os.Getenv("XURRENT_TEST_WORKFLOW_TEMPLATE_ID")
	if tplStr == "" {
		t.Skip("set XURRENT_TEST_WORKFLOW_TEMPLATE_ID for workflow lifecycle test")
	}
	tpl64, err := strconv.ParseInt(tplStr, 10, 32)
	require.NoError(t, err)
	templateID := int32(tpl64)
	wfType := os.Getenv("XURRENT_TEST_WORKFLOW_TYPE")
	if wfType == "" {
		wfType = "application_change"
	}

	client := lifecycleClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")

	var workflowID, svcID int32
	defer func() {
		bctx := context.Background()
		if workflowID != 0 {
			_, _, _ = client.WorkflowsAPI.PostWorkflowsIdTrash(bctx, workflowID).Execute()
		}
		if svcID != 0 {
			_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
				"disabled": true,
				"remarks":  "disabled after sdk lifecycle workflow test",
			}).Execute()
		}
	}()

	var sfx string
	workflowID, svcID, sfx = createWorkflowWithService(t, ctx, client, acc, templateID, wfType)

	_, resp, err := client.WorkflowsAPI.GetWorkflowsId(ctx, workflowID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	taskBody := map[string]interface{}{
		"subject":  fmt.Sprintf("sdk-lifecycle-task-%s", sfx),
		"category": "implementation",
		"note":     fmt.Sprintf("sdk lifecycle task note %s", sfx),
	}
	_, resp, err = client.WorkflowsAPI.PostWorkflowsIdTasks(ctx, workflowID).Body(taskBody).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		t.Fatalf("PostWorkflowsIdTasks: %v HTTP %v (adjust task body fields for your tenant)", err, statusOrZero(resp))
	}

	_, resp, err = client.WorkflowsAPI.PostWorkflowsIdArchive(ctx, workflowID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.WorkflowsAPI.PostWorkflowsIdRestore(ctx, workflowID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.WorkflowsAPI.PostWorkflowsIdTrash(ctx, workflowID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	workflowID = 0
}

// TestLifecycle_ProblemArchiveRestoreTrash creates a problem via POST, GET by id, then archive/restore/trash/restore
// and marks it solved for cleanup.
func TestLifecycle_ProblemArchiveRestoreTrash(t *testing.T) {
	client := lifecycleClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")

	orgID, ok := firstOrganizationID(t, ctx, client)
	require.True(t, ok, "organization")
	teamID, ok := firstTeamID(t, ctx, client)
	if !ok {
		t.Skip("need at least one team")
	}

	sfx := randomSuffix()
	svcName := fmt.Sprintf("sdk-lc-prob-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":     svcName,
		"provider": map[string]interface{}{"id": orgID},
	}).Execute()
	if err != nil || resp == nil || (resp.StatusCode != 201 && resp.StatusCode != 200) {
		if resp != nil && (resp.StatusCode == 401 || resp.StatusCode == 403) {
			t.Skipf("PostServices not permitted")
		}
		if resp != nil && resp.StatusCode == 422 {
			t.Skipf("PostServices validation failed (HTTP 422); cannot create problem fixture service: %v", err)
		}
		require.NoError(t, err)
	}
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, svcCreated)
	svcID, ok := mapIDInt32Any(svcCreated)
	require.True(t, ok, "service id")

	me := mePersonID(t, ctx, client)
	probBody := map[string]interface{}{
		"service_id":      svcID,
		"requested_by_id": me,
		"manager":         me,
		"impact":          "medium",
		"subject":         fmt.Sprintf("sdk-lc-problem-%s", sfx),
		"team":            fmt.Sprintf("%d", teamID),
		"category":        "proactive",
		"note":            fmt.Sprintf("sdk lifecycle problem %s", sfx),
	}
	pResp, err := client.ProblemsAPI.PostProblems(ctx).Body(probBody).Execute()
	require.NoError(t, err)
	require.NotNil(t, pResp)
	require.Equal(t, 200, pResp.StatusCode, "PostProblems expected 200 per spec")
	created := decodeMapFromHTTPResponse(t, pResp)
	require.NotNil(t, created)
	pid, ok := mapIDInt32(created)
	require.True(t, ok, "problem id from POST body")

	defer func() {
		bctx := context.Background()
		_, _, _ = client.ProblemsAPI.PostProblemsIdRestore(bctx, pid).Execute()
		pr, err := client.ProblemsAPI.PatchProblemsId(bctx, pid).Body(map[string]interface{}{
			"status": "solved",
			"note":   fmt.Sprintf("closed after sdk lifecycle test %s", sfx),
		}).Execute()
		if err == nil && pr != nil {
			_, _ = io.Copy(io.Discard, pr.Body)
			_ = pr.Body.Close()
		}
		_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
			"disabled": true,
			"remarks":  "disabled after sdk lifecycle problem test",
		}).Execute()
	}()

	_, resp, err = client.ProblemsAPI.GetProblemsId(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProblemsAPI.PostProblemsIdArchive(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProblemsAPI.PostProblemsIdRestore(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProblemsAPI.PostProblemsIdTrash(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProblemsAPI.PostProblemsIdRestore(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}

// TestLifecycle_RequestArchiveRestoreTrash creates a request, GET by id, archive/restore/trash/restore.
func TestLifecycle_RequestArchiveRestoreTrash(t *testing.T) {
	siStr := os.Getenv("XURRENT_TEST_SERVICE_INSTANCE_ID")
	tplStr := os.Getenv("XURRENT_TEST_REQUEST_TEMPLATE_ID")
	if siStr == "" || tplStr == "" {
		t.Skip("set XURRENT_TEST_SERVICE_INSTANCE_ID and XURRENT_TEST_REQUEST_TEMPLATE_ID (numeric)")
	}
	si64, err := strconv.ParseInt(siStr, 10, 32)
	require.NoError(t, err)
	tpl64, err := strconv.ParseInt(tplStr, 10, 32)
	require.NoError(t, err)
	serviceInstanceID := int32(si64)
	templateID := int32(tpl64)

	client := lifecycleClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")
	me := mePersonID(t, ctx, client)
	sfx := randomSuffix()

	body := map[string]interface{}{
		"requested_by":        map[string]interface{}{"id": me},
		"requested_for":       map[string]interface{}{"id": me},
		"service_instance_id": serviceInstanceID,
		"subject":             fmt.Sprintf("sdk-lc-request-%s", sfx),
		"template":            templateID,
		"category":            "rfi",
		"note":                fmt.Sprintf("sdk lifecycle request %s", sfx),
	}
	created, resp, err := client.RequestsAPI.PostRequests(ctx).Body(body).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, created)
	rid, ok := mapIDInt32Any(created)
	require.True(t, ok, "request id")

	defer func() {
		bctx := context.Background()
		_, _, _ = client.RequestsAPI.PostRequestsIdRestore(bctx, rid).Execute()
	}()

	_, resp, err = client.RequestsAPI.GetRequestsId(ctx, rid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.RequestsAPI.PostRequestsIdArchive(ctx, rid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.RequestsAPI.PostRequestsIdRestore(ctx, rid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.RequestsAPI.PostRequestsIdTrash(ctx, rid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.RequestsAPI.PostRequestsIdRestore(ctx, rid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}

// TestLifecycle_PersonArchiveRestoreTrash creates a person, GET by id, archive/restore/trash/restore, then disables.
func TestLifecycle_PersonArchiveRestoreTrash(t *testing.T) {
	client := lifecycleClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")
	sfx := randomSuffix()
	email := fmt.Sprintf("sdk-lc-%s@invalid.example", sfx)

	created, resp, err := client.PeopleAPI.PostPeople(ctx).Body(map[string]interface{}{
		"source":        fmt.Sprintf("sdk-lifecycle-%s", sfx),
		"primary_email": email,
		"name":          fmt.Sprintf("SDK Lifecycle %s", sfx),
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode == 403 || resp.StatusCode == 401 {
		t.Skipf("PostPeople not permitted or failed: %v HTTP %v", err, statusOrZero(resp))
	}
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	if id, ok := resourceAccountID(jsonMap(created)); ok && id != "" && !resourceAccountAcceptableForEnv(id, acc) {
		t.Skipf("person created in account %q not acceptable for XURRENT_ACCOUNT %q (set XURRENT_TRUSTED_ACCOUNT_IDS for trust peers); skip lifecycle", id, acc)
	}
	pid, ok := mapIDInt32Any(created)
	require.True(t, ok, "person id")

	defer func() {
		bctx := context.Background()
		_, _, _ = client.PeopleAPI.PostPeopleIdRestore(bctx, pid).Execute()
		_, _, _ = client.PeopleAPI.PatchPeopleId(bctx, pid).Body(map[string]interface{}{
			"disabled": true,
			"remarks":  fmt.Sprintf("disabled after sdk lifecycle person test %s", sfx),
		}).Execute()
	}()

	_, resp, err = client.PeopleAPI.GetPeopleId(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.PeopleAPI.PostPeopleIdArchive(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.PeopleAPI.PostPeopleIdRestore(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.PeopleAPI.PostPeopleIdTrash(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.PeopleAPI.PostPeopleIdRestore(ctx, pid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}
