package ueshka

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	address = "https://api.classcard.ru"
)

// Client ...
type Client struct {
	AppVersion string
	Token      string

	client *http.Client
}

// NewClient ...
func NewClient(ver string, token string) *Client {
	return &Client{
		AppVersion: ver,
		Token:      token,

		client: &http.Client{},
	}
}

// NewRequest ...
func (c *Client) NewRequest(method, uri string, data interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, address+uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("App-Version", c.AppVersion)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	return req, nil
}

// GetDailyStat ... 1756948 2020-10-13
func (c *Client) GetDailyStat(pupilID string, dateStart, dateEnd string) (map[string][]Operation, error) {
	req, err := c.NewRequest(http.MethodGet, "/attendance/daily-stat/"+pupilID, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("dateStart", dateStart)
	q.Set("dateEnd", dateEnd)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not valid status code: %d url: %s", resp.StatusCode, req.URL)
	}

	var stat = make(map[string][]Operation)

	if err := json.NewDecoder(resp.Body).Decode(&stat); err != nil {
		return nil, err
	}

	return stat, nil
}
