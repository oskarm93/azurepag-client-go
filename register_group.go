package azurepag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) RegisterGroup(objectID string) error {
	reqBody := struct {
		ExternalID string `json:"externalId"`
	}{
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
	_ = body // fuck this language
	if err != nil {
		return err
	}
	return nil
}
