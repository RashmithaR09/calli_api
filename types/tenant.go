package types

type TenantCreateRequest struct {
	TenantName  string `json:"tenant_name"`
	DisplayName string `json:"display_name"`
}

type TenantUpdateResponse struct {
	ID          uint   `json:"id"`
	TenantName  string `json:"tenant_name"`
	DisplayName string `json:"display_name"`
}

type TenantUpdateRequest struct {
	DisplayName string `json:"display_name"`
}
