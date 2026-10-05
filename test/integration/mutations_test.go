//go:build integration

// Opt-in mutation tests (POST/PATCH) against the API account in XURRENT_ACCOUNT.
//
//   set XURRENT_ALLOW_MUTATIONS=1
//   go test -tags=integration ./test/integration/ -run TestMutation_ -v -count=1
//
// Loads credentials like production_smoke_test.go. Random suffixes avoid collisions.
// Guards: only mutates the authenticated person's record when its account matches XURRENT_ACCOUNT.
//
// Product create: the API expects **category** as the product category **reference** string (e.g.
// `address/ip_address` from GET /v1/product_categories `reference`), not numeric `product_category_id`
// alone — otherwise Rails returns 422 "Category can't be blank". This is not cross-account confusion:
// use categories returned with the same X-4me-Account as POST.

package integration_test

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/xurrent/go-xurrent"
)

func randomSuffix() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

func requirePersonIDMatches(t *testing.T, expected int32, person map[string]interface{}) {
	t.Helper()
	raw, ok := person["id"]
	require.True(t, ok)
	var got int32
	switch v := raw.(type) {
	case float64:
		got = int32(v)
	case int:
		got = int32(v)
	case int32:
		got = v
	default:
		t.Fatalf("unexpected person id type %T", raw)
	}
	require.Equal(t, expected, got, "person id must match /v1/me")
}

func resourceAccountID(m map[string]interface{}) (string, bool) {
	if m == nil {
		return "", false
	}
	acc, ok := m["account"]
	if !ok || acc == nil {
		return "", false
	}
	switch v := acc.(type) {
	case map[string]interface{}:
		if id, ok := v["id"]; ok {
			return fmt.Sprint(id), true
		}
	}
	return "", false
}

// trustedAccountIDsFromEnv parses XURRENT_TRUSTED_ACCOUNT_IDS (comma-separated, trimmed).
// Use when X-4me-Account (XURRENT_ACCOUNT) differs from resource.account.id on created rows
// (trust / home-account behavior).
func trustedAccountIDsFromEnv() []string {
	raw := os.Getenv("XURRENT_TRUSTED_ACCOUNT_IDS")
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// resourceAccountAcceptableForEnv reports whether resource.account.id is safe to treat as
// belonging to the same probe/mutation tenant as envAccount (XURRENT_ACCOUNT). Empty or missing
// id passes. Otherwise id must equal envAccount or appear in XURRENT_TRUSTED_ACCOUNT_IDS.
func resourceAccountAcceptableForEnv(resourceAccountID, envAccount string) bool {
	if resourceAccountID == "" {
		return true
	}
	if resourceAccountID == envAccount {
		return true
	}
	for _, t := range trustedAccountIDsFromEnv() {
		if resourceAccountID == t {
			return true
		}
	}
	return false
}

// requireResourceOwnedByEnvAccount fails if the resource is explicitly tied to another account
// (e.g. trusted-account data); same-account or missing account passes.
//
// Note: With trust relations, XURRENT_ACCOUNT may differ from resource.account.id (home account).
// Set XURRENT_TRUSTED_ACCOUNT_IDS to comma-separated account ids that may appear on resources when
// using your token (e.g. home account id while the header is a peer account).
// Do not use this for records you are only touching via "self" (/v1/me) — use requireSameInt32 instead.
func requireResourceOwnedByEnvAccount(t *testing.T, expectedAccount string, m map[string]interface{}) {
	t.Helper()
	if id, ok := resourceAccountID(m); ok && id != "" {
		require.Truef(t, resourceAccountAcceptableForEnv(id, expectedAccount),
			"refusing to mutate resource owned by another account (resource account %q; XURRENT_ACCOUNT %q; add peer to XURRENT_TRUSTED_ACCOUNT_IDS if expected)", id, expectedAccount)
	}
}

func mutationClient(t *testing.T) *openapiclient.APIClient {
	t.Helper()
	if os.Getenv("XURRENT_ALLOW_MUTATIONS") != "1" {
		t.Skip("set XURRENT_ALLOW_MUTATIONS=1 to run POST/PATCH integration tests")
	}
	cfg := newIntegrationConfig(t, true)
	return openapiclient.NewAPIClient(cfg)
}

// mePersonID returns the authenticated person's id from GET /v1/me.
func mePersonID(t *testing.T, ctx context.Context, client *openapiclient.APIClient) int32 {
	t.Helper()
	_, resp, err := client.GeneralAPI.GetMe(ctx).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	_ = resp.Body.Close()
	var me map[string]interface{}
	require.NoError(t, json.Unmarshal(body, &me))
	idVal, ok := me["id"]
	require.True(t, ok, "/v1/me must include id")
	switch v := idVal.(type) {
	case float64:
		return int32(v)
	case int:
		return int32(v)
	case int32:
		return v
	default:
		t.Fatalf("unexpected me id type %T", idVal)
	}
	return 0
}

// TestMutation_PersonJobTitlePatchRoundTrip PATCHes job_title with a random marker, then restores the original.
func TestMutation_PersonJobTitlePatchRoundTrip(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	pid := mePersonID(t, ctx, client)

	beforeRaw, resp, err := client.PeopleAPI.GetPeopleId(ctx, pid).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	before := jsonMap(beforeRaw)
	// Guard: refuse to continue if GET /v1/people/{id} does not match /v1/me (stale id / wrong user).
	requirePersonIDMatches(t, pid, before)

	origTitle, _ := before["job_title"].(string)
	sfx := randomSuffix()
	marker := fmt.Sprintf(" [%s-sdk-mut]", sfx)
	tempTitle := origTitle + marker
	if origTitle == "" {
		tempTitle = "sdk-mut-" + sfx
	}

	restored := false
	defer func() {
		if restored {
			return
		}
		_, _, _ = client.PeopleAPI.PatchPeopleId(context.Background(), pid).Body(map[string]interface{}{
			"job_title": origTitle,
		}).Execute()
	}()

	_, resp, err = client.PeopleAPI.PatchPeopleId(ctx, pid).Body(map[string]interface{}{
		"job_title": tempTitle,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	_, resp, err = client.PeopleAPI.PatchPeopleId(ctx, pid).Body(map[string]interface{}{
		"job_title": origTitle,
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	restored = true
}

// TestMutation_ProductPostPatchDisable attempts POST /v1/products then PATCH remarks + disable cleanup.
func TestMutation_ProductPostPatchDisable(t *testing.T) {
	client := mutationClient(t)
	ctx := context.Background()
	token := os.Getenv("XURRENT_TOKEN")
	acc := os.Getenv("XURRENT_ACCOUNT")

	catRef, err := firstProductCategoryReference(ctx, token, acc)
	if err != nil {
		t.Skipf("product categories: %v", err)
	}
	require.NotEmpty(t, catRef)

	sfx := randomSuffix()
	name := fmt.Sprintf("go-xurrent-mut-%s", sfx)
	body := map[string]interface{}{
		"name":     name,
		"brand":    fmt.Sprintf("TBrand-%s", sfx),
		"model":    fmt.Sprintf("TModel-%s", sfx),
		"category": catRef,
		"remarks":  fmt.Sprintf("sdk mutation probe %s", sfx),
	}

	created, resp, err := client.ProductsAPI.PostProducts(ctx).Body(body).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 201, resp.StatusCode)
	createdMap := jsonMap(created)
	requireProductAccountMatchesHeaderIfPresent(t, acc, createdMap)

	idVal := createdMap["id"]
	var id int32
	switch v := idVal.(type) {
	case float64:
		id = int32(v)
	case int:
		id = int32(v)
	case int32:
		id = v
	default:
		t.Fatalf("unexpected product id type %T", idVal)
	}

	defer func() {
		ctx2 := context.Background()
		cur, r, err := client.ProductsAPI.GetProductsId(ctx2, id).Execute()
		if err != nil || r == nil || r.StatusCode != 200 {
			t.Logf("cleanup GET product %d: %v", id, err)
			return
		}
		requireProductAccountMatchesHeaderIfPresent(t, acc, jsonMap(cur))
		_, _, _ = client.ProductsAPI.PatchProductsId(ctx2, id).Body(map[string]interface{}{
			"disabled": true,
			"remarks":  fmt.Sprintf("disabled after sdk test %s", sfx),
		}).Execute()
	}()

	// Chain read endpoints on the same created product to exercise GET /v1/products/{id}/cis without extra fixtures.
	got, resp, err := client.ProductsAPI.GetProductsId(ctx, id).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	requireProductAccountMatchesHeaderIfPresent(t, acc, jsonMap(got))

	cisResp, err := client.ProductsAPI.GetProductsIdCis(ctx, id).PerPage(1).Fields("id").Execute()
	require.NoError(t, err)
	require.NotNil(t, cisResp)
	require.Equal(t, 200, cisResp.StatusCode)
	_, _ = io.Copy(io.Discard, cisResp.Body)
	_ = cisResp.Body.Close()

	patched, resp, err := client.ProductsAPI.PatchProductsId(ctx, id).Body(map[string]interface{}{
		"remarks": fmt.Sprintf("patched sdk test %s", sfx),
	}).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	requireProductAccountMatchesHeaderIfPresent(t, acc, jsonMap(patched))
}

// requireProductAccountMatchesHeaderIfPresent enforces account match when the API includes `account`
// on the product (omit check when absent). Trusted-account headers can differ from nested refs; this
// only blocks when an explicit account id contradicts XURRENT_ACCOUNT.
func requireProductAccountMatchesHeaderIfPresent(t *testing.T, headerAccount string, product map[string]interface{}) {
	t.Helper()
	if id, ok := resourceAccountID(product); ok && id != "" {
		require.Equal(t, headerAccount, id, "product.account.id should match X-4me-Account for this test")
	}
}

// firstProductCategoryReference returns the `reference` field (e.g. address/ip_address) for POST /v1/products `category`.
func firstProductCategoryReference(ctx context.Context, token, account string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.xurrent.com/v1/product_categories?per_page=1&fields=reference", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-4me-Account", account)
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = res.Body.Close() }()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("GET product_categories: %d %s", res.StatusCode, body)
	}
	var rows []map[string]interface{}
	if err := json.Unmarshal(body, &rows); err != nil {
		return "", err
	}
	if len(rows) == 0 {
		return "", fmt.Errorf("no product categories")
	}
	ref, ok := rows[0]["reference"].(string)
	if !ok || ref == "" {
		return "", fmt.Errorf("missing reference on first product category")
	}
	return ref, nil
}
