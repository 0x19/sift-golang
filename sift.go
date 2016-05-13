//
//

package sift

import "encoding/json"

// Sift -
type Sift struct {
	Client
}

// Track -
func (s *Sift) Track(event string, properties map[string]interface{}, args map[string]interface{}) (*Response, error) {
	// The name of the event to send. This can either be a reserved
	// event name such as "$transaction" or "$create_order" or a custom event
	// name (that does not start with a $).
	properties["$type"] = event

	// Whether the API response should include a score for this
	// user (the score will be calculated using this event).  This feature must be
	// enabled for your account in order to use it.  Please contact
	// support@siftscience.com if you are interested in using this feature.
	if score, ok := args["return_score"]; ok {
		properties["return_score"] = score.(bool)
	}

	// Whether the API response should include actions in the response. For
	// more information on how this works, please visit the tutorial at:
	// https://siftscience.com/resources/tutorials/formulas
	if action, ok := args["return_action"]; ok {
		properties["return_action"] = action.(bool)
	}

	return s.GetRequest("POST", s.GetEventsUrl(), properties)
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
