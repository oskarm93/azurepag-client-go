package azurepag

type RegisterGroupApiRequest struct {
	ExternalID string `json:"externalId"`
}

type RoleDefinition struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type RoleDefinitionsApiResponse struct {
	RoleDefinitions []RoleDefinition `json:"value"`
}

type RoleAssignmentRequest struct {
	ID               string `json:"id"`
	ResourceID       string `json:"resourceId"`
	RoleDefinitionID string `json:"roleDefinitionId"`
	SubjectID        string `json:"subjectId"`
	AssignmentState  string `json:"assignmentState"`
}

type RoleAssignmentRequestsApiResponse struct {
	RoleAssignmentRequests []RoleAssignmentRequest `json:"value"`
}

type RoleAssignmentRequestApiRequest struct {
	ResourceID       string `json:"resourceId"`
	RoleDefinitionID string `json:"roleDefinitionId"`
	SubjectID        string `json:"subjectId"`
	AssignmentState  string `json:"assignmentState"`
	Type             string `json:"type"`
}

type RoleSettingsExpirationRuleSetting struct {
	MaximumGrantPeriodInMinutes int  `json:"maximumGrantPeriodInMinutes"`
	PermanentAssignment         bool `json:"permanentAssignment"`
}

type RoleSettingsJustificationRuleSetting struct {
	Required bool `json:"required"`
}

type RoleSettingsMfaRuleSetting struct {
	MFARequired bool `json:"mfaRequired"`
}

type RoleSettingsTicketingRuleSetting struct {
	TicketingRequired bool `json:"ticketingRequired"`
}

type RoleSettingsRule struct {
	RuleIdentifier string `json:"ruleIdentifier"`
	Setting        string `json:"setting"`
}

type LifecycleManagement struct {
	Caller            string             `json:"caller"`
	Operation         string             `json:"operation"`
	Level             string             `json:"level"`
	RoleSettingsRules []RoleSettingsRule `json:"value"`
}

type RoleSettings struct {
	ID                  string                `json:"id"`
	ResourceID          string                `json:"resourceId"`
	RoleDefinitionID    string                `json:"roleDefinitionId"`
	LifecycleManagement []LifecycleManagement `json:"lifeCycleManagement"`
}

type RoleSettingsApiResponse struct {
	RoleSettingsList []RoleSettings `json:"value"`
}
