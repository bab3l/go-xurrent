//go:build integration

// Production smoke tests (read-only; no POST/PATCH that create or mutate tenant data).
//
//   go test -tags=integration ./test/integration/ -v -count=1
//
// Variables are documented in .env.example (no committed .env — use local file or export env vars).
// Optional: XURRENT_STRUCTURE_LOG_FILE — NDJSON path for redacted JSON structure logs (see utils/merge_skeleton_into_openapi.py).
//
// Unauthenticated:
//   (none required) — exercises GET /v1/rate_limit
//
// Authenticated (optional):
//   XURRENT_TOKEN       OAuth bearer token
//   XURRENT_ACCOUNT     Account id for X-4me-Account header
// Optional:
//   XURRENT_TEST_PROBLEM_ID  If set (or in parent_fixture.env), runs GET /v1/problems/{id}; otherwise discovers an id from problem list endpoints when possible.
//
// Mutations (POST workflow task, POST product) are intentionally not run here to avoid
// creating records that would require cleanup on a live tenant.

package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/xurrent/go-xurrent"
)

func prodClient(t *testing.T, withAuth bool) *openapiclient.APIClient {
	t.Helper()
	if withAuth {
		if os.Getenv("XURRENT_TOKEN") == "" || os.Getenv("XURRENT_ACCOUNT") == "" {
			t.Skip("set XURRENT_TOKEN and XURRENT_ACCOUNT for authenticated checks")
		}
	}
	cfg := newIntegrationConfig(t, withAuth)
	return openapiclient.NewAPIClient(cfg)
}

func TestProduction_GetRateLimit_Unauthenticated(t *testing.T) {
	client := prodClient(t, false)
	ctx := context.Background()
	body, resp, err := client.GeneralAPI.GetRateLimit(ctx).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.NotNil(t, body)
	_, ok := jsonMap(body)["resources"]
	require.True(t, ok, "expected JSON body with resources (see https://developer.xurrent.com/v1/rate_limit)")
}

func TestProduction_Authenticated_ReadOnlyEndpoints(t *testing.T) {
	client := prodClient(t, true)
	ctx := context.Background()
	// Authorization and X-4me-Account come from Configuration.DefaultHeader (see prodClient);
	// do not also set per-request Authorization — duplicate headers can yield 400 from the API.

	_, resp, err := client.TasksAPI.GetTasksFinished(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.TasksAPI.GetTasksApprovalByMe(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProductsAPI.GetProductsDisabled(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProductsAPI.GetProductsEnabled(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProductsAPI.GetProductsSupportedByMyTeams(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.AccountAPI.GetAccountBillableUsers(ctx).Execute()
	if err != nil {
		require.NotNil(t, resp)
		require.Equal(t, 403, resp.StatusCode, "billable users may require elevated account role")
	} else {
		require.Equal(t, 200, resp.StatusCode)
	}

	_, resp, err = client.AccountAPI.GetAccountUsageStatements(ctx).Execute()
	if err != nil {
		require.NotNil(t, resp)
		require.Equal(t, 403, resp.StatusCode, "usage statements require account owner per docs")
	} else {
		require.Equal(t, 200, resp.StatusCode)
	}
}

func TestProduction_GetProblemByID(t *testing.T) {
	client := prodClient(t, true)
	ctx := context.Background()
	account := os.Getenv("XURRENT_ACCOUNT")

	pid, ok := problemIDFromEnvOrAPI(t, ctx, client)
	if !ok {
		t.Skip("no problem id (set XURRENT_TEST_PROBLEM_ID, parent_fixture.env, or ensure a problem list returns rows)")
	}

	_, resp, err := client.ProblemsAPI.GetProblemsId(ctx, pid).X4meAccount(account).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}
