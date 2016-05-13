// Copyright 2016 Nevio Vesic
// Please check out LICENSE file for more information about what you CAN and what you CANNOT do!
// MIT License

package sift

var (
	// NoContentStatusCodes - Here as helper to define if request was successful or not
	NoContentStatusCodes = map[int]string{204: "", 304: ""}

	// AvailableMethods - List of all methods that Sift Science API accepts
	AvailableMethods = map[string]string{"GET": "get", "POST": "post", "DELETE": "delete"}

	// ErrorCodes A successful API request will respond with an HTTP 200. An invalid API
	// request will respond with an HTTP 400. The response body will be a JSON
	// object describing why the request failed.
	// These are JSON error response codes in case you need them
	ErrorCodes = map[int]string{
		-4:  "Service currently unavailable. Please try again later.",
		-3:  "Server-side timeout processing request. Please try again later.",
		-2:  "Unexpected server-side error",
		-1:  "Unexpected server-side error",
		0:   "Success",
		51:  "Invalid API Key. Please check your credentials and try again.",
		52:  "Invalid characters in field name",
		53:  "Invalid characters in field value",
		54:  "Specified user_id has no scoreable events",
		55:  "Missing required field",
		56:  "Invalid JSON in request",
		57:  "Invalid HTTP body",
		60:  "Rate limited; too many events have been received in a short period of time",
		104: "Invalid API version",
		105: "Not a valid reserved field",
	}
)
