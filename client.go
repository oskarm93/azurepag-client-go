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
	UserAgent  string
}

func NewClient(token *string, userAgent *string) *Client {
	c := Client{
		HTTPClient: &http.Client{Timeout: 1 * time.Minute},
		BaseURL:    BaseURL,
		Token:      *token,
		UserAgent:  *userAgent,
	}
	return &c
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", "Bearer", c.Token))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	success := res.StatusCode >= 200 && res.StatusCode < 300
	if !success {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
