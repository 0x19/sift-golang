// Copyright 2016 Nevio Vesic
// Please check out LICENSE file for more information about what you CAN and what you CANNOT do!
// MIT License

package sift

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	testApiKey         = "xxx"
	testValidApiConfig = fmt.Sprintf(`{"config": {"api_key": "%s"}}`, testApiKey)
	testApiError       = `{ "time" : 1463168681 , "status" : 51 , "request" : "{\"$api_key\":\"xxx\"}" , "error_message" : "Invalid API Key. Please check your credentials and try again."}`
	testEventOK        = `{ "status" : 0 , "error_message" : "OK" , "time" : 1463168322 , "request" : "{\"$api_key\":\"xxx\",\"$type\":\"my_custom_event\",\"$user_id\":\"test@someone.com\",\"hello\":\"world\",\"return_action\":true,\"return_score\":true}"}`
	testScoreOK        = `{ "status" : 0 , "error_message" : "OK" , "score" : 0.3275949342030485 , "user_id" : "test@someone.com" , "latest_label" : { "is_bad" : false , "time" : 1463167642} , "actions" : [ ]}`
	testLabelOK        = `{ "time" : 1463168323 , "status" : 0 , "request" : "{\"$api_key\":\"25a031275f740846\",\"$is_bad\":false}" , "error_message" : "OK"}`
)

func TestSiftStruct(t *testing.T) {
	Convey("New() returns pointer to Sift", t, func() {
		So(New(testApiKey), ShouldHaveSameTypeAs, &Sift{})
	})

	Convey("Make sure default api url is properly set", t, func() {
		s := New(testApiKey)
		So(s.ApiUrl, ShouldEqual, API_URL)
	})

	Convey("Make sure default api version is properly set", t, func() {
		s := New(testApiKey)
		So(s.ApiVersion, ShouldEqual, API_VERSION)
	})

	Convey("Make sure default api version is properly set", t, func() {
		s := New(testApiKey)
		So(s.Timeout, ShouldEqual, TIMEOUT)
	})

	Convey("Make sure api key is properly set", t, func() {
		s := New(testApiKey)
		So(s.ApiKey, ShouldEqual, testApiKey)
	})

}

func TestSiftStructJSON(t *testing.T) {
	s, err := NewFromJSON([]byte(testValidApiConfig))

	Convey("NewFromJSON() returns pointer to Sift", t, func() {
		So(err, ShouldBeNil)
		So(s, ShouldHaveSameTypeAs, &Sift{})
	})

	Convey("Make sure default api url is properly set", t, func() {
		So(s.ApiUrl, ShouldEqual, API_URL)
	})

	Convey("Make sure default api version is properly set", t, func() {
		So(s.ApiVersion, ShouldEqual, API_VERSION)
	})

	Convey("Make sure default api version is properly set", t, func() {
		So(s.Timeout, ShouldEqual, TIMEOUT)
	})

	Convey("Make sure api key is properly set", t, func() {
		So(s.ApiKey, ShouldEqual, testApiKey)
	})
}

func TestSiftClientStruct(t *testing.T) {
	s, _ := NewFromJSON([]byte(testValidApiConfig))

	Convey("SetApiUrl should set new Api Url", t, func() {
		s.SetApiUrl("http://test.url")
		So(s.ApiUrl, ShouldEqual, "http://test.url")
	})

	Convey("SetApiKey should set new Api Key", t, func() {
		s.SetApiKey("somekey")
		So(s.ApiKey, ShouldEqual, "somekey")
	})

	Convey("SetApiVersion should set new Api Version", t, func() {
		s.SetApiVersion(200)
		So(s.ApiVersion, ShouldEqual, 200)
	})

	Convey("SetTimeout should set new request timeout", t, func() {
		s.SetTimeout(20)
		So(s.Timeout, ShouldEqual, 20)
	})

	Convey("UserAgent should return properly generated agent as string", t, func() {
		So(s.UserAgent(), ShouldEqual, fmt.Sprintf("SiftScience/%d sift-golang/%s", s.ApiVersion, VERSION))
	})

	Convey("GetEventsUrl should return events full url", t, func() {
		So(s.GetEventsUrl(), ShouldEqual, fmt.Sprintf("%s/v%d/events", s.ApiUrl, s.ApiVersion))
	})

	Convey("GetLabelUrl should return label full url", t, func() {
		So(s.GetLabelUrl("abcd"), ShouldEqual, fmt.Sprintf("%s/v%d/users/%s/labels", s.ApiUrl, s.ApiVersion, "abcd"))
	})

	Convey("GetScoreUrl should return score full url", t, func() {
		So(s.GetScoreUrl("abcd"), ShouldEqual, fmt.Sprintf("%s/v%d/score/%s", s.ApiUrl, s.ApiVersion, "abcd"))
	})
}

func TestSiftClientTrackRequest(t *testing.T) {
	s, err := NewFromJSON([]byte(testValidApiConfig))

	Convey("NewFromJSON() returns pointer to Sift", t, func() {
		So(err, ShouldBeNil)
		So(s, ShouldHaveSameTypeAs, &Sift{})
	})

	Convey("Publishes new custom tracking event successfully", t, func() {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %v", r.URL)
			fmt.Fprintf(w, testEventOK)
		}))

		defer svr.Close()
		defer s.SetApiUrl(API_URL)

		s.SetApiUrl(svr.URL)

		r, err := s.Track(
			"my_custom_event",
			map[string]interface{}{
				"$user_id": "test@someone.com", "hello": "world",
			}, map[string]interface{}{
				"return_score":  true,
				"return_action": true,
			},
		)

		So(r, ShouldHaveSameTypeAs, &Response{})
		So(err, ShouldBeNil)

		So(r.IsOK(), ShouldBeTrue)

		So(r.HTTPStatusCode, ShouldEqual, 200)
		So(r.HTTPStatus, ShouldEqual, "200 OK")
		So(r.ErrorMessage, ShouldEqual, "OK")
	})

}

func TestSiftClientScoreRequest(t *testing.T) {
	s, err := NewFromJSON([]byte(testValidApiConfig))

	Convey("NewFromJSON() returns pointer to Sift", t, func() {
		So(err, ShouldBeNil)
		So(s, ShouldHaveSameTypeAs, &Sift{})
	})

	Convey("Fetches customer score", t, func() {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %v", r.URL)
			fmt.Fprintf(w, testScoreOK)
		}))

		defer svr.Close()
		defer s.SetApiUrl(API_URL)

		s.SetApiUrl(svr.URL)

		r, err := s.Score("test@someone.com")

		So(r, ShouldHaveSameTypeAs, &Response{})
		So(err, ShouldBeNil)

		So(r.IsOK(), ShouldBeTrue)

		So(r.HTTPStatusCode, ShouldEqual, 200)
		So(r.HTTPStatus, ShouldEqual, "200 OK")
		So(r.ErrorMessage, ShouldEqual, "OK")
	})
}

func TestSiftClientLabelRequest(t *testing.T) {
	s, err := NewFromJSON([]byte(testValidApiConfig))

	Convey("NewFromJSON() returns pointer to Sift", t, func() {
		So(err, ShouldBeNil)
		So(s, ShouldHaveSameTypeAs, &Sift{})
	})

	Convey("Labels customer as ok", t, func() {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %v", r.URL)
			fmt.Fprintf(w, testLabelOK)
		}))

		defer svr.Close()
		defer s.SetApiUrl(API_URL)

		s.SetApiUrl(svr.URL)

		r, err := s.Label("test@someone.com", map[string]interface{}{
			"$is_bad": false,
		})

		So(r, ShouldHaveSameTypeAs, &Response{})
		So(err, ShouldBeNil)

		So(r.IsOK(), ShouldBeTrue)

		So(r.HTTPStatusCode, ShouldEqual, 200)
		So(r.HTTPStatus, ShouldEqual, "200 OK")
		So(r.ErrorMessage, ShouldEqual, "OK")
	})
}

func TestSiftClientUnLabelRequest(t *testing.T) {
	s, err := NewFromJSON([]byte(testValidApiConfig))

	Convey("NewFromJSON() returns pointer to Sift", t, func() {
		So(err, ShouldBeNil)
		So(s, ShouldHaveSameTypeAs, &Sift{})
	})

	Convey("Removes Label from customer", t, func() {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %v", r.URL)
			w.WriteHeader(http.StatusNoContent)
		}))

		defer svr.Close()
		defer s.SetApiUrl(API_URL)

		s.SetApiUrl(svr.URL)

		r, err := s.UnLabel("test@someone.com")

		So(r, ShouldHaveSameTypeAs, &Response{})
		So(err, ShouldBeNil)

		So(r.IsOK(), ShouldBeTrue)

		So(r.HTTPStatusCode, ShouldEqual, 204)
		So(r.HTTPStatus, ShouldEqual, "204 No Content")
		So(r.ErrorMessage, ShouldEqual, "")
	})
}

func TestSiftClientErrorRequest(t *testing.T) {
	s, err := NewFromJSON([]byte(testValidApiConfig))

	Convey("NewFromJSON() returns pointer to Sift", t, func() {
		So(err, ShouldBeNil)
		So(s, ShouldHaveSameTypeAs, &Sift{})
	})

	Convey("Should return unauthorized error", t, func() {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %v", r.URL)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(testApiError))
		}))

		defer svr.Close()
		defer s.SetApiUrl(API_URL)

		s.SetApiUrl(svr.URL)

		r, err := s.UnLabel("test@someone.com")

		So(r, ShouldHaveSameTypeAs, &Response{})
		So(err, ShouldNotBeNil)

		So(r.IsOK(), ShouldBeFalse)

		So(r.HTTPStatusCode, ShouldEqual, 400)
		So(r.Status, ShouldEqual, 51)
		So(r.HTTPStatus, ShouldEqual, "400 Bad Request")
		So(r.ErrorMessage, ShouldEqual, ErrorCodes[51])
	})
}

// -----------------------------------------------------------------------------

func BenchmarkSiftClientTrackRequest(b *testing.B) {
	s, _ := NewFromJSON([]byte(testValidApiConfig))

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testEventOK)
	}))

	defer svr.Close()
	defer s.SetApiUrl(API_URL)

	s.SetApiUrl(svr.URL)

	for n := 0; n < b.N; n++ {
		s.Track(
			"my_custom_event",
			map[string]interface{}{
				"$user_id": "test@someone.com", "hello": "world",
			}, map[string]interface{}{
				"return_score":  true,
				"return_action": true,
			},
		)
	}
}
