// public async Task RegisterGroupAsync(Guid objectId, CancellationToken cancel)
// {
// 	var uri = $"{BaseUri}/privilegedAccess/aadGroups/resources/register";
// 	var request = new HttpRequestMessage(HttpMethod.Post, uri);
// 	var payload = new
// 	{
// 		externalId = objectId
// 	};
// 	var payloadJson = JsonSerializer.Serialize(payload);
// 	request.Content = new StringContent(payloadJson, Encoding.UTF8, "application/json");
// 	await SendRequestAsync(request, cancel);
// }

package azurepag

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) RegisterGroup(objectID string) error {
	reqBody := struct {
		externalID string `json:"externalId"`
	}{
		externalID: objectID,
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
	if err != nil {
		return err
	}
	return nil
}
