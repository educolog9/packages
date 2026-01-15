package types

// TenantContext represents the tenant information resolved from X-Tenant-Domain header.
// It contains the white label configuration for the current request.
type TenantContext struct {
	Domain         string `json:"domain"`
	OrganizationID string `json:"organizationId"`
	WhiteLabelID   string `json:"whiteLabelId"`
	HasWhiteLabel  bool   `json:"hasWhiteLabel"`
	Status         string `json:"status"` // "active" | "inactive" | "pending"
	BackofficeUrl  string `json:"backofficeUrl"`
	WebUrl         string `json:"webUrl"`
}

// IsActive checks if the tenant's white label is active.
func (tc *TenantContext) IsActive() bool {
	return tc.Status == "active"
}

// IsInactive checks if the tenant's white label is inactive.
func (tc *TenantContext) IsInactive() bool {
	return tc.Status == "inactive"
}
