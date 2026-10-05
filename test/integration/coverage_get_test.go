//go:build integration

// Broad authenticated GET smoke to exercise collection and filter endpoints that are easy to miss in generated stubs.
//
//   set XURRENT_ENDPOINT_COVERAGE=1
//   go test -tags=integration ./test/integration/ -run TestCoverage_ -v -count=1
//
// Uses the same credentials as production smoke tests. Some calls may return 403 (role / feature); those are logged and do not fail the test.
// TestCoverage_HighValueGETs hits workflows, teams, sites, organizations→services, calendars, service offerings,
// request templates (when a service exists), and SLA list filters; list-first id for GET /v1/workflows/{id}.
// TestCoverage_EntityGETByID hits GET /v1/teams/{id}, /v1/sites/{id}, /v1/slas/{id}, /v1/tasks/{id} using ids from list endpoints.
// (There is no GET /v1/calendars/{id} in the spec — only PATCH on that path.)

package integration_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/xurrent/go-xurrent"
)

func coverageClient(t *testing.T) *openapiclient.APIClient {
	t.Helper()
	if os.Getenv("XURRENT_ENDPOINT_COVERAGE") != "1" {
		t.Skip("set XURRENT_ENDPOINT_COVERAGE=1 to run broad GET coverage tests")
	}
	if os.Getenv("XURRENT_TOKEN") == "" || os.Getenv("XURRENT_ACCOUNT") == "" {
		t.Skip("set XURRENT_TOKEN and XURRENT_ACCOUNT")
	}
	cfg := newIntegrationConfig(t, true)
	return openapiclient.NewAPIClient(cfg)
}

// assertReadOK expects 200, or logs and accepts 403/404 when the API refuses or has nothing to return for this tenant.
func assertReadOK(t *testing.T, resp *http.Response, err error, label string) {
	t.Helper()
	if err != nil {
		if resp != nil && (resp.StatusCode == 403 || resp.StatusCode == 404) {
			t.Logf("%s: status %d (skipped)", label, resp.StatusCode)
			return
		}
		require.NoError(t, err, label)
	}
	require.NotNil(t, resp)
	if resp.StatusCode == 403 {
		t.Logf("%s: 403 (skipped)", label)
		return
	}
	require.Equal(t, 200, resp.StatusCode, label)
}

// assertReadOK200LenientDecode accepts HTTP 200 even when the generated client fails JSON decode
// (e.g. collection body typed as map in OpenAPI).
func assertReadOK200LenientDecode(t *testing.T, resp *http.Response, err error, label string) {
	t.Helper()
	if resp != nil && resp.StatusCode == 200 {
		if err != nil {
			t.Logf("%s: 200 OK (decode mismatch ignored): %v", label, err)
		}
		return
	}
	assertReadOK(t, resp, err, label)
}

func firstUsageStatementID(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.AccountAPI.GetAccountUsageStatements(ctx).PerPage(1).Fields("id").Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 {
		return 0, false
	}
	if len(rows) == 0 {
		return 0, false
	}
	if !rows[0].HasId() {
		return 0, false
	}
	return int32(rows[0].GetId()), true
}

func TestCoverage_HighValueGETs(t *testing.T) {
	client := coverageClient(t)
	ctx := context.Background()

	wfRows, resp, err := client.WorkflowsAPI.GetWorkflows(ctx).PerPage(1).Fields("id,subject").Execute()
	assertReadOK200LenientDecode(t, resp, err, "GetWorkflows")
	wid, haveWid := requestIDFromFirstRow(wfRows)
	if !haveWid {
		wid, haveWid = int32FromEnvOrFixture("XURRENT_TEST_WORKFLOW_ID")
	}

	_, resp, err = client.CalendarsAPI.GetCalendars(ctx).PerPage(1).Fields("id,name").Execute()
	assertReadOK(t, resp, err, "GetCalendars")

	_, resp, err = client.TeamsAPI.GetTeams(ctx).PerPage(1).Fields("id,name").Execute()
	assertReadOK(t, resp, err, "GetTeams")

	_, resp, err = client.SitesAPI.GetSites(ctx).PerPage(1).Fields("id,name").Execute()
	assertReadOK(t, resp, err, "GetSites")

	orgs, resp, err := client.OrganizationsAPI.GetOrganizations(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetOrganizations")
	if len(orgs) > 0 && orgs[0].Id != nil {
		svcs, resp, err := client.ServicesAPI.GetServices(ctx).Provider(int32(*orgs[0].Id)).PerPage(1).Fields("id,name").Execute()
		assertReadOK(t, resp, err, "GetServices")
		var svcID int32
		if len(svcs) > 0 {
			switch v := svcs[0]["id"].(type) {
			case float64:
				svcID = int32(v)
			case int:
				svcID = int32(v)
			case int32:
				svcID = v
			}
		}
		if svcID != 0 {
			respSO, errSO := client.ServiceOfferingsAPI.GetServiceOfferings(ctx).PerPage(1).Fields("id,name").Execute()
			assertReadOK(t, respSO, errSO, "GetServiceOfferings")
			if respSO != nil && respSO.StatusCode == 200 {
				_, _ = io.Copy(io.Discard, respSO.Body)
				_ = respSO.Body.Close()
			}
			_, resp, err = client.RequestTemplatesAPI.GetRequestTemplates(ctx).Service(strconv.FormatInt(int64(svcID), 10)).PerPage(1).Fields("id,subject").Execute()
			assertReadOK200LenientDecode(t, resp, err, "GetRequestTemplates")
		}
	}

	_, resp, err = client.ServiceLevelAgreementsAPI.GetSlasActive(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK200LenientDecode(t, resp, err, "GetSlasActive")

	_, resp, err = client.ServiceLevelAgreementsAPI.GetSlasInactive(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK200LenientDecode(t, resp, err, "GetSlasInactive")

	_, resp, err = client.ServiceLevelAgreementsAPI.GetSlas(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK200LenientDecode(t, resp, err, "GetSlas")

	if haveWid && wid != 0 {
		_, resp, err := client.WorkflowsAPI.GetWorkflowsId(ctx, wid).Execute()
		assertReadOK(t, resp, err, "GetWorkflowsId")
	}
}

func TestCoverage_EntityGETByID(t *testing.T) {
	client := coverageClient(t)
	ctx := context.Background()

	if tid, ok := firstTeamIDFromAPI(t, ctx, client); ok {
		_, resp, err := client.TeamsAPI.GetTeamsId(ctx, tid).Execute()
		assertReadOK200LenientDecode(t, resp, err, "GetTeamsId")
	} else {
		t.Log("no team id from GET /v1/teams; skipping GetTeamsId")
	}

	if sid, ok := firstSiteIDFromAPI(t, ctx, client); ok {
		_, resp, err := client.SitesAPI.GetSitesId(ctx, sid).Execute()
		assertReadOK200LenientDecode(t, resp, err, "GetSitesId")
	} else {
		t.Log("no site id from GET /v1/sites; skipping GetSitesId")
	}

	if slaID, ok := firstSLAIDFromAPI(t, ctx, client); ok {
		_, resp, err := client.ServiceLevelAgreementsAPI.GetSlasId(ctx, slaID).Execute()
		assertReadOK200LenientDecode(t, resp, err, "GetSlasId")
	} else {
		t.Log("no SLA id from SLA lists; skipping GetSlasId")
	}

	if taskID, ok := firstTaskIDFromAPI(t, ctx, client); ok {
		_, resp, err := client.TasksAPI.GetTasksId(ctx, taskID).Execute()
		assertReadOK200LenientDecode(t, resp, err, "GetTasksId")
	} else {
		t.Log("no task id from task lists; skipping GetTasksId")
	}
}

func TestCoverage_ProblemsFiltersAndPeopleDirectory(t *testing.T) {
	client := coverageClient(t)
	ctx := context.Background()

	_, resp, err := client.ProblemsAPI.GetProblemsActive(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetProblemsActive")

	_, resp, err = client.ProblemsAPI.GetProblemsAssignedToMe(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetProblemsAssignedToMe")

	_, resp, err = client.ProblemsAPI.GetProblemsAssignedToMyTeam(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetProblemsAssignedToMyTeam")

	_, resp, err = client.ProblemsAPI.GetProblemsKnownErrors(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetProblemsKnownErrors")

	_, resp, err = client.ProblemsAPI.GetProblemsManagedByMe(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetProblemsManagedByMe")

	_, resp, err = client.ProblemsAPI.GetProblemsProgressHalted(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetProblemsProgressHalted")

	_, resp, err = client.ProblemsAPI.GetProblemsSolved(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetProblemsSolved")

	_, resp, err = client.PeopleAPI.GetPeopleDirectory(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetPeopleDirectory")
}

func TestCoverage_RequestsFiltersAndNotes(t *testing.T) {
	client := coverageClient(t)
	ctx := context.Background()

	_, resp, err := client.RequestsAPI.GetRequestsWaitingForMe(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetRequestsWaitingForMe")

	_, resp, err = client.RequestsAPI.GetRequestsProblemManagementReview(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetRequestsProblemManagementReview")

	_, resp, err = client.RequestsAPI.GetRequestsSlaAccountability(ctx).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetRequestsSlaAccountability")

	rid, ok := firstOpenRequestID(t, ctx, client)
	if !ok {
		t.Log("no open request id; skipping request-scoped GETs")
		return
	}

	_, resp, err = client.RequestsAPI.GetRequestsIdNotesPublic(ctx, rid).PerPage(1).Execute()
	assertReadOK(t, resp, err, "GetRequestsIdNotesPublic")

	_, resp, err = client.RequestsAPI.GetRequestsIdNotesInternal(ctx, rid).PerPage(1).Execute()
	assertReadOK(t, resp, err, "GetRequestsIdNotesInternal")

	_, resp, err = client.RequestsAPI.GetRequestsIdCisActive(ctx, rid).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetRequestsIdCisActive")

	_, resp, err = client.RequestsAPI.GetRequestsIdCisInactive(ctx, rid).PerPage(1).Fields("id").Execute()
	assertReadOK(t, resp, err, "GetRequestsIdCisInactive")
}

func TestCoverage_AccountUsageStatementByID(t *testing.T) {
	client := coverageClient(t)
	ctx := context.Background()

	id, ok := firstUsageStatementID(t, ctx, client)
	if !ok {
		t.Log("no usage statements list or not account owner; skipping GetAccountUsageStatementsId")
		return
	}

	_, resp, err := client.AccountAPI.GetAccountUsageStatementsId(ctx, id).Execute()
	assertReadOK(t, resp, err, "GetAccountUsageStatementsId")
}

func TestCoverage_ProductCisByExistingProduct(t *testing.T) {
	client := coverageClient(t)
	ctx := context.Background()

	rows, resp, err := client.ProductsAPI.GetProductsEnabled(ctx).PerPage(1).Fields("id").Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 {
		if resp != nil && resp.StatusCode == 403 {
			t.Log("GetProductsEnabled 403; skip")
			return
		}
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	}
	if len(rows) == 0 {
		t.Log("no enabled products; skip GetProductsIdCis")
		return
	}
	raw, ok := rows[0]["id"]
	require.True(t, ok)
	var pid int32
	switch v := raw.(type) {
	case float64:
		pid = int32(v)
	case int:
		pid = int32(v)
	case int32:
		pid = v
	default:
		t.Fatalf("unexpected id type %T", raw)
	}

	resp, err = client.ProductsAPI.GetProductsIdCis(ctx, pid).PerPage(1).Fields("id").Execute()
	if err != nil {
		require.NotNil(t, resp)
		if resp.StatusCode == 403 {
			t.Log("GetProductsIdCis 403; skip")
			return
		}
		require.NoError(t, err)
	}
	require.Equal(t, 200, resp.StatusCode)
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
}
