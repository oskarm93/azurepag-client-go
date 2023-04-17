package azurepag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) RegisterGroup(objectID string) error {
	reqBody := RegisterGroupApiRequest{
		ExternalID: objectID,
	}
	rb, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/privilegedAccess/aadGroups/resources/register", c.BaseURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	_ = body
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetRoleDefinitions(objectID string) ([]RoleDefinition, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/privilegedAccess/aadGroups/resources/%s/roleDefinitions", c.BaseURL, objectID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := RoleDefinitionsApiResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.RoleDefinitions, nil
}

func (c *Client) GetRoleDefinition(objectID string, roleName string) (*RoleDefinition, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/privilegedAccess/aadGroups/resources/%s/roleDefinitions?$filter=(displayName%%20eq%%20%%27%s%%27)", c.BaseURL, objectID, roleName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := RoleDefinitionsApiResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response.RoleDefinitions[0], nil
}

func (c *Client) GetRoleAssignmentRequest(objectID string, subjectID string, roleDefinitionID string, assignmentState string) (*RoleAssignmentRequest, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/privilegedAccess/aadGroups/roleAssignments?$filter=(roleDefinition/resource/id%%20eq%%20%%27%s%%27)+and+(roleDefinition/id%%20eq%%20%%27%s%%27)+and+(subjectId%%20eq%%20%%27%s%%27)+and+(assignmentState%%20eq%%20%%27%s%%27)", c.BaseURL, objectID, roleDefinitionID, subjectID, assignmentState), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := RoleAssignmentRequestsApiResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response.RoleAssignmentRequests[0], nil
}

func (c *Client) CreateRoleAssignmentRequest(objectID string, subjectID string, roleDefinitionID string, assignmentState string) (*RoleAssignmentRequest, error) {
	reqBody := RoleAssignmentRequestApiRequest{
		ResourceID:       objectID,
		RoleDefinitionID: roleDefinitionID,
		SubjectID:        subjectID,
		AssignmentState:  assignmentState,
		Type:             "AdminAdd",
	}
	rb, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/privilegedAccess/aadGroups/roleAssignmentRequests", c.BaseURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := RoleAssignmentRequest{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteRoleAssignmentRequest(objectID string, subjectID string, roleDefinitionID string, assignmentState string) error {
	reqBody := RoleAssignmentRequestApiRequest{
		ResourceID:       objectID,
		RoleDefinitionID: roleDefinitionID,
		SubjectID:        subjectID,
		AssignmentState:  assignmentState,
		Type:             "AdminRemove",
	}
	rb, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/privilegedAccess/aadGroups/roleAssignmentRequests", c.BaseURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	_ = body
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetRoleSettings(objectID string, roleDefinitionID string) (*RoleSettings, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/privilegedAccess/aadGroups/roleSettingsv2?$filter=(resource/id+eq+%%27%s%%27)+and+(roleDefinition/id+eq+%%27%s%%27)", c.BaseURL, objectID, roleDefinitionID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := RoleSettingsApiResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response.RoleSettingsList[0], nil
}

func (c *Client) UpdateRoleSettings(roleSettings *RoleSettings) error {
	rb, err := json.Marshal(roleSettings)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/privilegedAccess/aadGroups/roleSettingsV2/%s", c.BaseURL, roleSettings.ID), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	_ = body
	if err != nil {
		return err
	}
	return nil
}
