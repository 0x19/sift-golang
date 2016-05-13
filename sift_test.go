//
//
package sift

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	testApiKey         = "A"
	testValidApiConfig = fmt.Sprintf(`{"config": {"api_key": "%s"}}`, testApiKey)
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

func TestSiftClientRequest(t *testing.T) {
	s, err := NewFromJSON([]byte(testValidApiConfig))

	Convey("NewFromJSON() returns pointer to Sift", t, func() {
		So(err, ShouldBeNil)
		So(s, ShouldHaveSameTypeAs, &Sift{})
	})

	Convey("NewFromJSON() returns pointer to Sift", t, func() {
		r, err := s.Track(
			"my_custom_event",
			map[string]interface{}{
				"$user_id": "nevio@someone.com", "hello": "world",
			}, map[string]interface{}{
				"return_score":  true,
				"return_action": true,
			},
		)

		So(r, ShouldHaveSameTypeAs, &Response{})
		So(err, ShouldBeNil)
	})

}
