package parser

import (
	. "github.com/smartystreets/goconvey/convey"
	"gojav/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	Convey("ParseCityList given a html file", t, func() {
		content, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
		if err != nil {
			panic(err)
		}
		result := ParseCityList(content)
		resultSize := 470
		Convey("result requests length should equal with the given num",  func() {
			So(len(result.Requests), ShouldEqual, resultSize)
		})
	})
}
