package azurepag

type RegisterGroupRequest struct {
	ExternalID string `json:"externalId"`
}

type RoleDefinitionsResponse struct {
	RoleDefinitions []RoleDefinition `json:"value"`
}

type RoleDefinition struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}
