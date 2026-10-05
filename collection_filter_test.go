package xurrent

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCollectionFilterValues(t *testing.T) {
	v := CollectionFilterValues(map[string]string{
		"status":         "assigned",
		"next_target_at": ">2011-10-01T00:00:00Z",
		"team_id":        "!4",
	})
	require.Equal(t, "assigned", v.Get("status"))
	require.Equal(t, ">2011-10-01T00:00:00Z", v.Get("next_target_at"))
	require.Equal(t, "!4", v.Get("team_id"))
}

func TestMergeQuery(t *testing.T) {
	a := url.Values{"per_page": {"1"}, "fields": {"id,name"}}.Encode()
	b := CollectionFilterValues(map[string]string{"status": "open"}).Encode()
	require.Contains(t, MergeQuery(a, b), "per_page=1")
	require.Contains(t, MergeQuery(a, b), "status=open")
}

func TestFieldsParam(t *testing.T) {
	require.Equal(t, "id,name,subject", FieldsParam("id", "name", "subject"))
}

func TestListCollectionQuery(t *testing.T) {
	q := ListCollectionQuery(50, "id,name", map[string]string{"name": "foo"})
	require.Equal(t, "50", q.Get("per_page"))
	require.Equal(t, "id,name", q.Get("fields"))
	require.Equal(t, "foo", q.Get("name"))
}

func TestCloneURLValues(t *testing.T) {
	q := url.Values{"a": {"1"}, "b": {"2", "3"}}
	c := CloneURLValues(q)
	require.Equal(t, q.Encode(), c.Encode())
	c.Set("a", "x")
	require.Equal(t, "1", q.Get("a"))
}

func TestWithSearchAfter(t *testing.T) {
	q := ListCollectionQuery(10, "id", map[string]string{"name": "n"})
	q2 := WithSearchAfter(q, "cursor")
	require.Equal(t, "cursor", q2.Get("search_after"))
	require.Empty(t, q.Get("search_after"))
	require.Empty(t, WithSearchAfter(q, "").Get("search_after"))
}
