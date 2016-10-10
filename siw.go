package siw

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTimeout = 30 * time.Second
)

type Client struct {
	BaseURL *url.URL
	Timeout time.Duration
}

func NewClient(webhookURL string) *Client {
	u, _ := url.Parse(webhookURL)
	return &Client{
		BaseURL: u,
		Timeout: defaultTimeout,
	}
}

func (c *Client) PostPayload(payload Payload) (string, error) {
	var reqBody bytes.Buffer
	err := json.NewEncoder(&reqBody).Encode(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.BaseURL.String(), &reqBody)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	client.Timeout = c.Timeout
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", fmt.Errorf("API result failed: %s", resp.Status)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), err

}

type Payload struct {
	Channel   string `json:"channel"`
	UserName  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
}
