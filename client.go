package gohome

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client   *http.Client
	endpoint string
	port     string
}

func NewClient(client *http.Client, endpoint string) *Client {
	return &Client{
		client:   client,
		endpoint: endpoint,
	}
}
func (c *Client) url(path string) string {
	return fmt.Sprintf("http://%s/%s", c.endpoint, path)
}
func (c *Client) get(path string) (body []byte, err error) {
	req, err := http.NewRequest("GET", c.url(path), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func (c *Client) DeviceInfo() (*DeviceInfo, error) {
	body, err := c.get(Info.String())
	if err != nil {
		return &DeviceInfo{}, err
	}
	z := &DeviceInfo{}
	err = json.Unmarshal(body, &z)
	if err != nil {
		return &DeviceInfo{}, err
	}
	return z, nil
}
