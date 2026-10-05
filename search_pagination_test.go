package xurrent

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchNextPageToken(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name   string
		header string
		want   string
	}{
		{"empty", "", ""},
		{"nil response", "_nil_", ""},
		{"opaque", "eyJxdWVyeSI6IndpbmRvd3MifQ", "eyJxdWVyeSI6IndpbmRvd3MifQ"},
		{
			"absolute URL",
			"https://api.xurrent.com/v1/search?q=windows&per_page=25&page=nextTok",
			"nextTok",
		},
		{
			"path only",
			"/v1/search?q=windows&page=abc%2Bdef",
			"abc+def",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var resp *http.Response
			if tc.name != "nil response" {
				resp = &http.Response{Header: make(http.Header)}
				if tc.header != "" {
					resp.Header.Set(SearchNextPageHeader, tc.header)
				}
			}
			require.Equal(t, tc.want, SearchNextPageToken(resp))
		})
	}
}

func TestWithSearchPage(t *testing.T) {
	t.Parallel()
	q := url.Values{}
	q.Set("q", "windows")
	q.Set("per_page", "10")

	out := WithSearchPage(q, "tok1")
	require.Equal(t, "tok1", out.Get("page"))
	require.Equal(t, "windows", out.Get("q"))
	require.Equal(t, "", q.Get("page"))

	cleared := WithSearchPage(out, "")
	require.Equal(t, "", cleared.Get("page"))
	require.Equal(t, "windows", cleared.Get("q"))
}

func TestForEachSearchPage(t *testing.T) {
	var n int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/v1/search", r.URL.Path)
		n++
		switch n {
		case 1:
			require.Equal(t, "", r.URL.Query().Get("page"))
			w.Header().Set(SearchNextPageHeader, "/v1/search?q=windows&page=second")
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`[{"id":1}]`))
		case 2:
			require.Equal(t, "second", r.URL.Query().Get("page"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`[{"id":2}]`))
		default:
			t.Fatal("unexpected extra request")
		}
	}))
	defer srv.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: srv.URL, Description: "t"}}
	cfg.AddDefaultHeader("Authorization", "Bearer t")
	c := NewAPIClient(cfg)

	q := url.Values{}
	q.Set("q", "windows")
	var bodies int
	err := c.ForEachSearchPage(context.Background(), q, 10, func(body []byte, resp *http.Response) error {
		bodies++
		require.Equal(t, http.StatusOK, resp.StatusCode)
		return nil
	})
	require.NoError(t, err)
	require.Equal(t, 2, bodies)
	require.Equal(t, 2, n)
}

func TestForEachSearchPage_HTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"bad"}`))
	}))
	defer srv.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: srv.URL, Description: "t"}}
	cfg.AddDefaultHeader("Authorization", "Bearer t")
	c := NewAPIClient(cfg)

	err := c.ForEachSearchPage(context.Background(), url.Values{"q": {"x"}}, 5, func([]byte, *http.Response) error {
		return nil
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "HTTP 400")
}

func TestForEachSearchPage_truncationError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(SearchNextPageHeader, "/v1/search?q=x&page=more")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[]`))
	}))
	defer srv.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: srv.URL, Description: "t"}}
	cfg.AddDefaultHeader("Authorization", "Bearer t")
	c := NewAPIClient(cfg)

	err := c.ForEachSearchPage(context.Background(), url.Values{"q": {"x"}}, 1, func([]byte, *http.Response) error {
		return nil
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "max pages")
}
