package types

// WhiteLabelType represents the type of white label (parent or child).
type WhiteLabelType string

const (
	// WhiteLabelTypeParent represents a parent white label that can manage other tenants.
	WhiteLabelTypeParent WhiteLabelType = "parent"
	// WhiteLabelTypeChild represents a child white label managed by a parent.
	WhiteLabelTypeChild WhiteLabelType = "child"
)

// TenantContext represents the tenant information resolved from X-Tenant-Domain header.
// It contains the white label configuration for the current request.
type TenantContext struct {
	Domain         string         `json:"domain"`
	OrganizationID string         `json:"organizationId"`
	WhiteLabelID   string         `json:"whiteLabelId"`
	HasWhiteLabel  bool           `json:"hasWhiteLabel"`
	Status         string         `json:"status"` // "active" | "inactive" | "pending"
	Type           WhiteLabelType `json:"type"`   // "parent" | "child"
	BackofficeUrl  string         `json:"backofficeUrl"`
	WebUrl         string         `json:"webUrl"`
	LogoURL        string         `json:"logoUrl"`
	Colors         []string       `json:"colors"`
	LoginImage     string         `json:"loginImage"`
}

// IsActive checks if the tenant's white label is active.
func (tc *TenantContext) IsActive() bool {
	return tc.Status == "active"
}

// IsInactive checks if the tenant's white label is inactive.
func (tc *TenantContext) IsInactive() bool {
	return tc.Status == "inactive"
}

// IsParent checks if the tenant is a parent white label.
func (tc *TenantContext) IsParent() bool {
	return tc.Type == WhiteLabelTypeParent
}

// IsChild checks if the tenant is a child white label.
func (tc *TenantContext) IsChild() bool {
	return tc.Type == WhiteLabelTypeChild
}
