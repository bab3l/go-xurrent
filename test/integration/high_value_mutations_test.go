//go:build integration

// Opt-in POST/PATCH for high-value configuration resources (teams, sites, services, SLAs).
// Uses the same gate as other mutations: XURRENT_ALLOW_MUTATIONS=1.
//
// Creates minimal records, verifies GET by id + PATCH, then disables via PATCH (no archive/trash
// endpoints in spec for some resources). Chained config tests cover calendars, service offerings,
// request templates, and (optionally) workflows; see TestMutation_Chained_* below.
//
//   go test -tags=integration ./test/integration/ -run TestMutation_HighValue_ -v -count=1
//   go test -tags=integration ./test/integration/ -run TestMutation_Chained_ -v -count=1
//   go test -tags=integration ./test/integration/ -run TestMutation_CIReservationAutomation -v -count=1

package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/xurrent/go-xurrent"
)

func firstOrganizationID(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.OrganizationsAPI.GetOrganizations(ctx).PerPage(1).Fields("id").Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 || len(rows) == 0 {
		return 0, false
	}
	if rows[0].Id == nil {
		return 0, false
	}
	return int32(*rows[0].Id), true
}

func mapIDInt32(m map[string]interface{}) (int32, bool) {
	if m == nil {
		return 0, false
	}
	raw, ok := m["id"]
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
	default:
		return 0, false
	}
}

// TestMutation_HighValue_TeamSiteServiceSLA chains minimal creates when an organization exists.
func TestMutation_HighValue_TeamSiteServiceSLA(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")

	orgID, ok := firstOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need at least one organization for provider/customer references")
	}

	sfx := randomSuffix()

	// --- Team ---
	teamName := fmt.Sprintf("sdk-hv-team-%s", sfx)
	teamCreated, resp, err := client.TeamsAPI.PostTeams(ctx).Body(map[string]interface{}{
		"name": teamName,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, teamCreated)
	tid, ok := mapIDInt32Any(teamCreated)
	require.True(t, ok, "team id")
	defer func() {
		_, _, _ = client.TeamsAPI.PatchTeamsId(context.Background(), tid).Body(map[string]interface{}{
			"disabled": true,
			"remarks":  fmt.Sprintf("disabled after sdk high-value test %s", sfx),
		}).Execute()
	}()

	_, resp, err = client.TeamsAPI.GetTeamsId(ctx, tid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.TeamsAPI.PatchTeamsId(ctx, tid).Body(map[string]interface{}{
		"remarks": fmt.Sprintf("sdk patch %s", sfx),
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	// --- Site ---
	siteName := fmt.Sprintf("sdk-hv-site-%s", sfx)
	siteCreated, resp, err := client.SitesAPI.PostSites(ctx).Body(map[string]interface{}{
		"name": siteName,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, siteCreated)
	sid, ok := mapIDInt32Any(siteCreated)
	require.True(t, ok, "site id")
	defer func() {
		_, _, _ = client.SitesAPI.PatchSitesId(context.Background(), sid).Body(map[string]interface{}{
			"disabled": true,
			"remarks":  fmt.Sprintf("disabled after sdk high-value test %s", sfx),
		}).Execute()
	}()

	_, resp, err = client.SitesAPI.GetSitesId(ctx, sid).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()

	_, resp, err = client.SitesAPI.PatchSitesId(ctx, sid).Body(map[string]interface{}{
		"remarks": fmt.Sprintf("sdk patch %s", sfx),
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	// --- Service (requires provider organization) ---
	svcName := fmt.Sprintf("sdk-hv-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":     svcName,
		"provider": map[string]interface{}{"id": orgID},
	}).Execute()
	if err != nil {
		t.Logf("PostServices skipped or failed (permissions/template constraints): %v", err)
	} else {
		require.Equal(t, 201, resp.StatusCode)
		requireResourceOwnedByEnvAccountAny(t, acc, svcCreated)
		svcID, ok := mapIDInt32Any(svcCreated)
		if ok {
			defer func() {
				_, _, _ = client.ServicesAPI.PatchServicesId(context.Background(), svcID).Body(map[string]interface{}{
					"disabled": true,
					"remarks":  fmt.Sprintf("disabled after sdk high-value test %s", sfx),
				}).Execute()
			}()
			_, resp, err = client.ServicesAPI.GetServicesId(ctx, svcID).Execute()
			require.NoError(t, err)
			require.Equal(t, 200, resp.StatusCode)
			_, resp, err = client.ServicesAPI.PatchServicesId(ctx, svcID).Body(map[string]interface{}{
				"remarks": fmt.Sprintf("sdk patch %s", sfx),
			}).Execute()
			require.NoError(t, err)
			require.Equal(t, 200, resp.StatusCode)
		}
	}

	// --- SLA (customer org required; draft-style to avoid customer rep requirements for active SLAs) ---
	slaName := fmt.Sprintf("sdk-hv-sla-%s", sfx)
	slaBody := map[string]interface{}{
		"name":     slaName,
		"customer": map[string]interface{}{"id": orgID},
	}
	slaCreated, resp, err := client.ServiceLevelAgreementsAPI.PostSlas(ctx).Body(slaBody).Execute()
	if err != nil {
		t.Logf("PostSlas skipped or failed (validation/permissions): %v", err)
	} else {
		require.Equal(t, 201, resp.StatusCode)
		requireResourceOwnedByEnvAccountAny(t, acc, slaCreated)
		slaID, ok := mapIDInt32Any(slaCreated)
		require.True(t, ok, "sla id")
		defer func() {
			_, _, _ = client.ServiceLevelAgreementsAPI.PatchSlasId(context.Background(), slaID).Body(map[string]interface{}{
				"disabled": true,
				"remarks":  fmt.Sprintf("disabled after sdk high-value test %s", sfx),
			}).Execute()
		}()

		_, resp, err = client.ServiceLevelAgreementsAPI.GetSlasId(ctx, slaID).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()

		_, resp, err = client.ServiceLevelAgreementsAPI.PatchSlasId(ctx, slaID).Body(map[string]interface{}{
			"remarks": fmt.Sprintf("sdk patch %s", sfx),
		}).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	}
}

// TestMutation_Chained_CalendarServiceOfferingRequestTemplate chains POST/PATCH across calendar,
// service, service offering, and request template, then disables created records (dependency order).
func TestMutation_Chained_CalendarServiceOfferingRequestTemplate(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")

	orgID, ok := firstOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need at least one organization for service provider reference")
	}

	sfx := randomSuffix()
	var calID, svcID, offID, rtID int32
	defer func() {
		bctx := context.Background()
		if rtID != 0 {
			_, _, _ = client.RequestTemplatesAPI.PatchRequestTemplatesId(bctx, rtID).Body(map[string]interface{}{
				"disabled": true,
				"remarks":  fmt.Sprintf("disabled after sdk chained config test %s", sfx),
			}).Execute()
		}
		if offID != 0 {
			_, _, _ = client.ServiceOfferingsAPI.PatchServiceOfferingsId(bctx, offID).Body(map[string]interface{}{
				"disabled": true,
				"remarks":  fmt.Sprintf("disabled after sdk chained config test %s", sfx),
			}).Execute()
		}
		if svcID != 0 {
			_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
				"disabled": true,
				"remarks":  fmt.Sprintf("disabled after sdk chained config test %s", sfx),
			}).Execute()
		}
		if calID != 0 {
			_, _, _ = client.CalendarsAPI.PatchCalendarsId(bctx, calID).Body(map[string]interface{}{
				"disabled": true,
				"remarks":  fmt.Sprintf("disabled after sdk chained config test %s", sfx),
			}).Execute()
		}
	}()

	calName := fmt.Sprintf("sdk-chain-cal-%s", sfx)
	calCreated, resp, err := client.CalendarsAPI.PostCalendars(ctx).Body(map[string]interface{}{
		"name": calName,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, calCreated)
	var okID bool
	calID, okID = mapIDInt32Any(calCreated)
	require.True(t, okID, "calendar id")

	_, resp, err = client.CalendarsAPI.PatchCalendarsId(ctx, calID).Body(map[string]interface{}{
		"remarks": fmt.Sprintf("sdk patch cal %s", sfx),
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	svcName := fmt.Sprintf("sdk-chain-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":     svcName,
		"provider": map[string]interface{}{"id": orgID},
	}).Execute()
	if err != nil && resp != nil && (resp.StatusCode == 401 || resp.StatusCode == 403) {
		t.Skipf("PostServices not permitted for this token/account (HTTP %d)", resp.StatusCode)
	}
	if err != nil && resp != nil && resp.StatusCode == 422 {
		t.Skipf("PostServices validation failed (HTTP 422); tenant may require extra fields: %v", err)
	}
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, svcCreated)
	svcID, okID = mapIDInt32Any(svcCreated)
	require.True(t, okID, "service id")

	_, resp, err = client.ServicesAPI.GetServicesId(ctx, svcID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	offName := fmt.Sprintf("sdk-chain-offering-%s", sfx)
	offCreated, resp, err := client.ServiceOfferingsAPI.PostServiceOfferings(ctx).Body(map[string]interface{}{
		"name":       offName,
		"service_id": svcID,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, offCreated)
	offID, okID = mapIDInt32Any(offCreated)
	require.True(t, okID, "service offering id")

	_, offResp, err := client.ServiceOfferingsAPI.GetServiceOfferingsId(ctx, offID).Execute()
	require.NoError(t, err)
	require.NotNil(t, offResp)
	require.Equal(t, 200, offResp.StatusCode)
	_, _ = io.Copy(io.Discard, offResp.Body)
	_ = offResp.Body.Close()

	_, resp, err = client.ServiceOfferingsAPI.PatchServiceOfferingsId(ctx, offID).Body(map[string]interface{}{
		"remarks": fmt.Sprintf("sdk patch offering %s", sfx),
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	rtSubject := fmt.Sprintf("sdk-chain-rt-%s", sfx)
	rtCreated, resp, err := client.RequestTemplatesAPI.PostRequestTemplates(ctx).Body(map[string]interface{}{
		"subject":    rtSubject,
		"service_id": svcID,
		"category":   "rfi",
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, rtCreated)
	rtID, okID = mapIDInt32Any(rtCreated)
	require.True(t, okID, "request template id")

	_, resp, err = client.RequestTemplatesAPI.GetRequestTemplatesId(ctx, rtID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.RequestTemplatesAPI.PatchRequestTemplatesId(ctx, rtID).Body(map[string]interface{}{
		"remarks": fmt.Sprintf("sdk patch rt %s", sfx),
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}

// TestMutation_Chained_WorkflowCreateTrash creates a workflow from XURRENT_TEST_WORKFLOW_TEMPLATE_ID,
// patches it, trashes it, and disables the temporary service. Skips when the template id is unset.
func TestMutation_Chained_WorkflowCreateTrash(t *testing.T) {
	tplStr := os.Getenv("XURRENT_TEST_WORKFLOW_TEMPLATE_ID")
	if tplStr == "" {
		t.Skip("set XURRENT_TEST_WORKFLOW_TEMPLATE_ID (numeric id in your tenant) to run workflow create+trash")
	}
	tpl64, err := strconv.ParseInt(tplStr, 10, 32)
	require.NoError(t, err)
	templateID := int32(tpl64)

	wfType := os.Getenv("XURRENT_TEST_WORKFLOW_TYPE")
	if wfType == "" {
		wfType = "application_change"
	}

	client := mutationClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")
	orgID, ok := firstOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need at least one organization for service provider reference")
	}

	sfx := randomSuffix()
	var workflowID, svcID int32
	defer func() {
		bctx := context.Background()
		if workflowID != 0 {
			_, _, _ = client.WorkflowsAPI.PostWorkflowsIdTrash(bctx, workflowID).Execute()
		}
		if svcID != 0 {
			_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
				"disabled": true,
				"remarks":  fmt.Sprintf("disabled after sdk workflow chain test %s", sfx),
			}).Execute()
		}
	}()

	svcName := fmt.Sprintf("sdk-wf-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":     svcName,
		"provider": map[string]interface{}{"id": orgID},
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, acc, svcCreated)
	svcID, ok = mapIDInt32Any(svcCreated)
	require.True(t, ok, "service id")

	managerID := mePersonID(t, ctx, client)
	wfSubject := fmt.Sprintf("sdk-wf-%s", sfx)
	wfBody := map[string]interface{}{
		"subject":           wfSubject,
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

	_, resp, err = client.WorkflowsAPI.GetWorkflowsId(ctx, workflowID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.WorkflowsAPI.PatchWorkflowsId(ctx, workflowID).Body(map[string]interface{}{
		"remarks": fmt.Sprintf("sdk patch wf %s", sfx),
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.WorkflowsAPI.PostWorkflowsIdTrash(ctx, workflowID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	workflowID = 0
}

const envHTTPMinIntervalKey = "XURRENT_HTTP_MIN_INTERVAL"

func mutationClientWithThrottle(t *testing.T) *openapiclient.APIClient {
	t.Helper()
	if os.Getenv(envHTTPMinIntervalKey) == "" {
		t.Setenv(envHTTPMinIntervalKey, "0.4")
	}
	return mutationClient(t)
}

// firstProductIDForPostCI prefers enabled products, then disabled, then generic product list.
func firstProductIDForPostCI(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.ProductsAPI.GetProductsEnabled(ctx).PerPage(1).Fields("id").Execute()
	if err == nil && resp != nil && resp.StatusCode == 200 && len(rows) > 0 {
		if id, ok := mapIDInt32(rows[0]); ok {
			return id, true
		}
	}
	rowsD, respD, errD := client.ProductsAPI.GetProductsDisabled(ctx).PerPage(1).Fields("id").Execute()
	if errD == nil && respD != nil && respD.StatusCode == 200 && len(rowsD) > 0 {
		if id, ok := int32FromOptionalFloat32ID(rowsD[0].Id); ok {
			return id, true
		}
	}
	rowsG, respG, errG := client.ProductsAPI.GetProducts(ctx).PerPage(1).Fields("id").Execute()
	if errG != nil || respG == nil || respG.StatusCode != 200 || len(rowsG) == 0 {
		return 0, false
	}
	return int32FromOptionalFloat32ID(rowsG[0].Id)
}

func firstServiceIDForOrg(t *testing.T, ctx context.Context, client *openapiclient.APIClient, orgID int32) (int32, bool) {
	t.Helper()
	rows, resp, err := client.ServicesAPI.GetServices(ctx).Provider(orgID).PerPage(1).Fields("id").Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 || len(rows) == 0 {
		return 0, false
	}
	return mapIDInt32(rows[0])
}

func firstAnyServiceID(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.ServicesAPI.GetServices(ctx).PerPage(25).Fields("id").Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 || len(rows) == 0 {
		return 0, false
	}
	for _, row := range rows {
		if id, ok := mapIDInt32(row); ok {
			return id, true
		}
	}
	return 0, false
}

func readJSONBodyMap(t *testing.T, resp *http.Response) map[string]interface{} {
	t.Helper()
	if resp == nil || resp.Body == nil {
		return nil
	}
	defer func() { _ = resp.Body.Close() }()
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	var m map[string]interface{}
	if json.Unmarshal(b, &m) != nil {
		return nil
	}
	return m
}

// postServiceInstanceForCITest creates a service instance (see https://developer.xurrent.com/v1/service_instances/).
func postServiceInstanceForCITest(t *testing.T, ctx context.Context, client *openapiclient.APIClient, svcID, teamID int32, name string) int32 {
	t.Helper()
	created, resp, err := client.ServiceInstancesAPI.PostServiceInstances(ctx).Body(map[string]interface{}{
		"name":         name,
		"service":      map[string]interface{}{"id": svcID},
		"support_team": map[string]interface{}{"id": teamID},
		"status":       "in_production",
	}).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 201, resp.StatusCode, "PostServiceInstances")
	id, ok := mapIDInt32Any(created)
	require.True(t, ok, "service instance id from PostServiceInstances")
	return id
}

func patchServiceInstanceStatusForCITest(t *testing.T, ctx context.Context, client *openapiclient.APIClient, siID int32, status string) {
	t.Helper()
	_, resp, err := client.ServiceInstancesAPI.PatchServiceInstancesId(ctx, siID).Body(map[string]interface{}{
		"status": status,
	}).Execute()
	if err != nil {
		t.Logf("PatchServiceInstancesId: %v", err)
		return
	}
	if resp != nil && resp.StatusCode != 200 {
		t.Logf("PatchServiceInstancesId: HTTP %d", resp.StatusCode)
	}
}

// archiveCreatedCI prefers POST /cis/:id/archive (see Configuration items API), then PATCH status=archived.
// Trash/restore (PostCisIdTrash, PostCisIdRestore) are documented alternatives; not used here to avoid extra admin-only calls.
func archiveCreatedCI(t *testing.T, ctx context.Context, client *openapiclient.APIClient, ciID int32, sfx string) {
	t.Helper()
	_, resp, err := client.ConfigurationItemsAPI.PostCisIdArchive(ctx, ciID).Execute()
	if err == nil && resp != nil && resp.StatusCode == 200 {
		t.Logf("CI id=%d archived (PostCisIdArchive)", ciID)
		return
	}
	if err != nil {
		t.Logf("PostCisIdArchive: %v", err)
	} else if resp != nil {
		t.Logf("PostCisIdArchive: HTTP %d", resp.StatusCode)
	}
	_, resp2, err2 := client.ConfigurationItemsAPI.PatchCisId(ctx, ciID).Body(map[string]interface{}{
		"status":  "archived",
		"remarks": fmt.Sprintf("archived after integration test %s", sfx),
	}).Execute()
	if err2 != nil {
		t.Logf("PatchCisId archive fallback: %v", err2)
		return
	}
	if resp2 != nil && resp2.StatusCode == 200 {
		t.Logf("CI id=%d archived (PatchCisId status)", ciID)
		return
	}
	if resp2 != nil {
		t.Logf("PatchCisId archive fallback: HTTP %d", resp2.StatusCode)
	}
}

// TestMutation_CIReservationAutomation creates a service (POST /v1/services), a service instance (POST /v1/service_instances),
// then POSTs a CI linked to both, GETs CI subresources, POSTs a reservation, GETs it, PATCHes reservation to canceled,
// archives the CI (PostCisIdArchive or PatchCisId), discontinues the service instance, and disables the service.
// Automation rules: OpenAPI has no POST; we GET list → GetAutomationRulesId when present.
func TestMutation_CIReservationAutomation(t *testing.T) {
	client := mutationClientWithThrottle(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")
	sfx := randomSuffix()

	orgID, ok := firstOrganizationID(t, ctx, client)
	require.True(t, ok, "need an organization")

	teamID, ok := firstTeamIDFromAPI(t, ctx, client)
	require.True(t, ok, "need a team for support_team_id")

	productID, ok := firstProductIDForPostCI(t, ctx, client)
	require.True(t, ok, "need a product (enabled, disabled, or from list)")

	var svcID int32
	var createdService bool

	svcName := fmt.Sprintf("sdk-ci-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":     svcName,
		"provider": map[string]interface{}{"id": orgID},
	}).Execute()
	if err != nil && resp != nil && (resp.StatusCode == 401 || resp.StatusCode == 403) {
		t.Skipf("PostServices not permitted (HTTP %d); need a service for CI + service instance chain", resp.StatusCode)
	}
	if err != nil && resp != nil && resp.StatusCode == 422 {
		t.Skipf("PostServices validation failed (HTTP 422): %v", err)
	}
	if err != nil || resp == nil {
		t.Logf("PostServices failed: %v", err)
		var ok2 bool
		svcID, ok2 = firstServiceIDForOrg(t, ctx, client, orgID)
		if !ok2 {
			svcID, ok2 = firstAnyServiceID(t, ctx, client)
		}
		if !ok2 {
			if id, ok3 := int32FromEnvOrFixture("XURRENT_TEST_SERVICE_ID"); ok3 {
				svcID, ok2 = id, true
				t.Logf("using XURRENT_TEST_SERVICE_ID=%d", svcID)
			}
		}
		if !ok2 {
			t.Skip("no service: PostServices failed and no service from GET /v1/services or XURRENT_TEST_SERVICE_ID")
		}
	} else {
		require.Equal(t, 201, resp.StatusCode)
		requireResourceOwnedByEnvAccountAny(t, acc, svcCreated)
		var okID bool
		svcID, okID = mapIDInt32Any(svcCreated)
		require.True(t, okID, "service id from PostServices")
		createdService = true
	}

	siName := fmt.Sprintf("sdk-ci-si-%s", sfx)
	siID := postServiceInstanceForCITest(t, ctx, client, svcID, teamID, siName)

	_, siGetResp, err := client.ServiceInstancesAPI.GetServiceInstancesId(ctx, siID).Execute()
	require.NoError(t, err)
	require.NotNil(t, siGetResp)
	require.Equal(t, 200, siGetResp.StatusCode, "GetServiceInstancesId")
	_, _ = io.Copy(io.Discard, siGetResp.Body)
	_ = siGetResp.Body.Close()

	personID := mePersonID(t, ctx, client)

	ciBody := map[string]interface{}{
		"source":           "4me",
		"systemID":         fmt.Sprintf("sdk-ci-%s", sfx),
		"label":            fmt.Sprintf("LAB-%s", sfx),
		"name":             fmt.Sprintf("integration-ci-%s", sfx),
		"status":           "in_production",
		"product_id":       productID,
		"service":          svcID,
		"service_instance": map[string]interface{}{"id": siID},
		"support_team_id":  teamID,
	}

	postCIResp, err := client.ConfigurationItemsAPI.PostCis(ctx).Body(ciBody).Execute()
	require.NoError(t, err)
	require.NotNil(t, postCIResp)
	if postCIResp.StatusCode != 200 && postCIResp.StatusCode != 201 {
		b, _ := io.ReadAll(postCIResp.Body)
		_ = postCIResp.Body.Close()
		t.Fatalf("PostCis: HTTP %d body %s", postCIResp.StatusCode, string(b))
	}
	ciCreated := readJSONBodyMap(t, postCIResp)
	requireResourceOwnedByEnvAccountAny(t, acc, ciCreated)
	ciID, ok := mapIDInt32(ciCreated)
	require.True(t, ok, "ci id from PostCis body")
	t.Logf("created CI id=%d (service=%d service_instance=%d)", ciID, svcID, siID)

	_, getResp, err := client.ConfigurationItemsAPI.GetCisId(ctx, ciID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, getResp.StatusCode)
	_, _ = io.Copy(io.Discard, getResp.Body)
	_ = getResp.Body.Close()

	_, relResp, err := client.ConfigurationItemRelationsAPI.GetCisIdCiRelations(ctx, ciID).PerPage(5).Fields("id").Execute()
	if err == nil && relResp != nil && relResp.StatusCode == 200 {
		_, _ = io.Copy(io.Discard, relResp.Body)
		_ = relResp.Body.Close()
	} else {
		t.Logf("GetCisIdCiRelations: %v", err)
	}

	_, uResp, err := client.ConfigurationItemRelationsAPI.GetCisIdUsers(ctx, ciID).PerPage(5).Fields("id").Execute()
	if err == nil && uResp != nil && uResp.StatusCode == 200 {
		_, _ = io.Copy(io.Discard, uResp.Body)
		_ = uResp.Body.Close()
	} else {
		t.Logf("GetCisIdUsers: %v", err)
	}

	start := time.Now().UTC().Add(48 * time.Hour).Truncate(time.Minute)
	end := start.Add(2 * time.Hour)
	resBody := map[string]interface{}{
		"ci":       map[string]interface{}{"id": ciID},
		"person":   map[string]interface{}{"id": personID},
		"name":     fmt.Sprintf("int-res-%s", sfx),
		"start_at": start.Format(time.RFC3339),
		"end_at":   end.Format(time.RFC3339),
		"status":   "confirmed",
	}

	postResHTTP, err := client.ReservationsAPI.PostReservations(ctx).Body(resBody).Execute()
	require.NoError(t, err)
	require.NotNil(t, postResHTTP)
	if postResHTTP.StatusCode != 201 {
		b, _ := io.ReadAll(postResHTTP.Body)
		_ = postResHTTP.Body.Close()
		t.Fatalf("PostReservations: status %d body %s", postResHTTP.StatusCode, string(b))
	}
	resCreated := readJSONBodyMap(t, postResHTTP)
	resID, ok := mapIDInt32(resCreated)
	require.True(t, ok, "reservation id")
	t.Logf("created reservation id=%d", resID)

	defer func() {
		bctx := context.Background()
		if resID != 0 {
			patchResp, perr := client.ReservationsAPI.PatchReservationsId(bctx, resID).Body(map[string]interface{}{
				"status": "canceled",
			}).Execute()
			if perr != nil {
				t.Logf("PatchReservationsId cancel: %v", perr)
			} else if patchResp != nil {
				_, _ = io.Copy(io.Discard, patchResp.Body)
				_ = patchResp.Body.Close()
				t.Logf("canceled reservation id=%d", resID)
			}
		}
		if ciID != 0 {
			archiveCreatedCI(t, bctx, client, ciID, sfx)
		}
		if siID != 0 {
			patchServiceInstanceStatusForCITest(t, bctx, client, siID, "discontinued")
		}
		if createdService && svcID != 0 {
			_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
				"disabled": true,
				"remarks":  fmt.Sprintf("disabled after CI mutation test %s", sfx),
			}).Execute()
		}
	}()

	oneRes, err := client.ReservationsAPI.GetReservationsId(ctx, resID).Execute()
	require.NoError(t, err)
	require.NotNil(t, oneRes)
	require.Equal(t, 200, oneRes.StatusCode)
	_, _ = io.Copy(io.Discard, oneRes.Body)
	_ = oneRes.Body.Close()

	arListHTTP, err := client.AutomationRulesAPI.GetAutomationRules(ctx).PerPage(5).Fields("id").Execute()
	require.NoError(t, err)
	require.NotNil(t, arListHTTP)
	require.Equal(t, 200, arListHTTP.StatusCode)
	arBody, err := io.ReadAll(arListHTTP.Body)
	_ = arListHTTP.Body.Close()
	require.NoError(t, err)
	var arRows []map[string]interface{}
	if json.Unmarshal(arBody, &arRows) == nil && len(arRows) > 0 {
		if aid, ok := mapIDInt32(arRows[0]); ok {
			arGet, err := client.AutomationRulesAPI.GetAutomationRulesId(ctx, aid).Execute()
			require.NoError(t, err)
			require.NotNil(t, arGet)
			require.Equal(t, 200, arGet.StatusCode)
			_, _ = io.Copy(io.Discard, arGet.Body)
			_ = arGet.Body.Close()
			t.Logf("GetAutomationRulesId id=%d OK", aid)
		}
	} else {
		t.Log("no automation rules in list; skip GetAutomationRulesId")
	}
}
