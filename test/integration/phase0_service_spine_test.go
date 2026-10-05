//go:build integration

// Phase 0 (live capture spine): verify the integration principal can POST /v1/services.
// See docs/internal/LIVE_CAPTURE_GAP_CLOSURE_PLAN.md and .env.example (XURRENT_PHASE0_POST_SERVICES).
//
//	go test -tags=integration ./test/integration/ -run TestPhase0_PostServicesSpine -v -count=1
//
// Requires XURRENT_TOKEN, XURRENT_ACCOUNT, and XURRENT_PHASE0_POST_SERVICES=1.

package integration_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/xurrent/go-xurrent"
)

const phase0PostServicesEnv = "XURRENT_PHASE0_POST_SERVICES"

func phase0Client(t *testing.T) *openapiclient.APIClient {
	t.Helper()
	cfg := newIntegrationConfig(t, true)
	return openapiclient.NewAPIClient(cfg)
}

// TestPhase0_PostServicesSpine fails with troubleshooting text if POST /v1/services returns 401/403.
// Success (201) disables the service in a defer. Run this after granting configuration roles to the API user.
func TestPhase0_PostServicesSpine(t *testing.T) {
	if os.Getenv(phase0PostServicesEnv) != "1" {
		t.Skipf("set %s=1 to run service spine check (see docs/internal/LIVE_CAPTURE_GAP_CLOSURE_PLAN.md)", phase0PostServicesEnv)
	}

	client := phase0Client(t)
	ctx := context.Background()
	acc := os.Getenv("XURRENT_ACCOUNT")

	orgID, ok := firstOrganizationID(t, ctx, client)
	if !ok {
		t.Fatal("need at least one organization (GET /v1/organizations) to use as service provider")
	}

	sfx := randomSuffix()
	name := fmt.Sprintf("sdk-phase0-svc-%s", sfx)

	created, resp, err := client.ServicesAPI.PostServices(ctx).Body(map[string]interface{}{
		"name":     name,
		"provider": map[string]interface{}{"id": orgID},
	}).Execute()

	if resp != nil && (resp.StatusCode == 401 || resp.StatusCode == 403) {
		t.Fatalf(
			"POST /v1/services returned HTTP %d — the bearer identity cannot create services in account %q.\n\n"+
				"Fix (pick one or combine):\n"+
				"  • In Xurrent UI: grant this person a role that may maintain services, e.g. **configuration_manager** or **account_administrator** on that account.\n"+
				"  • Via API: POST /v1/people/{id}/permissions/{accountId}?roles=configuration_manager (requires an admin token). See:\n"+
				"    https://developer.xurrent.com/v1/people/permissions/\n"+
				"  • Use an OAuth token for a **person** with those roles (not a limited app-only token if your tenant restricts it).\n"+
				"  • Confirm XURRENT_ACCOUNT matches the account where those roles were granted.\n\n"+
				"After fixing, re-run with %s=1.",
			resp.StatusCode, acc, phase0PostServicesEnv,
		)
	}

	require.NoError(t, err, "PostServices")
	require.NotNil(t, resp)
	require.Equal(t, 201, resp.StatusCode, "PostServices should return 201 Created")
	requireResourceOwnedByEnvAccount(t, acc, created)

	svcID, ok := mapIDInt32(created)
	require.True(t, ok, "service id in response")
	defer func() {
		_, _, _ = client.ServicesAPI.PatchServicesId(context.Background(), svcID).Body(map[string]interface{}{
			"disabled": true,
			"remarks":  fmt.Sprintf("disabled after Phase 0 spine check %s", sfx),
		}).Execute()
	}()

	_, getResp, err := client.ServicesAPI.GetServicesId(ctx, svcID).Execute()
	require.NoError(t, err)
	require.NotNil(t, getResp)
	require.Equal(t, 200, getResp.StatusCode)
	t.Logf("Phase 0 OK: created and fetched service id=%d", svcID)
}
