package azurepag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) RegisterGroup(objectID string) error {
	reqBody := RegisterGroupRequest{
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

// public async Task<RoleDefinitionsList> GetRoleDefinitionsAsync(Guid objectId, CancellationToken cancel)
// {
// 	var uri = $"{BaseUri}/privilegedAccess/aadGroups/resources/{objectId}/roleDefinitions";
// 	var request = new HttpRequestMessage(HttpMethod.Get, uri);
// 	var result = await SendRequestAsync<RoleDefinitionsList>(request, cancel);
// 	return result;
// }

func (c *Client) GetRoleDefinitions(objectID string) ([]RoleDefinition, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/privilegedAccess/aadGroups/resources/%s/roleDefinitions", c.BaseURL, objectID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := RoleDefinitionsResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.RoleDefinitions, nil
}
