package azurepag

type RegistrationRequest struct {
	externalID string `json:"externalId"`
}

type GovernanceResource struct {
	ID string `json:"id"`
}
