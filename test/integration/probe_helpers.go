//go:build integration

package integration_test

import (
	"context"
	"testing"

	openapiclient "github.com/xurrent/go-xurrent"
)

// firstEnabledOrganizationID returns an organization id from GET /v1/organizations?state=enabled
// (provider must be an **enabled** org — disabled orgs return 422 "Service provider must be enabled").
func firstEnabledOrganizationID(t *testing.T, ctx context.Context, client *openapiclient.APIClient) (int32, bool) {
	t.Helper()
	rows, resp, err := client.OrganizationsAPI.GetOrganizations(ctx).State("enabled").PerPage(25).Fields("id").Execute()
	if err == nil && resp != nil && resp.StatusCode == 200 {
		for i := range rows {
			if rows[i].Id != nil {
				return int32(*rows[i].Id), true
			}
		}
	}
	return 0, false
}

// firstEnabledOrganizationIDOtherThan returns an enabled organization id different from exclude
// (e.g. SLA customer vs service provider). Uses listOrganizations state=enabled.
func firstEnabledOrganizationIDOtherThan(t *testing.T, ctx context.Context, client *openapiclient.APIClient, exclude int32) (int32, bool) {
	t.Helper()
	rows, resp, err := client.OrganizationsAPI.GetOrganizations(ctx).State("enabled").PerPage(100).Fields("id").Execute()
	if err != nil || resp == nil || resp.StatusCode != 200 {
		return 0, false
	}
	for i := range rows {
		if rows[i].Id != nil {
			id := int32(*rows[i].Id)
			if id != exclude {
				return id, true
			}
		}
	}
	return 0, false
}
