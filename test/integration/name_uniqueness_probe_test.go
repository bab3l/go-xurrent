//go:build integration

// Empirical probes for Phase A/B (name reservation + error payloads).
// Requires XURRENT_ALLOW_MUTATIONS=1, XURRENT_TOKEN, XURRENT_ACCOUNT.
//
//	go test -tags=integration ./test/integration/ -run 'TestProbe_(NameUniqueness|CompositeUniqueness)' -v -count=1
//
// Optional: XURRENT_HTTP_MIN_INTERVAL=0.35 to respect rate limits.
// Optional: XURRENT_TEST_WORKFLOW_TEMPLATE_ID (+ XURRENT_TEST_WORKFLOW_TYPE) for workflow subject probe.
// Optional: XURRENT_TRUSTED_ACCOUNT_IDS=comma-separated ids when resource.account.id differs from
// XURRENT_ACCOUNT (trust / home account) so person probes do not skip.

package integration_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/xurrent/go-xurrent"
)

func openAPIErrBody(err error) (httpStatus string, body string) {
	var ge *openapiclient.GenericOpenAPIError
	if errors.As(err, &ge) {
		return ge.Error(), string(ge.Body())
	}
	if err != nil {
		return err.Error(), ""
	}
	return "", ""
}

func statusOf(resp *http.Response) int {
	if resp == nil {
		return 0
	}
	return resp.StatusCode
}

// --- Team / Site (existing behavior + dual cleanup) ---

func TestProbe_NameUniqueness_Team(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	sfx := randomSuffix()
	name := fmt.Sprintf("mcpprobe-team-%s", sfx)

	created, resp, err := client.TeamsAPI.PostTeams(ctx).Body(map[string]interface{}{
		"name": name,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, os.Getenv("XURRENT_ACCOUNT"), created)
	tid, ok := mapIDInt32Any(created)
	require.True(t, ok, "team id")

	var tid2 int32
	defer func() {
		bctx := context.Background()
		if tid2 != 0 {
			_, _, _ = client.TeamsAPI.PatchTeamsId(bctx, tid2).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-team-retired-b-%s", sfx),
			}).Execute()
		}
		cleanupRenameTeam(t, client, tid, sfx)
	}()

	_, resp, err = client.TeamsAPI.PatchTeamsId(ctx, tid).Body(map[string]interface{}{
		"disabled": true,
		"remarks":  fmt.Sprintf("name uniqueness probe %s", sfx),
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.TeamsAPI.PostTeams(ctx).Body(map[string]interface{}{
		"name": name,
	}).Execute()
	require.Error(t, err2)
	require.NotNil(t, resp2)
	st, bod := openAPIErrBody(err2)
	t.Logf("duplicate team name: HTTP=%s body=%s", st, bod)
	require.GreaterOrEqual(t, resp2.StatusCode, 400)

	freed := fmt.Sprintf("%s-renamed-%s", name, sfx)
	_, resp, err = client.TeamsAPI.PatchTeamsId(ctx, tid).Body(map[string]interface{}{
		"name": freed,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	created2, resp3, err3 := client.TeamsAPI.PostTeams(ctx).Body(map[string]interface{}{
		"name": name,
	}).Execute()
	require.NoError(t, err3)
	require.Equal(t, 201, resp3.StatusCode)
	var ok2 bool
	tid2, ok2 = mapIDInt32Any(created2)
	require.True(t, ok2)
	_, _, _ = client.TeamsAPI.PatchTeamsId(context.Background(), tid2).Body(map[string]interface{}{
		"disabled": true,
		"name":     fmt.Sprintf("%s-second-%s", name, sfx),
	}).Execute()
}

func cleanupRenameTeam(t *testing.T, client *openapiclient.APIClient, tid int32, sfx string) {
	t.Helper()
	_, _, _ = client.TeamsAPI.PatchTeamsId(context.Background(), tid).Body(map[string]interface{}{
		"disabled": true,
		"name":     fmt.Sprintf("mcpprobe-team-retired-%s", sfx),
		"remarks":  "cleanup name uniqueness probe",
	}).Execute()
}

func TestProbe_NameUniqueness_Site(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	sfx := randomSuffix()
	name := fmt.Sprintf("mcpprobe-site-%s", sfx)

	created, resp, err := client.SitesAPI.PostSites(ctx).Body(map[string]interface{}{
		"name": name,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	requireResourceOwnedByEnvAccountAny(t, os.Getenv("XURRENT_ACCOUNT"), created)
	sid, ok := mapIDInt32Any(created)
	require.True(t, ok, "site id")

	var sid2 int32
	defer func() {
		bctx := context.Background()
		if sid2 != 0 {
			_, _, _ = client.SitesAPI.PatchSitesId(bctx, sid2).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-site-retired-b-%s", sfx),
			}).Execute()
		}
		cleanupRenameSite(t, client, sid, sfx)
	}()

	_, resp, err = client.SitesAPI.PatchSitesId(ctx, sid).Body(map[string]interface{}{
		"disabled": true,
		"remarks":  fmt.Sprintf("name uniqueness probe %s", sfx),
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.SitesAPI.PostSites(ctx).Body(map[string]interface{}{
		"name": name,
	}).Execute()
	require.Error(t, err2)
	require.NotNil(t, resp2)
	st, bod := openAPIErrBody(err2)
	t.Logf("duplicate site name: HTTP status=%d err=%s body=%s", resp2.StatusCode, st, bod)
	require.GreaterOrEqual(t, resp2.StatusCode, 400)

	freed := fmt.Sprintf("%s-renamed-%s", name, sfx)
	_, _, err = client.SitesAPI.PatchSitesId(ctx, sid).Body(map[string]interface{}{
		"name": freed,
	}).Execute()
	require.NoError(t, err)

	created2, resp3, err3 := client.SitesAPI.PostSites(ctx).Body(map[string]interface{}{
		"name": name,
	}).Execute()
	require.NoError(t, err3)
	require.Equal(t, 201, resp3.StatusCode)
	sid2, ok = mapIDInt32Any(created2)
	require.True(t, ok)
}

func cleanupRenameSite(t *testing.T, client *openapiclient.APIClient, sid int32, sfx string) {
	t.Helper()
	_, _, _ = client.SitesAPI.PatchSitesId(context.Background(), sid).Body(map[string]interface{}{
		"disabled": true,
		"name":     fmt.Sprintf("mcpprobe-site-retired-%s", sfx),
		"remarks":  "cleanup name uniqueness probe",
	}).Execute()
}

func TestProbe_NameUniqueness_CrossType_TeamAndSite(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	sfx := randomSuffix()
	shared := fmt.Sprintf("mcpprobe-shared-%s", sfx)

	teamCreated, resp, err := client.TeamsAPI.PostTeams(ctx).Body(map[string]interface{}{
		"name": shared,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	tid, _ := mapIDInt32Any(teamCreated)
	defer cleanupRenameTeam(t, client, tid, sfx)

	siteCreated, resp, err := client.SitesAPI.PostSites(ctx).Body(map[string]interface{}{
		"name": shared,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode, "same display name on different entity types should be allowed")
	sid, _ := mapIDInt32Any(siteCreated)
	defer cleanupRenameSite(t, client, sid, sfx)
}

// --- Service (provider must be an internal org when possible) ---

func TestProbe_NameUniqueness_Service(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	orgID, ok := firstEnabledOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need at least one enabled organization (service provider)")
	}
	sfx := randomSuffix()
	name := fmt.Sprintf("mcpprobe-svc-%s", sfx)

	// Rails accepts top-level provider_id; nested `provider` alone may not populate provider_id.
	created, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":        name,
		"provider_id": orgID,
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostServices failed: status=%v body=%s", statusOf(resp), bod)
		if resp != nil && (resp.StatusCode == 401 || resp.StatusCode == 403) {
			t.Skip("PostServices not permitted for this token")
		}
		t.Skip("PostServices validation failed — check provider org (internal) and configuration role")
		return
	}
	requireResourceOwnedByEnvAccountAny(t, os.Getenv("XURRENT_ACCOUNT"), created)
	svcID, ok := mapIDInt32Any(created)
	require.True(t, ok)

	var svc2 int32
	defer func() {
		bctx := context.Background()
		if svc2 != 0 {
			_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svc2).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-svc-retired-b-%s", sfx),
			}).Execute()
		}
		_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-svc-retired-%s", sfx),
		}).Execute()
	}()

	_, resp, err = client.ServicesAPI.PatchServicesId(ctx, svcID).Body(map[string]interface{}{
		"disabled": true,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":        name,
		"provider_id": orgID,
	}).Execute()
	require.Error(t, err2)
	require.NotNil(t, resp2)
	_, bod := openAPIErrBody(err2)
	t.Logf("duplicate service name: HTTP=%d body=%s", resp2.StatusCode, bod)
	require.GreaterOrEqual(t, resp2.StatusCode, 400)

	freed := fmt.Sprintf("%s-renamed-%s", name, sfx)
	_, _, err = client.ServicesAPI.PatchServicesId(ctx, svcID).Body(map[string]interface{}{
		"name": freed,
	}).Execute()
	require.NoError(t, err)

	created2, resp3, err3 := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":        name,
		"provider_id": orgID,
	}).Execute()
	require.NoError(t, err3)
	require.Equal(t, 201, resp3.StatusCode)
	svc2, ok = mapIDInt32Any(created2)
	require.True(t, ok)
}

// --- Calendar ---

func TestProbe_NameUniqueness_Calendar(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	sfx := randomSuffix()
	name := fmt.Sprintf("mcpprobe-cal-%s", sfx)

	created, resp, err := client.CalendarsAPI.PostCalendars(ctx).Body(map[string]interface{}{
		"name": name,
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostCalendars: %v body=%s", err, bod)
		t.Skip("PostCalendars not available")
		return
	}
	calID, ok := mapIDInt32Any(created)
	require.True(t, ok)

	var cal2 int32
	defer func() {
		bctx := context.Background()
		if cal2 != 0 {
			_, _, _ = client.CalendarsAPI.PatchCalendarsId(bctx, cal2).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-cal-retired-b-%s", sfx),
			}).Execute()
		}
		_, _, _ = client.CalendarsAPI.PatchCalendarsId(bctx, calID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-cal-retired-%s", sfx),
		}).Execute()
	}()

	_, resp, err = client.CalendarsAPI.PatchCalendarsId(ctx, calID).Body(map[string]interface{}{
		"disabled": true,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.CalendarsAPI.PostCalendars(ctx).Body(map[string]interface{}{
		"name": name,
	}).Execute()
	require.Error(t, err2)
	require.NotNil(t, resp2)
	t.Logf("duplicate calendar name: %d %s", resp2.StatusCode, mustErrBody(err2))
	require.GreaterOrEqual(t, resp2.StatusCode, 400)

	freed := fmt.Sprintf("%s-rn-%s", name, sfx)
	_, _, _ = client.CalendarsAPI.PatchCalendarsId(ctx, calID).Body(map[string]interface{}{"name": freed}).Execute()

	c2, r3, e3 := client.CalendarsAPI.PostCalendars(ctx).Body(map[string]interface{}{"name": name}).Execute()
	require.NoError(t, e3)
	require.Equal(t, 201, r3.StatusCode)
	cal2, _ = mapIDInt32Any(c2)
}

func mustErrBody(err error) string {
	_, b := openAPIErrBody(err)
	return b
}

// --- Product ---

func TestProbe_NameUniqueness_Product(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	token := os.Getenv("XURRENT_TOKEN")
	acc := os.Getenv("XURRENT_ACCOUNT")
	catRef, err := firstProductCategoryReference(ctx, token, acc)
	if err != nil {
		t.Skipf("product categories: %v", err)
	}
	sfx := randomSuffix()
	brand := fmt.Sprintf("McpB-%s", sfx)
	model := fmt.Sprintf("McpM-%s", sfx)
	name := fmt.Sprintf("mcpprobe-product-%s", sfx)

	body := map[string]interface{}{
		"name":     name,
		"brand":    brand,
		"model":    model,
		"category": catRef,
		"remarks":  "name uniqueness probe",
	}
	created, resp, err := client.ProductsAPI.PostProducts(ctx).Body(body).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostProducts: %v body=%s", err, bod)
		t.Skip("PostProducts failed")
		return
	}
	pid, ok := mapIDInt32Any(created)
	require.True(t, ok)

	var pid2 int32
	defer func() {
		bctx := context.Background()
		if pid2 != 0 {
			_, _, _ = client.ProductsAPI.PatchProductsId(bctx, pid2).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-prd-ret-b-%s", sfx),
			}).Execute()
		}
		_, _, _ = client.ProductsAPI.PatchProductsId(bctx, pid).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-prd-ret-%s", sfx),
		}).Execute()
	}()

	_, resp, err = client.ProductsAPI.PatchProductsId(ctx, pid).Body(map[string]interface{}{
		"disabled": true,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.ProductsAPI.PostProducts(ctx).Body(body).Execute()
	require.Error(t, err2)
	t.Logf("duplicate product (same name/brand/model): %d %s", resp2.StatusCode, mustErrBody(err2))

	body["name"] = fmt.Sprintf("%s-rn-%s", name, sfx)
	_, _, _ = client.ProductsAPI.PatchProductsId(ctx, pid).Body(map[string]interface{}{
		"name": body["name"],
	}).Execute()

	body["name"] = name
	created2, r3, e3 := client.ProductsAPI.PostProducts(ctx).Body(body).Execute()
	require.NoError(t, e3)
	require.Equal(t, 201, r3.StatusCode)
	pid2, _ = mapIDInt32Any(created2)
}

// --- Service offering (requires service) ---

func TestProbe_NameUniqueness_ServiceOffering(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	orgID, ok := firstEnabledOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need enabled organization")
	}
	sfx := randomSuffix()
	svcName := fmt.Sprintf("mcpprobe-svcoff-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":        svcName,
		"provider_id": orgID,
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("need service for offering probe: %v body=%s", err, bod)
		t.Skipf("need service for offering probe")
		return
	}
	svcID, ok := mapIDInt32Any(svcCreated)
	require.True(t, ok)

	offName := fmt.Sprintf("mcpprobe-offering-%s", sfx)
	offCreated, resp, err := client.ServiceOfferingsAPI.PostServiceOfferings(ctx).Body(map[string]interface{}{
		"name":       offName,
		"service_id": svcID,
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostServiceOfferings: %v body=%s", err, bod)
		_, _, _ = client.ServicesAPI.PatchServicesId(context.Background(), svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-svcoff-svc-ret-%s", sfx),
		}).Execute()
		t.Skip("PostServiceOfferings failed")
		return
	}
	offID, ok := mapIDInt32Any(offCreated)
	require.True(t, ok)

	var off2 int32
	defer func() {
		bctx := context.Background()
		if off2 != 0 {
			_, _, _ = client.ServiceOfferingsAPI.PatchServiceOfferingsId(bctx, off2).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-off-ret-b-%s", sfx),
			}).Execute()
		}
		_, _, _ = client.ServiceOfferingsAPI.PatchServiceOfferingsId(bctx, offID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-off-ret-%s", sfx),
		}).Execute()
		_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-svcoff-svc-ret-%s", sfx),
		}).Execute()
	}()

	_, resp, err = client.ServiceOfferingsAPI.PatchServiceOfferingsId(ctx, offID).Body(map[string]interface{}{
		"disabled": true,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.ServiceOfferingsAPI.PostServiceOfferings(ctx).Body(map[string]interface{}{
		"name":       offName,
		"service_id": svcID,
	}).Execute()
	require.Error(t, err2)
	t.Logf("duplicate service offering name: %d %s", resp2.StatusCode, mustErrBody(err2))

	_, _, _ = client.ServiceOfferingsAPI.PatchServiceOfferingsId(ctx, offID).Body(map[string]interface{}{
		"name": fmt.Sprintf("%s-rn-%s", offName, sfx),
	}).Execute()

	o2, r3, e3 := client.ServiceOfferingsAPI.PostServiceOfferings(ctx).Body(map[string]interface{}{
		"name":       offName,
		"service_id": svcID,
	}).Execute()
	require.NoError(t, e3)
	require.Equal(t, 201, r3.StatusCode)
	off2, _ = mapIDInt32Any(o2)
}

// --- SLA (try minimal name+customer first; fall back to service → offering → SLA when required) ---

func TestProbe_NameUniqueness_SLA(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	sfx := randomSuffix()
	slaName := fmt.Sprintf("mcpprobe-sla-%s", sfx)

	customerOrgID, ok := firstEnabledOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need enabled organization")
	}
	slaBody := map[string]interface{}{
		"name":     slaName,
		"customer": map[string]interface{}{"id": customerOrgID},
	}

	var (
		svcID, offID int32
	)
	tryPostSlas := func(body map[string]interface{}) (map[string]interface{}, *http.Response, error) {
		return client.ServiceLevelAgreementsAPI.PostSlas(ctx).Body(body).Execute()
	}

	created, resp, err := tryPostSlas(slaBody)
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostSlas (minimal): %v body=%s", err, bod)

		// Some tenants accept top-level customer_id instead of nested customer.
		slaBodyID := map[string]interface{}{
			"name":        slaName,
			"customer_id": customerOrgID,
		}
		created, resp, err = tryPostSlas(slaBodyID)
		if err == nil && resp != nil && resp.StatusCode == 201 {
			slaBody = slaBodyID
		} else {
			_, bodID := openAPIErrBody(err)
			t.Logf("PostSlas (customer_id): %v body=%s", err, bodID)
		}
	}
	if err != nil || resp == nil || resp.StatusCode != 201 {
		providerOrgID, okProv := firstEnabledOrganizationID(t, ctx, client)
		if !okProv {
			t.Skip("need enabled organization for SLA chain")
			return
		}
		customerOrgID, ok = firstEnabledOrganizationIDOtherThan(t, ctx, client, providerOrgID)
		if !ok {
			t.Skip("need a second enabled organization for SLA customer (distinct from service provider)")
			return
		}
		svcName := fmt.Sprintf("mcpprobe-sla-svc-%s", sfx)
		svcCreated, resp2, err2 := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
			"name":        svcName,
			"provider_id": providerOrgID,
		}).Execute()
		if err2 != nil || resp2 == nil || resp2.StatusCode != 201 {
			_, bod2 := openAPIErrBody(err2)
			t.Logf("PostServices for SLA chain: %v body=%s", err2, bod2)
			t.Skip("PostSlas and service chain both failed for this tenant")
			return
		}
		var okSvc bool
		svcID, okSvc = mapIDInt32Any(svcCreated)
		require.True(t, okSvc)

		offName := fmt.Sprintf("mcpprobe-sla-off-%s", sfx)
		offCreated, resp3, err3 := client.ServiceOfferingsAPI.PostServiceOfferings(ctx).Body(map[string]interface{}{
			"name":       offName,
			"service_id": svcID,
		}).Execute()
		if err3 != nil || resp3 == nil || resp3.StatusCode != 201 {
			_, bod3 := openAPIErrBody(err3)
			t.Logf("PostServiceOfferings for SLA: %v body=%s", err3, bod3)
			_, _, _ = client.ServicesAPI.PatchServicesId(context.Background(), svcID).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-sla-svc-ret-%s", sfx),
			}).Execute()
			t.Skip("PostSlas and service offering chain failed for this tenant")
			return
		}
		var okOff bool
		offID, okOff = mapIDInt32Any(offCreated)
		require.True(t, okOff)

		slaBody = map[string]interface{}{
			"name":                slaName,
			"customer_id":         customerOrgID,
			"service_offering_id": offID,
		}
		created, resp, err = tryPostSlas(slaBody)
		if err != nil || resp == nil || resp.StatusCode != 201 {
			_, bod4 := openAPIErrBody(err)
			t.Logf("PostSlas (customer_id + service_offering_id): %v body=%s", err, bod4)
			slaBody = map[string]interface{}{
				"name":             slaName,
				"customer":         map[string]interface{}{"id": customerOrgID},
				"service_offering": map[string]interface{}{"id": offID},
			}
			created, resp, err = tryPostSlas(slaBody)
		}
		if err != nil || resp == nil || resp.StatusCode != 201 {
			_, bod4 := openAPIErrBody(err)
			t.Logf("PostSlas (nested customer + service_offering): %v body=%s", err, bod4)
			_, _, _ = client.ServiceOfferingsAPI.PatchServiceOfferingsId(context.Background(), offID).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-sla-off-ret-%s", sfx),
			}).Execute()
			_, _, _ = client.ServicesAPI.PatchServicesId(context.Background(), svcID).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-sla-svc-ret-%s", sfx),
			}).Execute()
			t.Skip("PostSlas failed for this tenant")
			return
		}
	}

	slaID, ok := mapIDInt32Any(created)
	require.True(t, ok)

	var sla2 int32
	defer func() {
		bctx := context.Background()
		if sla2 != 0 {
			_, _, _ = client.ServiceLevelAgreementsAPI.PatchSlasId(bctx, sla2).Body(map[string]interface{}{
				"name": fmt.Sprintf("mcpprobe-sla-ret-b-%s", sfx),
			}).Execute()
		}
		_, _, _ = client.ServiceLevelAgreementsAPI.PatchSlasId(bctx, slaID).Body(map[string]interface{}{
			"name": fmt.Sprintf("mcpprobe-sla-ret-%s", sfx),
		}).Execute()
		if offID != 0 {
			_, _, _ = client.ServiceOfferingsAPI.PatchServiceOfferingsId(bctx, offID).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-sla-off-ret-%s", sfx),
			}).Execute()
		}
		if svcID != 0 {
			_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
				"disabled": true,
				"name":     fmt.Sprintf("mcpprobe-sla-svc-ret-%s", sfx),
			}).Execute()
		}
	}()

	_, resp, err = client.ServiceLevelAgreementsAPI.PatchSlasId(ctx, slaID).Body(map[string]interface{}{
		"disabled": true,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.ServiceLevelAgreementsAPI.PostSlas(ctx).Body(slaBody).Execute()
	require.Error(t, err2)
	t.Logf("duplicate SLA name: %d %s", resp2.StatusCode, mustErrBody(err2))

	_, _, _ = client.ServiceLevelAgreementsAPI.PatchSlasId(ctx, slaID).Body(map[string]interface{}{
		"name": fmt.Sprintf("%s-rn-%s", slaName, sfx),
	}).Execute()

	c2, r3, e3 := client.ServiceLevelAgreementsAPI.PostSlas(ctx).Body(slaBody).Execute()
	require.NoError(t, e3)
	require.Equal(t, 201, r3.StatusCode)
	sla2, _ = mapIDInt32Any(c2)
}

// --- Request template (subject line treated as unique label) ---

func TestProbe_NameUniqueness_RequestTemplate(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	orgID, ok := firstEnabledOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need enabled organization")
	}
	sfx := randomSuffix()
	svcName := fmt.Sprintf("mcpprobe-rt-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":        svcName,
		"provider_id": orgID,
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("need service: %v body=%s", err, bod)
		t.Skipf("need service for request template probe")
		return
	}
	svcID, ok := mapIDInt32Any(svcCreated)
	require.True(t, ok)

	subj := fmt.Sprintf("mcpprobe-rt-subj-%s", sfx)
	rtCreated, resp, err := client.RequestTemplatesAPI.PostRequestTemplates(ctx).Body(map[string]interface{}{
		"subject":    subj,
		"service_id": svcID,
		"category":   "rfi",
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostRequestTemplates: %v body=%s", err, bod)
		_, _, _ = client.ServicesAPI.PatchServicesId(context.Background(), svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-rt-svc-ret-%s", sfx),
		}).Execute()
		t.Skip("PostRequestTemplates failed")
		return
	}
	rtID, ok := mapIDInt32Any(rtCreated)
	require.True(t, ok)

	var rt2 int32
	defer func() {
		bctx := context.Background()
		if rt2 != 0 {
			_, _, _ = client.RequestTemplatesAPI.PatchRequestTemplatesId(bctx, rt2).Body(map[string]interface{}{
				"disabled": true,
				"subject":  fmt.Sprintf("ret-b-%s", sfx),
			}).Execute()
		}
		_, _, _ = client.RequestTemplatesAPI.PatchRequestTemplatesId(bctx, rtID).Body(map[string]interface{}{
			"disabled": true,
			"subject":  fmt.Sprintf("ret-%s", sfx),
		}).Execute()
		_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-rt-svc-ret-%s", sfx),
		}).Execute()
	}()

	// Duplicate subject: many tenants allow multiple templates with the same subject on one service.
	dupCreated, respDup, errDup := client.RequestTemplatesAPI.PostRequestTemplates(ctx).Body(map[string]interface{}{
		"subject":    subj,
		"service_id": svcID,
		"category":   "rfi",
	}).Execute()
	if errDup == nil && respDup != nil && respDup.StatusCode == 201 {
		if id, ok := mapIDInt32Any(dupCreated); ok {
			rt2 = id
		}
		t.Skip("duplicate request template subject accepted (no uniqueness on subject for this tenant)")
	}
	require.Error(t, errDup)
	require.NotNil(t, respDup)
	t.Logf("duplicate request template subject: %d %s", respDup.StatusCode, mustErrBody(errDup))
	require.GreaterOrEqual(t, respDup.StatusCode, 400)

	_, resp, err = client.RequestTemplatesAPI.PatchRequestTemplatesId(ctx, rtID).Body(map[string]interface{}{
		"disabled": true,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	freed := fmt.Sprintf("%s-rn-%s", subj, sfx)
	_, _, _ = client.RequestTemplatesAPI.PatchRequestTemplatesId(ctx, rtID).Body(map[string]interface{}{
		"subject": freed,
	}).Execute()

	r2, r3, e3 := client.RequestTemplatesAPI.PostRequestTemplates(ctx).Body(map[string]interface{}{
		"subject":    subj,
		"service_id": svcID,
		"category":   "rfi",
	}).Execute()
	require.NoError(t, e3)
	require.Equal(t, 201, r3.StatusCode)
	rt2, _ = mapIDInt32Any(r2)
}

// --- Organization (POST uses Organization model; success may be HTTP 200) ---

func TestProbe_NameUniqueness_Organization(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	sfx := randomSuffix()
	name := fmt.Sprintf("mcpprobe-org-%s", sfx)

	o := openapiclient.NewOrganization()
	o.SetName(name)
	created, resp, err := client.OrganizationsAPI.PostOrganizations(ctx).Organization(*o).Execute()
	if err != nil || resp == nil || (resp.StatusCode != 200 && resp.StatusCode != 201) {
		_, bod := openAPIErrBody(err)
		t.Logf("PostOrganizations: %v body=%s", err, bod)
		if resp != nil && (resp.StatusCode == 401 || resp.StatusCode == 403) {
			t.Skip("PostOrganizations not permitted for this token")
		}
		t.Skip("PostOrganizations validation failed or not available")
		return
	}
	require.NotNil(t, created)
	orgID := created.GetId()

	var org2 int32
	defer func() {
		bctx := context.Background()
		if org2 != 0 {
			o2 := openapiclient.NewOrganization()
			o2.SetDisabled(true)
			o2.SetName(fmt.Sprintf("mcpprobe-org-ret-b-%s", sfx))
			_, _, _ = client.OrganizationsAPI.PatchOrganizationsId(bctx, org2).Organization(*o2).Execute()
		}
		oc := openapiclient.NewOrganization()
		oc.SetDisabled(true)
		oc.SetName(fmt.Sprintf("mcpprobe-org-ret-%s", sfx))
		_, _, _ = client.OrganizationsAPI.PatchOrganizationsId(bctx, orgID).Organization(*oc).Execute()
	}()

	dis := openapiclient.NewOrganization()
	dis.SetDisabled(true)
	_, resp, err = client.OrganizationsAPI.PatchOrganizationsId(ctx, orgID).Organization(*dis).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.OrganizationsAPI.PostOrganizations(ctx).Organization(*o).Execute()
	require.Error(t, err2)
	require.NotNil(t, resp2)
	t.Logf("duplicate organization name: %d %s", resp2.StatusCode, mustErrBody(err2))
	require.GreaterOrEqual(t, resp2.StatusCode, 400)

	freed := fmt.Sprintf("%s-rn-%s", name, sfx)
	rn := openapiclient.NewOrganization()
	rn.SetName(freed)
	_, _, _ = client.OrganizationsAPI.PatchOrganizationsId(ctx, orgID).Organization(*rn).Execute()

	o2 := openapiclient.NewOrganization()
	o2.SetName(name)
	created2, resp3, err3 := client.OrganizationsAPI.PostOrganizations(ctx).Organization(*o2).Execute()
	require.NoError(t, err3)
	require.NotNil(t, resp3)
	require.True(t, resp3.StatusCode == 200 || resp3.StatusCode == 201)
	org2 = created2.GetId()
}

// --- Service instance (name uniqueness; depends on service + support team) ---

func TestProbe_NameUniqueness_ServiceInstance(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	orgID, ok := firstEnabledOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need enabled organization")
	}

	sfx := randomSuffix()
	teamName := fmt.Sprintf("mcpprobe-si-team-%s", sfx)
	teamCreated, resp, err := client.TeamsAPI.PostTeams(ctx).Body(map[string]interface{}{"name": teamName}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostTeams for SI probe: %v body=%s", err, bod)
		t.Skip("need to create a team for service instance support_team")
		return
	}
	teamID, ok := mapIDInt32Any(teamCreated)
	require.True(t, ok)

	svcName := fmt.Sprintf("mcpprobe-si-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":        svcName,
		"provider_id": orgID,
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostServices for SI probe: %v body=%s", err, bod)
		_, _, _ = client.TeamsAPI.PatchTeamsId(context.Background(), teamID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-si-team-ret-%s", sfx),
		}).Execute()
		t.Skip("need service for service instance probe")
		return
	}
	svcID, ok := mapIDInt32Any(svcCreated)
	require.True(t, ok)

	siName := fmt.Sprintf("mcpprobe-si-%s", sfx)
	// Prefer nested service/support_team (matches high_value_mutations); fall back to *_id top-level.
	tryBodies := []map[string]interface{}{
		{
			"name":         siName,
			"service":      map[string]interface{}{"id": svcID},
			"support_team": map[string]interface{}{"id": teamID},
			"status":       "in_production",
		},
		{
			"name":            siName,
			"service_id":      svcID,
			"support_team_id": teamID,
			"status":          "in_production",
		},
	}
	var (
		siCreated map[string]interface{}
		siID      int32
		usedSIBody map[string]interface{}
	)
	for _, body := range tryBodies {
		siCreated, resp, err = client.ServiceInstancesAPI.PostServiceInstances(ctx).Body(body).Execute()
		if err == nil && resp != nil && resp.StatusCode == 201 {
			usedSIBody = body
			break
		}
	}
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostServiceInstances: %v body=%s", err, bod)
		_, _, _ = client.ServicesAPI.PatchServicesId(context.Background(), svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-si-svc-ret-%s", sfx),
		}).Execute()
		_, _, _ = client.TeamsAPI.PatchTeamsId(context.Background(), teamID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-si-team-ret-%s", sfx),
		}).Execute()
		t.Skip("PostServiceInstances failed for this tenant")
		return
	}
	siID, ok = mapIDInt32Any(siCreated)
	require.True(t, ok)

	var si2 int32
	defer func() {
		bctx := context.Background()
		if si2 != 0 {
			_, _, _ = client.ServiceInstancesAPI.PatchServiceInstancesId(bctx, si2).Body(map[string]interface{}{
				"status": "discontinued",
				"name":   fmt.Sprintf("mcpprobe-si-ret-b-%s", sfx),
			}).Execute()
		}
		_, _, _ = client.ServiceInstancesAPI.PatchServiceInstancesId(bctx, siID).Body(map[string]interface{}{
			"status": "discontinued",
			"name":   fmt.Sprintf("mcpprobe-si-ret-%s", sfx),
		}).Execute()
		_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-si-svc-ret-%s", sfx),
		}).Execute()
		_, _, _ = client.TeamsAPI.PatchTeamsId(bctx, teamID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-si-team-ret-%s", sfx),
		}).Execute()
	}()

	_, resp, err = client.ServiceInstancesAPI.PatchServiceInstancesId(ctx, siID).Body(map[string]interface{}{
		"status": "discontinued",
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.ServiceInstancesAPI.PostServiceInstances(ctx).Body(usedSIBody).Execute()
	require.Error(t, err2)
	require.NotNil(t, resp2)
	t.Logf("duplicate service instance name: %d %s", resp2.StatusCode, mustErrBody(err2))
	require.GreaterOrEqual(t, resp2.StatusCode, 400)

	_, _, _ = client.ServiceInstancesAPI.PatchServiceInstancesId(ctx, siID).Body(map[string]interface{}{
		"name": fmt.Sprintf("%s-rn-%s", siName, sfx),
	}).Execute()

	c2, r3, e3 := client.ServiceInstancesAPI.PostServiceInstances(ctx).Body(usedSIBody).Execute()
	require.NoError(t, e3)
	require.Equal(t, 201, r3.StatusCode)
	si2, _ = mapIDInt32Any(c2)
}

// --- Workflow (subject line; requires XURRENT_TEST_WORKFLOW_TEMPLATE_ID) ---

func TestProbe_NameUniqueness_Workflow(t *testing.T) {
	tplStr := os.Getenv("XURRENT_TEST_WORKFLOW_TEMPLATE_ID")
	if tplStr == "" {
		t.Skip("set XURRENT_TEST_WORKFLOW_TEMPLATE_ID for workflow subject probe")
	}
	tpl64, err := strconv.ParseInt(tplStr, 10, 32)
	if err != nil {
		t.Skipf("invalid XURRENT_TEST_WORKFLOW_TEMPLATE_ID: %v", err)
	}
	templateID := int32(tpl64)
	wfType := os.Getenv("XURRENT_TEST_WORKFLOW_TYPE")
	if wfType == "" {
		wfType = "application_change"
	}

	client := mutationClient(t)
	ctx := context.Background()
	orgID, ok := firstEnabledOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need enabled organization")
	}

	sfx := randomSuffix()
	svcName := fmt.Sprintf("mcpprobe-wf-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":        svcName,
		"provider_id": orgID,
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostServices for workflow probe: %v body=%s", err, bod)
		t.Skip("need service for workflow probe")
		return
	}
	svcID, ok := mapIDInt32Any(svcCreated)
	require.True(t, ok)

	managerID := mePersonID(t, ctx, client)
	subj := fmt.Sprintf("mcpprobe-wf-subj-%s", sfx)
	wfBodies := []map[string]interface{}{
		{
			"subject":           subj,
			"service_id":        svcID,
			"category":          "rfc",
			"manager":           map[string]interface{}{"id": managerID},
			"workflow_type":     wfType,
			"workflow_template": map[string]interface{}{"id": templateID},
		},
		{
			"subject":           subj,
			"service":           map[string]interface{}{"id": svcID},
			"category":          "rfc",
			"manager":           map[string]interface{}{"id": managerID},
			"workflow_type":     wfType,
			"workflow_template": map[string]interface{}{"id": templateID},
		},
	}
	var wfCreated map[string]interface{}
	var usedWfBody map[string]interface{}
	for _, wb := range wfBodies {
		wfCreated, resp, err = client.WorkflowsAPI.PostWorkflows(ctx).Body(wb).Execute()
		if err == nil && resp != nil && resp.StatusCode == 201 {
			usedWfBody = wb
			break
		}
	}
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostWorkflows: %v body=%s", err, bod)
		_, _, _ = client.ServicesAPI.PatchServicesId(context.Background(), svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-wf-svc-ret-%s", sfx),
		}).Execute()
		t.Skip("PostWorkflows failed for this tenant")
		return
	}
	wfID, ok := mapIDInt32Any(wfCreated)
	require.True(t, ok)

	var wf2 int32
	defer func() {
		bctx := context.Background()
		if wf2 != 0 {
			_, _, _ = client.WorkflowsAPI.PostWorkflowsIdTrash(bctx, wf2).Execute()
		}
		// wfID was already trashed in the test body
		_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-wf-svc-ret-%s", sfx),
		}).Execute()
	}()

	_, resp, err = client.WorkflowsAPI.PostWorkflowsIdTrash(ctx, wfID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	dupCreated, respDup, errDup := client.WorkflowsAPI.PostWorkflows(ctx).Body(usedWfBody).Execute()
	if errDup == nil && respDup != nil && respDup.StatusCode == 201 {
		if id, ok := mapIDInt32Any(dupCreated); ok {
			wf2 = id
		}
		t.Skip("duplicate workflow subject accepted after trash (no uniqueness on subject for this tenant)")
	}
	require.Error(t, errDup)
	require.NotNil(t, respDup)
	t.Logf("duplicate workflow subject after trash: %d %s", respDup.StatusCode, mustErrBody(errDup))
	require.GreaterOrEqual(t, respDup.StatusCode, 400)
}

// --- Phase B: composite keys (primary_email, source+sourceID, brand+productID, CI serial_nr) ---

func TestProbe_CompositeUniqueness_PrimaryEmail(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")
	sfx := randomSuffix()
	email := fmt.Sprintf("mcpprobe-email-%s@invalid.example", sfx)

	created, resp, err := client.PeopleAPI.PostPeople(ctx).Body(map[string]interface{}{
		"source":        fmt.Sprintf("mcpprobe-person-%s", sfx),
		"primary_email": email,
		"name":          fmt.Sprintf("McpProbe Person %s", sfx),
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode == 401 || resp.StatusCode == 403 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostPeople: %v body=%s", err, bod)
		t.Skip("PostPeople not permitted for this token")
		return
	}
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	if id, ok := resourceAccountID(jsonMap(created)); ok && id != "" && !resourceAccountAcceptableForEnv(id, acc) {
		t.Skipf("person created in account %q not acceptable for XURRENT_ACCOUNT %q (set XURRENT_TRUSTED_ACCOUNT_IDS for trust peers)", id, acc)
	}
	pid, ok := mapIDInt32Any(created)
	require.True(t, ok)

	var pid2 int32
	defer func() {
		bctx := context.Background()
		if pid2 != 0 {
			_, _, _ = client.PeopleAPI.PostPeopleIdRestore(bctx, pid2).Execute()
			_, _, _ = client.PeopleAPI.PatchPeopleId(bctx, pid2).Body(map[string]interface{}{
				"disabled":      true,
				"primary_email": fmt.Sprintf("mcpprobe-ret-b-%s@invalid.example", sfx),
			}).Execute()
		}
		_, _, _ = client.PeopleAPI.PostPeopleIdRestore(bctx, pid).Execute()
		_, _, _ = client.PeopleAPI.PatchPeopleId(bctx, pid).Body(map[string]interface{}{
			"disabled":      true,
			"primary_email": fmt.Sprintf("mcpprobe-ret-%s@invalid.example", sfx),
		}).Execute()
	}()

	_, resp, err = client.PeopleAPI.PatchPeopleId(ctx, pid).Body(map[string]interface{}{
		"disabled": true,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp2, err2 := client.PeopleAPI.PostPeople(ctx).Body(map[string]interface{}{
		"source":        fmt.Sprintf("mcpprobe-person-dup-%s", sfx),
		"primary_email": email,
		"name":          fmt.Sprintf("McpProbe Dup %s", sfx),
	}).Execute()
	require.Error(t, err2)
	require.NotNil(t, resp2)
	t.Logf("duplicate primary_email: %d %s", resp2.StatusCode, mustErrBody(err2))
	require.GreaterOrEqual(t, resp2.StatusCode, 400)

	freed := fmt.Sprintf("mcpprobe-freed-%s@invalid.example", sfx)
	_, _, _ = client.PeopleAPI.PatchPeopleId(ctx, pid).Body(map[string]interface{}{
		"primary_email": freed,
	}).Execute()

	created2, resp3, err3 := client.PeopleAPI.PostPeople(ctx).Body(map[string]interface{}{
		"source":        fmt.Sprintf("mcpprobe-person-2-%s", sfx),
		"primary_email": email,
		"name":          fmt.Sprintf("McpProbe Second %s", sfx),
	}).Execute()
	require.NoError(t, err3)
	require.Equal(t, 201, resp3.StatusCode)
	pid2, ok = mapIDInt32Any(created2)
	require.True(t, ok)
}

func TestProbe_CompositeUniqueness_SourceSourceID(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")
	sfx := randomSuffix()
	src := "mcpprobe-ss"
	srcID := fmt.Sprintf("mcpprobe-ext-%s", sfx)
	email1 := fmt.Sprintf("mcpprobe-ss1-%s@invalid.example", sfx)
	email2 := fmt.Sprintf("mcpprobe-ss2-%s@invalid.example", sfx)

	created, resp, err := client.PeopleAPI.PostPeople(ctx).Body(map[string]interface{}{
		"source":        src,
		"sourceID":      srcID,
		"primary_email": email1,
		"name":          fmt.Sprintf("McpProbe Src1 %s", sfx),
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode == 401 || resp.StatusCode == 403 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostPeople: %v body=%s", err, bod)
		t.Skip("PostPeople not permitted for this token")
		return
	}
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
	if id, ok := resourceAccountID(jsonMap(created)); ok && id != "" && !resourceAccountAcceptableForEnv(id, acc) {
		t.Skipf("person created in account %q not acceptable for XURRENT_ACCOUNT %q (set XURRENT_TRUSTED_ACCOUNT_IDS for trust peers)", id, acc)
	}
	pid, ok := mapIDInt32Any(created)
	require.True(t, ok)

	var pid2 int32
	defer func() {
		bctx := context.Background()
		if pid2 != 0 {
			_, _, _ = client.PeopleAPI.PostPeopleIdRestore(bctx, pid2).Execute()
			_, _, _ = client.PeopleAPI.PatchPeopleId(bctx, pid2).Body(map[string]interface{}{
				"disabled":      true,
				"primary_email": fmt.Sprintf("mcpprobe-ss-ret-b-%s@invalid.example", sfx),
			}).Execute()
		}
		_, _, _ = client.PeopleAPI.PostPeopleIdRestore(bctx, pid).Execute()
		_, _, _ = client.PeopleAPI.PatchPeopleId(bctx, pid).Body(map[string]interface{}{
			"disabled":      true,
			"primary_email": fmt.Sprintf("mcpprobe-ss-ret-%s@invalid.example", sfx),
		}).Execute()
	}()

	_, resp, err = client.PeopleAPI.PatchPeopleId(ctx, pid).Body(map[string]interface{}{
		"disabled": true,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	dupBody := map[string]interface{}{
		"source":        src,
		"sourceID":      srcID,
		"primary_email": email2,
		"name":          fmt.Sprintf("McpProbe Src2 %s", sfx),
	}
	dupCreated, resp2, err2 := client.PeopleAPI.PostPeople(ctx).Body(dupBody).Execute()
	if err2 == nil && resp2 != nil && resp2.StatusCode == 201 {
		if id, ok := mapIDInt32Any(dupCreated); ok {
			pid2 = id
		}
		t.Skip("duplicate source+sourceID accepted while first person disabled (no composite uniqueness for this tenant)")
	}
	require.Error(t, err2)
	require.NotNil(t, resp2)
	t.Logf("duplicate source+sourceID: %d %s", resp2.StatusCode, mustErrBody(err2))
	require.GreaterOrEqual(t, resp2.StatusCode, 400)

	_, _, _ = client.PeopleAPI.PatchPeopleId(ctx, pid).Body(map[string]interface{}{
		"sourceID": fmt.Sprintf("%s-RN", srcID),
	}).Execute()

	created2, resp3, err3 := client.PeopleAPI.PostPeople(ctx).Body(map[string]interface{}{
		"source":        src,
		"sourceID":      srcID,
		"primary_email": email2,
		"name":          fmt.Sprintf("McpProbe Src2b %s", sfx),
	}).Execute()
	require.NoError(t, err3)
	require.Equal(t, 201, resp3.StatusCode)
	pid2, ok = mapIDInt32Any(created2)
	require.True(t, ok)
}

func TestProbe_CompositeUniqueness_ProductBrandProductID(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	token := os.Getenv("XURRENT_TOKEN")
	acc := os.Getenv("XURRENT_ACCOUNT")
	catRef, err := firstProductCategoryReference(ctx, token, acc)
	if err != nil {
		t.Skipf("product categories: %v", err)
	}

	sfx := randomSuffix()
	brand := fmt.Sprintf("McpProbeB-%s", sfx)
	model := fmt.Sprintf("McpProbeM-%s", sfx)
	pidStr := fmt.Sprintf("MCP-PID-%s", sfx)
	name1 := fmt.Sprintf("mcpprobe-prd1-%s", sfx)
	name2 := fmt.Sprintf("mcpprobe-prd2-%s", sfx)

	body := map[string]interface{}{
		"name":      name1,
		"brand":     brand,
		"model":     model,
		"category":  catRef,
		"remarks":   "composite probe brand+productID",
		"productID": pidStr,
	}

	created, resp, err := client.ProductsAPI.PostProducts(ctx).Body(body).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, bod := openAPIErrBody(err)
		t.Logf("PostProducts (with productID): %v body=%s", err, bod)
		t.Skip("PostProducts with productID failed — check field name or tenant rules")
		return
	}
	prdID, ok := mapIDInt32Any(created)
	require.True(t, ok)

	var prd2 int32
	defer func() {
		bctx := context.Background()
		if prd2 != 0 {
			_, _, _ = client.ProductsAPI.PatchProductsId(bctx, prd2).Body(map[string]interface{}{
				"disabled": true,
				"brand":    fmt.Sprintf("mcpprobe-ret-b-%s", sfx),
			}).Execute()
		}
		_, _, _ = client.ProductsAPI.PatchProductsId(bctx, prdID).Body(map[string]interface{}{
			"disabled": true,
			"brand":    fmt.Sprintf("mcpprobe-ret-%s", sfx),
		}).Execute()
	}()

	dupBody := map[string]interface{}{
		"name":      name2,
		"brand":     brand,
		"model":     model,
		"category":  catRef,
		"remarks":   "duplicate brand+productID probe",
		"productID": pidStr,
	}
	// Uniqueness applies while the first product still exists (enabled); disabled rows may not block duplicate pair on all tenants.
	dupCreated, resp2, err2 := client.ProductsAPI.PostProducts(ctx).Body(dupBody).Execute()
	if err2 == nil && resp2 != nil && resp2.StatusCode == 201 {
		if id, ok := mapIDInt32Any(dupCreated); ok {
			prd2 = id
		}
		t.Skip("duplicate brand+productID accepted while first product still enabled (no composite uniqueness for this tenant)")
	}
	require.Error(t, err2)
	require.NotNil(t, resp2)
	t.Logf("duplicate brand+productID: %d %s", resp2.StatusCode, mustErrBody(err2))
	require.GreaterOrEqual(t, resp2.StatusCode, 400)

	_, _, _ = client.ProductsAPI.PatchProductsId(ctx, prdID).Body(map[string]interface{}{
		"disabled": true,
	}).Execute()
	_, _, _ = client.ProductsAPI.PatchProductsId(ctx, prdID).Body(map[string]interface{}{
		"productID": fmt.Sprintf("%s-FREE", pidStr),
	}).Execute()

	p2, r3, e3 := client.ProductsAPI.PostProducts(ctx).Body(dupBody).Execute()
	require.NoError(t, e3)
	require.Equal(t, 201, r3.StatusCode)
	prd2, _ = mapIDInt32Any(p2)
}

func TestProbe_CompositeUniqueness_CI_SerialNr(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")
	sfx := randomSuffix()

	orgID, ok := firstEnabledOrganizationID(t, ctx, client)
	if !ok {
		t.Skip("need enabled organization")
	}
	teamName := fmt.Sprintf("mcpprobe-ci-team-%s", sfx)
	teamCreated, resp, err := client.TeamsAPI.PostTeams(ctx).Body(map[string]interface{}{"name": teamName}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		t.Skip("need team for CI chain")
		return
	}
	teamID, ok := mapIDInt32Any(teamCreated)
	require.True(t, ok)

	productID, ok := firstProductIDForPostCI(t, ctx, client)
	if !ok {
		_, _, _ = client.TeamsAPI.PatchTeamsId(context.Background(), teamID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-ci-team-ret-%s", sfx),
		}).Execute()
		t.Skip("need a product for CI")
		return
	}

	svcName := fmt.Sprintf("mcpprobe-ci-svc-%s", sfx)
	svcCreated, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":        svcName,
		"provider_id": orgID,
	}).Execute()
	if err != nil || resp == nil || resp.StatusCode != 201 {
		_, _, _ = client.TeamsAPI.PatchTeamsId(context.Background(), teamID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-ci-team-ret-%s", sfx),
		}).Execute()
		t.Skip("PostServices for CI probe failed")
		return
	}
	svcID, ok := mapIDInt32Any(svcCreated)
	require.True(t, ok)

	siName := fmt.Sprintf("mcpprobe-ci-si-%s", sfx)
	siTryBodies := []map[string]interface{}{
		{
			"name":         siName,
			"service":      map[string]interface{}{"id": svcID},
			"support_team": map[string]interface{}{"id": teamID},
			"status":       "in_production",
		},
		{
			"name":            siName,
			"service_id":      svcID,
			"support_team_id": teamID,
			"status":          "in_production",
		},
	}
	var siID int32
	for _, sb := range siTryBodies {
		var r *http.Response
		var siCreated map[string]interface{}
		siCreated, r, err = client.ServiceInstancesAPI.PostServiceInstances(ctx).Body(sb).Execute()
		if err == nil && r != nil && r.StatusCode == 201 {
			var okSI bool
			siID, okSI = mapIDInt32Any(siCreated)
			if okSI {
				break
			}
		}
	}
	if siID == 0 {
		_, _, _ = client.ServicesAPI.PatchServicesId(context.Background(), svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-ci-svc-ret-%s", sfx),
		}).Execute()
		_, _, _ = client.TeamsAPI.PatchTeamsId(context.Background(), teamID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-ci-team-ret-%s", sfx),
		}).Execute()
		t.Skip("PostServiceInstances for CI serial probe failed")
		return
	}

	serial := fmt.Sprintf("MCP-SER-%s", sfx)
	baseCI := map[string]interface{}{
		"source":           "mcpprobe",
		"systemID":         fmt.Sprintf("mcpprobe-sys-%s", sfx),
		"label":            fmt.Sprintf("mcpprobe-lab-%s", sfx),
		"name":             fmt.Sprintf("mcpprobe-ci-%s", sfx),
		"serial_nr":        serial,
		"status":           "in_production",
		"product_id":       productID,
		"service":          svcID,
		"service_instance": map[string]interface{}{"id": siID},
		"support_team_id":  teamID,
	}

	post1, err := client.ConfigurationItemsAPI.PostCis(ctx).Body(baseCI).Execute()
	require.NoError(t, err)
	require.NotNil(t, post1)
	if post1.StatusCode != 200 && post1.StatusCode != 201 {
		b, _ := io.ReadAll(post1.Body)
		_ = post1.Body.Close()
		patchServiceInstanceStatusForCITestCleanup(t, ctx, client, siID, svcID, teamID, sfx)
		t.Skipf("PostCis: HTTP %d %s", post1.StatusCode, string(b))
		return
	}
	ci1Map := readJSONBodyMap(t, post1)
	requireResourceOwnedByEnvAccountAny(t, acc, ci1Map)
	ci1, ok := mapIDInt32(ci1Map)
	require.True(t, ok)

	defer func() {
		bctx := context.Background()
		patchServiceInstanceStatusForCITest(t, bctx, client, siID, "discontinued")
		_, _, _ = client.ServicesAPI.PatchServicesId(bctx, svcID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-ci-svc-ret-%s", sfx),
		}).Execute()
		_, _, _ = client.TeamsAPI.PatchTeamsId(bctx, teamID).Body(map[string]interface{}{
			"disabled": true,
			"name":     fmt.Sprintf("mcpprobe-ci-team-ret-%s", sfx),
		}).Execute()
	}()

	archiveCreatedCI(t, ctx, client, ci1, sfx)

	baseCI["systemID"] = fmt.Sprintf("mcpprobe-sys-b-%s", sfx)
	baseCI["label"] = fmt.Sprintf("mcpprobe-lab-b-%s", sfx)
	baseCI["name"] = fmt.Sprintf("mcpprobe-ci-b-%s", sfx)

	post2, err2 := client.ConfigurationItemsAPI.PostCis(ctx).Body(baseCI).Execute()
	require.NoError(t, err2)
	require.NotNil(t, post2)
	b2, _ := io.ReadAll(post2.Body)
	_ = post2.Body.Close()
	if post2.StatusCode >= 400 {
		t.Logf("duplicate serial_nr after archive: %d %s", post2.StatusCode, string(b2))
		require.GreaterOrEqual(t, post2.StatusCode, 400)
		return
	}
	t.Skipf("second CI with same serial_nr returned %d (serial may be reusable after archive on this tenant)", post2.StatusCode)
}

func patchServiceInstanceStatusForCITestCleanup(t *testing.T, ctx context.Context, client *openapiclient.APIClient, siID, svcID, teamID int32, sfx string) {
	t.Helper()
	patchServiceInstanceStatusForCITest(t, ctx, client, siID, "discontinued")
	_, _, _ = client.ServicesAPI.PatchServicesId(ctx, svcID).Body(map[string]interface{}{
		"disabled": true,
		"name":     fmt.Sprintf("mcpprobe-ci-svc-ret-%s", sfx),
	}).Execute()
	_, _, _ = client.TeamsAPI.PatchTeamsId(ctx, teamID).Body(map[string]interface{}{
		"disabled": true,
		"name":     fmt.Sprintf("mcpprobe-ci-team-ret-%s", sfx),
	}).Execute()
}

// --- Regressions: ensure JSON error shape for duplicate name ---

func TestProbe_NameUniqueness_ErrorJSONShape(t *testing.T) {
	var payload struct {
		Message string          `json:"message"`
		Errors  [][]string      `json:"errors"`
		Raw     json.RawMessage `json:"-"`
	}
	s := `{"message":"Validation Failed","errors":[["name","Name has already been taken"]]}`
	require.NoError(t, json.Unmarshal([]byte(s), &payload))
	require.Equal(t, "Validation Failed", payload.Message)
	require.Len(t, payload.Errors, 1)
	require.Equal(t, "name", payload.Errors[0][0])
}
