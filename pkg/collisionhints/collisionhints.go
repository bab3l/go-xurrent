// Package collisionhints provides ordered API steps to locate rows that may still
// hold a unique identifier ("ghost" rows) after a duplicate-name or duplicate-key error.
//
// It is the programmatic counterpart to Phase C in xurrent-mcp/docs/PROBE_RESULTS.md
// and is intended for MCP tools and CLIs that guide operators without performing writes.
package collisionhints

import (
	"fmt"
	"strings"
)

// Kind identifies an API resource family (not cross-type uniqueness).
type Kind string

const (
	KindTeam         Kind = "team"
	KindSite         Kind = "site"
	KindService      Kind = "service"
	KindPerson       Kind = "person"
	KindProduct      Kind = "product"
	KindOrganization Kind = "organization"
	KindCI           Kind = "configuration_item"
	KindSLA          Kind = "sla"
)

// ParseKind maps CLI/MCP strings to Kind (accepts plural and short aliases).
func ParseKind(s string) (Kind, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "team", "teams":
		return KindTeam, nil
	case "site", "sites":
		return KindSite, nil
	case "service", "services":
		return KindService, nil
	case "person", "people":
		return KindPerson, nil
	case "product", "products":
		return KindProduct, nil
	case "organization", "organizations", "org", "orgs":
		return KindOrganization, nil
	case "configuration_item", "configurationitem", "ci", "cis":
		return KindCI, nil
	case "sla", "slas":
		return KindSLA, nil
	default:
		return "", fmt.Errorf("collisionhints: unknown entity %q (use team, site, service, person, product, organization, configuration_item, sla)", s)
	}
}

// Step is one discovery step (typically GET list or predefined collection).
type Step struct {
	Order       int    `json:"order"`
	Summary     string `json:"summary"`
	Method      string `json:"method"`
	Path        string `json:"path"`
	QueryOrBody string `json:"query_or_body,omitempty"`
	Notes       string `json:"notes,omitempty"`
}

// Plan is the full guidance for find-then-decide (read-only discovery).
type Plan struct {
	Entity       Kind         `json:"entity"`
	Field        string       `json:"field"`
	Steps        []Step       `json:"steps"`
	RenamePolicy RenamePolicy `json:"rename_policy"`
}

// RenamePolicy documents when suffix/rename is appropriate vs risky (Phase D).
type RenamePolicy struct {
	Summary   string   `json:"summary"`
	OKWhen    []string `json:"ok_when"`
	AvoidWhen []string `json:"avoid_when"`
}

// DetectIdentifierCollision returns read-only discovery steps for locating a blocking row.
// field is the duplicate field from the API error (e.g. "name", "primary_email", "serial_nr",
// "productID", "sourceID"). For name-based entities, pass "name" or "".
func DetectIdentifierCollision(entity Kind, field string) (Plan, error) {
	f := strings.TrimSpace(strings.ToLower(field))
	if f == "" {
		f = "name"
	}

	switch entity {
	case KindTeam:
		return planTeam(f), nil
	case KindSite:
		return planSite(f), nil
	case KindService:
		return planService(f), nil
	case KindPerson:
		return planPerson(f), nil
	case KindProduct:
		return planProduct(f), nil
	case KindOrganization:
		return planOrganization(f), nil
	case KindCI:
		return planCI(f), nil
	case KindSLA:
		return planSLA(f), nil
	default:
		return Plan{}, fmt.Errorf("collisionhints: unknown entity kind %q", entity)
	}
}

func defaultRenamePolicy() RenamePolicy {
	return RenamePolicy{
		Summary: "After locating the row, PATCH may free a unique value by renaming; confirm tenant policy and audit requirements first.",
		OKWhen: []string{
			"Test or disposable rows created for automation (suffix in name).",
			"Non-regulated display names with no downstream sync (teams, sites, calendars).",
		},
		AvoidWhen: []string{
			"Regulated personal data: prefer process over blind rename of real person records.",
			"Identifiers tied to external sync: person source/sourceID, HR feeds, or CMDB correlation — renaming may break integrations.",
			"Production CIs or assets under change control — use formal archive/trash workflows.",
		},
	}
}

func planTeam(field string) Plan {
	return Plan{
		Entity: KindTeam,
		Field:  field,
		Steps: []Step{
			{1, "List disabled teams (reservation often on disabled rows)", "GET", "/v1/teams/disabled", "fields=id,name,disabled", "Or GET /v1/teams?state=disabled per Filtering docs."},
			{2, "If not found, scan active teams", "GET", "/v1/teams", "state=enabled&fields=id,name", ""},
		},
		RenamePolicy: defaultRenamePolicy(),
	}
}

func planSite(field string) Plan {
	return Plan{
		Entity: KindSite,
		Field:  field,
		Steps: []Step{
			{1, "List disabled sites", "GET", "/v1/sites/disabled", "fields=id,name", "Or GET /v1/sites?state=disabled."},
			{2, "Scan enabled sites if needed", "GET", "/v1/sites", "fields=id,name", ""},
		},
		RenamePolicy: defaultRenamePolicy(),
	}
}

func planService(field string) Plan {
	return Plan{
		Entity: KindService,
		Field:  field,
		Steps: []Step{
			{1, "List services with provider/state filters", "GET", "/v1/services", "fields=id,name,disabled,provider", "Use provider_id or state per tenant."},
		},
		RenamePolicy: defaultRenamePolicy(),
	}
}

func planPerson(field string) Plan {
	f := strings.TrimSpace(strings.ToLower(field))
	if f == "" {
		f = "name"
	}
	switch f {
	case "primary_email", "email":
		return Plan{
			Entity: KindPerson,
			Field:  f,
			Steps: []Step{
				{1, "Disabled people", "GET", "/v1/people/disabled", "fields=id,primary_email,name&primary_email=…", "MCP find_identifier_candidates tries primary_email= when supported; else full list scan."},
				{2, "Enabled people", "GET", "/v1/people/enabled", "fields=id,primary_email,name", ""},
				{3, "Search (if enabled)", "GET", "/v1/search", "q=…", "Tenant-dependent; use for email fragments."},
			},
			RenamePolicy: defaultRenamePolicy(),
		}
	case "name":
		return Plan{
			Entity: KindPerson,
			Field:  f,
			Steps: []Step{
				{1, "Disabled people", "GET", "/v1/people/disabled", "fields=id,name,primary_email&name=…", "MCP find_identifier_candidates tries name= when supported; else full list scan."},
				{2, "Enabled people", "GET", "/v1/people/enabled", "fields=id,name,primary_email", ""},
			},
			RenamePolicy: defaultRenamePolicy(),
		}
	case "source_sourceid", "source+sourceid", "source_source_id", "sourceid", "source_id":
		return Plan{
			Entity: KindPerson,
			Field:  f,
			Steps: []Step{
				{1, "Disabled people (match sourceID)", "GET", "/v1/people/disabled", "fields=id,primary_email,name,sourceID", "MCP find_identifier_candidates: field source_sourceid (value_secondary=sourceID) or field sourceid (value=sourceID only). List rows cannot filter by source."},
				{2, "Enabled people", "GET", "/v1/people/enabled", "fields=id,primary_email,name,sourceID", ""},
				{3, "Confirm source", "GET", "/v1/people/{id}", "", "Detail may include source for directory/HR correlation."},
			},
			RenamePolicy: defaultRenamePolicy(),
		}
	default:
		notes := "Ghost may be disabled but still reserve email or source+sourceID."
		return Plan{
			Entity: KindPerson,
			Field:  f,
			Steps: []Step{
				{1, "Disabled people", "GET", "/v1/people/disabled", "fields=id,primary_email,name,sourceID", notes},
				{2, "Search (if enabled)", "GET", "/v1/search", "q=…", "Tenant-dependent; use for email fragments."},
			},
			RenamePolicy: defaultRenamePolicy(),
		}
	}
}

func planProduct(field string) Plan {
	f := strings.TrimSpace(strings.ToLower(field))
	if f == "" {
		f = "name"
	}
	switch f {
	case "brand_productid", "brand+productid", "brand_product_id":
		return Plan{
			Entity: KindProduct,
			Field:  f,
			Steps: []Step{
				{1, "Enabled products (brand + productID)", "GET", "/v1/products/enabled", "fields=id,name,brand,productID", "MCP find_identifier_candidates: field brand_productid, value=brand, value_secondary=productID."},
				{2, "Disabled products", "GET", "/v1/products/disabled", "fields=id,name", "Typed list may omit brand/productID; use name scan or API detail if ghost is disabled."},
			},
			RenamePolicy: defaultRenamePolicy(),
		}
	case "productid", "product_id":
		return Plan{
			Entity: KindProduct,
			Field:  f,
			Steps: []Step{
				{1, "Products (productID filter)", "GET", "/v1/products/disabled|enabled", "productID=…&fields=id,name,brand,productID", "MCP find_identifier_candidates: field productid, value=productID (falls back to full scan if filter rejected)."},
				{2, "Composite alternative", "GET", "/v1/products/enabled", "fields=id,name,brand,productID", "If uniqueness is brand-scoped, use field brand_productid with value_secondary."},
			},
			RenamePolicy: defaultRenamePolicy(),
		}
	default:
		return Plan{
			Entity: KindProduct,
			Field:  f,
			Steps: []Step{
				{1, "Disabled products", "GET", "/v1/products/disabled", "fields=id,name,brand,productID", ""},
				{2, "Enabled products", "GET", "/v1/products/enabled", "fields=id,name,brand,productID", "Composite uniqueness may involve brand+productID; use field brand_productid + value_secondary for MCP live search."},
			},
			RenamePolicy: defaultRenamePolicy(),
		}
	}
}

func planOrganization(field string) Plan {
	return Plan{
		Entity: KindOrganization,
		Field:  field,
		Steps: []Step{
			{1, "Organizations by state", "GET", "/v1/organizations", "state=disabled or state=enabled", "Disabled org may still hold name."},
		},
		RenamePolicy: defaultRenamePolicy(),
	}
}

func planCI(field string) Plan {
	return Plan{
		Entity: KindCI,
		Field:  field,
		Steps: []Step{
			{1, "List CIs by serial (case-insensitive per API docs)", "GET", "/v1/cis", "serial_nr=…&fields=id,label,serial_nr,status", "Include archived/trash states as separate queries."},
			{2, "Detail after id known", "GET", "/v1/cis/{id}", "", ""},
		},
		RenamePolicy: defaultRenamePolicy(),
	}
}

func planSLA(field string) Plan {
	return Plan{
		Entity: KindSLA,
		Field:  field,
		Steps: []Step{
			{1, "SLA collections", "GET", "/v1/slas", "inactive / active per spec", "Use tenant-specific inactive list if duplicate after disable."},
		},
		RenamePolicy: defaultRenamePolicy(),
	}
}
