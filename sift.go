// Copyright 2016 Nevio Vesic
// Please check out LICENSE file for more information about what you CAN and what you CANNOT do!
// MIT License

package sift

import (
	"encoding/json"
	"fmt"
)

type Sift struct {
	Client
}

// Track - Send tracking event towards sift science
func (s *Sift) Track(event string, params map[string]interface{}, args map[string]interface{}) (*Response, error) {
	// The name of the event to send. This can either be a reserved
	// event name such as "$transaction" or "$create_order" or a custom event
	// name (that does not start with a $).
	params["$type"] = event

	// Whether the API response should include a score for this
	// user (the score will be calculated using this event).  This feature must be
	// enabled for your account in order to use it.  Please contact
	// support@siftscience.com if you are interested in using this feature.
	if score, ok := args["return_score"]; ok {
		params["return_score"] = score.(bool)
	}

	// Whether the API response should include actions in the response. For
	// more information on how this works, please visit the tutorial at:
	// https://siftscience.com/resources/tutorials/formulas
	if action, ok := args["return_action"]; ok {
		params["return_action"] = action.(bool)
	}

	return s.HttpRequest("POST", s.GetEventsUrl(), params)
}

// Score - Get out user score
func (s *Sift) Score(userID string) (*Response, error) {
	return s.HttpRequest("GET", fmt.Sprintf("%s?api_key=%s", s.GetScoreUrl(userID), s.ApiKey), map[string]interface{}{})
}

// Label - Request labeling for specific user
func (s *Sift) Label(userID string, params map[string]interface{}) (*Response, error) {
	return s.HttpRequest("POST", s.GetLabelUrl(userID), params)
}

// UnLabel - Request unlabeling for specific user
func (s *Sift) UnLabel(userID string) (*Response, error) {
	return s.HttpRequest("DELETE", fmt.Sprintf("%s?api_key=%s", s.GetLabelUrl(userID), s.ApiKey), map[string]interface{}{})
}

// New - Return Sift API client.
func New(apiKey string) *Sift {
	return &Sift{
		Client: Client{
			Config: Config{
				ApiUrl:     API_URL,
				ApiKey:     apiKey,
				ApiVersion: API_VERSION,
				Timeout:    TIMEOUT,
			},
		},
	}
}

// NewFromJSON - Decode JSON file into new Sift client.
func NewFromJSON(config []byte) (*Sift, error) {
	var client Sift

	if err := json.Unmarshal(config, &client); err != nil {
		return &client, err
	}

	if client.ApiUrl == "" {
		client.ApiUrl = API_URL
	}

	if client.ApiVersion == 0 {
		client.ApiVersion = API_VERSION
	}

	if client.Timeout == 0 {
		client.Timeout = TIMEOUT
	}

	return &client, nil
}
