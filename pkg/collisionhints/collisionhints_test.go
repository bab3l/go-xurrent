package collisionhints

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetectIdentifierCollision_allKinds(t *testing.T) {
	t.Parallel()
	kinds := []Kind{
		KindTeam, KindSite, KindService, KindPerson, KindProduct,
		KindOrganization, KindCI, KindSLA,
	}
	for _, k := range kinds {
		k := k
		t.Run(string(k), func(t *testing.T) {
			t.Parallel()
			p, err := DetectIdentifierCollision(k, "")
			require.NoError(t, err)
			require.Equal(t, k, p.Entity)
			require.NotEmpty(t, p.Steps)
			for _, s := range p.Steps {
				require.NotEmpty(t, s.Method)
				require.NotEmpty(t, s.Path)
			}
			require.NotEmpty(t, p.RenamePolicy.AvoidWhen)
			require.NotEmpty(t, p.RenamePolicy.OKWhen)
		})
	}
}

func TestDetectIdentifierCollision_unknownKind(t *testing.T) {
	t.Parallel()
	_, err := DetectIdentifierCollision(Kind("unknown"), "name")
	require.Error(t, err)
}

func TestDetectIdentifierCollision_defaultsFieldToName(t *testing.T) {
	t.Parallel()
	p, err := DetectIdentifierCollision(KindTeam, "")
	require.NoError(t, err)
	require.Equal(t, "name", p.Field)
}

func TestDetectIdentifierCollision_primaryEmailField(t *testing.T) {
	t.Parallel()
	p, err := DetectIdentifierCollision(KindPerson, "primary_email")
	require.NoError(t, err)
	require.Equal(t, "primary_email", p.Field)
	require.Len(t, p.Steps, 3)
}

func TestDetectIdentifierCollision_personSourceComposite(t *testing.T) {
	t.Parallel()
	p, err := DetectIdentifierCollision(KindPerson, "source_sourceid")
	require.NoError(t, err)
	require.Equal(t, "source_sourceid", p.Field)
	require.Len(t, p.Steps, 3)
	require.Contains(t, p.Steps[0].QueryOrBody, "sourceID")
}

func TestDetectIdentifierCollision_productBrandComposite(t *testing.T) {
	t.Parallel()
	p, err := DetectIdentifierCollision(KindProduct, "brand_productid")
	require.NoError(t, err)
	require.Equal(t, "brand_productid", p.Field)
	require.Len(t, p.Steps, 2)
	require.Contains(t, p.Steps[0].Path, "/products/enabled")
}

func TestDetectIdentifierCollision_productProductIDOnly(t *testing.T) {
	t.Parallel()
	p, err := DetectIdentifierCollision(KindProduct, "productid")
	require.NoError(t, err)
	require.Equal(t, "productid", p.Field)
	require.Len(t, p.Steps, 2)
	require.Contains(t, p.Steps[0].QueryOrBody, "productID")
}

func TestDetectIdentifierCollision_personNameField(t *testing.T) {
	t.Parallel()
	p, err := DetectIdentifierCollision(KindPerson, "name")
	require.NoError(t, err)
	require.Equal(t, "name", p.Field)
	require.Len(t, p.Steps, 2)
	require.Contains(t, p.Steps[0].QueryOrBody, "name=")
}

func TestParseKind_aliases(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		in   string
		want Kind
	}{
		{"teams", KindTeam},
		{"CI", KindCI},
		{"orgs", KindOrganization},
		{"people", KindPerson},
	} {
		got, err := ParseKind(tc.in)
		require.NoError(t, err)
		require.Equal(t, tc.want, got)
	}
	_, err := ParseKind("nope")
	require.Error(t, err)
}
