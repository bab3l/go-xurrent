//go:build integration

// Production smoke tests (read-only; no POST/PATCH that create or mutate tenant data).
//
//   go test -tags=integration ./test/integration/ -v -count=1
//
// Variables are documented in .env.example (no committed .env — use local file or export env vars).
//
// Unauthenticated:
//   (none required) — exercises GET /v1/rate_limit
//
// Authenticated (optional):
//   XURRENT_TOKEN       OAuth bearer token
//   XURRENT_ACCOUNT     Account id for X-4me-Account header
// Optional:
//   XURRENT_TEST_PROBLEM_ID  If set, runs GET /v1/problems/{id}
//
// Mutations (POST workflow task, POST product) are intentionally not run here to avoid
// creating records that would require cleanup on a live tenant.

package integration_test

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/xurrent/go-xurrent"
)

func prodClient(t *testing.T, withAuth bool) *openapiclient.APIClient {
	t.Helper()
	cfg := openapiclient.NewConfiguration()
	if withAuth {
		token := os.Getenv("XURRENT_TOKEN")
		account := os.Getenv("XURRENT_ACCOUNT")
		if token == "" || account == "" {
			t.Skip("set XURRENT_TOKEN and XURRENT_ACCOUNT for authenticated checks")
		}
		cfg.AddDefaultHeader("Authorization", "Bearer "+token)
		cfg.AddDefaultHeader("X-4me-Account", account)
	}
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
	_, ok := body["resources"]
	require.True(t, ok, "expected JSON body with resources (see https://developer.xurrent.com/v1/rate_limit)")
}

func TestProduction_Authenticated_ReadOnlyEndpoints(t *testing.T) {
	client := prodClient(t, true)
	ctx := context.Background()
	token := os.Getenv("XURRENT_TOKEN")
	account := os.Getenv("XURRENT_ACCOUNT")

	auth := "Bearer " + token

	_, resp, err := client.TasksAPI.GetTasksFinished(ctx).Authorization(auth).X4meAccount(account).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.TasksAPI.GetTasksApprovalByMe(ctx).Authorization(auth).X4meAccount(account).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProductsAPI.GetProductsDisabled(ctx).Authorization(auth).X4meAccount(account).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProductsAPI.GetProductsEnabled(ctx).Authorization(auth).X4meAccount(account).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.ProductsAPI.GetProductsSupportedByMyTeams(ctx).Authorization(auth).X4meAccount(account).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.AccountAPI.GetAccountBillableUsers(ctx).Authorization(auth).X4meAccount(account).Execute()
	if err != nil {
		require.NotNil(t, resp)
		require.Equal(t, 403, resp.StatusCode, "billable users may require elevated account role")
	} else {
		require.Equal(t, 200, resp.StatusCode)
	}

	_, resp, err = client.AccountAPI.GetAccountUsageStatements(ctx).Authorization(auth).X4meAccount(account).Execute()
	if err != nil {
		require.NotNil(t, resp)
		require.Equal(t, 403, resp.StatusCode, "usage statements require account owner per docs")
	} else {
		require.Equal(t, 200, resp.StatusCode)
	}
}

func TestProduction_GetProblemByID(t *testing.T) {
	pid := os.Getenv("XURRENT_TEST_PROBLEM_ID")
	if pid == "" {
		t.Skip("set XURRENT_TEST_PROBLEM_ID to verify GET /v1/problems/{id}")
	}
	id64, err := strconv.ParseInt(pid, 10, 32)
	require.NoError(t, err)

	client := prodClient(t, true)
	ctx := context.Background()
	account := os.Getenv("XURRENT_ACCOUNT")

	_, resp, err := client.ProblemsAPI.GetProblemsId(ctx, int32(id64)).X4meAccount(account).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}
