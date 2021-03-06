// Copyright 2016 Nevio Vesic
// Please check out LICENSE file for more information about what you CAN and what you CANNOT do!
// MIT License

package sift

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// IsOK - Check status of response. Is it error'ed or succeed?
func (r *Response) IsOK() bool {
	if _, ok := NoContentStatusCodes[r.HTTPStatusCode]; ok {
		return 204 == r.HTTPStatusCode
	}

	return r.Status == 0
}

// Config - Configuration struct used once per Sift Environment
type Config struct {
	ApiUrl     string        `json:"api_url"`
	ApiVersion int           `json:"api_version"`
	ApiKey     string        `json:"api_key"`
	Timeout    time.Duration `json:"timeout"`
}

// Client - Designed to be as connection point between code and Sift Science API
type Client struct {
	Config `json:"config"`
}

// SetApiUrl - Set Sift API Url. Should not be modified unless you know
// what you're doing. Default API Url can be seen in constants.go
func (c *Client) SetApiUrl(url string) {
	c.ApiUrl = url
}

// SetApiKey - Set Sift API key. You can find your keys at https://siftscience.com/console/developer/api-keys
// Pay closer attention to Production/Sandbox Mode as keys are different.
func (c *Client) SetApiKey(key string) {
	c.ApiKey = key
}

// SetApiVersion - Set Sift API version. Default API Url can be seen in constants.go
func (c *Client) SetApiVersion(version int) {
	c.ApiVersion = version
}

// SetTimeout - Set new API request timeout. Default API Timeout can be
// seen in constants.go
func (c *Client) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

// UserAgent - Returns User Agent that will be used with request towards Sift Science
func (c *Client) UserAgent() string {
	return fmt.Sprintf("SiftScience/%d sift-golang/%s", c.ApiVersion, VERSION)
}

// GetEventsUrl - Returning full url to sift science events API
func (c *Client) GetEventsUrl() string {
	return c.apiUrl("events")
}

// GetScoreUrl - Returning full url to sift science score API
func (c *Client) GetScoreUrl(userId string) string {
	return c.apiUrl(fmt.Sprintf("score/%s", userId))
}

// GetScoreUrl - Returning full url to sift science label API
func (c *Client) GetLabelUrl(userId string) string {
	return c.apiUrl(fmt.Sprintf("users/%s/labels", userId))
}

// HttpRequest - Doing HTTP requests towards sift science API
// Should be used only if there's no other choice. There are helper methods defined
// in sift.go for tracking, labeling, etc...
func (c *Client) HttpRequest(method string, url string, params map[string]interface{}) (*Response, error) {
	if _, ok := AvailableMethods[strings.ToUpper(method)]; !ok {
		return nil, fmt.Errorf("Passed request (method: %s) is not supported by Sift Science yet!", method)
	}

	var rbytes io.Reader
	var err error

	// Set this here so it acts as global configuration
	if method != "GET" {
		params["$api_key"] = c.ApiKey

		b, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		rbytes = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, rbytes)
	req.Header.Set("User-Agent", c.UserAgent())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(c.Timeout * time.Second),
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	r := Response{
		HTTPStatus:       resp.Status,
		HTTPStatusCode:   resp.StatusCode,
		HTTPStatusHeader: resp.Header,
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &r, err
	}

	// If it's 204 or 304 return what we have, without error
	if _, ok := NoContentStatusCodes[resp.StatusCode]; ok {
		return &r, nil
	}

	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return &r, err
	}

	if r.IsOK() == false {
		return &r, errors.New(r.ErrorMessage)
	}

	return &r, nil
}

// apiUrl - private helper to get base API url including uri that is provided
// as argument
func (c *Client) apiUrl(uri string) string {
	// Make sure correct API version is set. Fail-safe and to setup defaults
	if c.ApiVersion == 0 {
		c.SetApiVersion(API_VERSION)
	}

	// Make sure correct API URL is set. Fail-safe and setup defaults.
	if c.ApiUrl == "" {
		c.SetApiUrl(API_URL)
	}

	return fmt.Sprintf("%s/v%d/%s", c.ApiUrl, c.ApiVersion, uri)
}
