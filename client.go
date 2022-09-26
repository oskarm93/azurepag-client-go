package azurepag

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const BaseURL string = "https://api.azrbac.mspim.azure.com/api/v2"

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func NewClient(token *string) *Client {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    BaseURL,
		Token:      *token,
	}
	return &c
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", "Bearer", c.Token))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
