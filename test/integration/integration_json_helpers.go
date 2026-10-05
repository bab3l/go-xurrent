//go:build integration

package integration_test

import (
	"encoding/json"
	"testing"
)

// jsonMap converts a generated model (struct pointer) to a generic map for legacy helpers
// that expect map[string]interface{} (e.g. requireResourceOwnedByEnvAccount, mapIDInt32).
func jsonMap(v interface{}) map[string]interface{} {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil
	}
	return m
}

// mapIDInt32Any supports both map responses and typed *PostFoo201Response models.
func mapIDInt32Any(v interface{}) (int32, bool) {
	if v == nil {
		return 0, false
	}
	if m, ok := v.(map[string]interface{}); ok {
		return mapIDInt32(m)
	}
	return mapIDInt32(jsonMap(v))
}

func requireResourceOwnedByEnvAccountAny(t *testing.T, expectedAccount string, v interface{}) {
	t.Helper()
	requireResourceOwnedByEnvAccount(t, expectedAccount, jsonMap(v))
}

// int32IDFromSliceFirst converts the first row of a typed list response (e.g. GetTeams200ResponseInner) via JSON.
func int32IDFromSliceFirst(row interface{}) (int32, bool) {
	return mapIDInt32(jsonMap(row))
}
