package parser

import (
	"io/ioutil"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParseCityList(t *testing.T) {
	Convey("ParseCityList given a html file", t, func() {
		content, err := ioutil.ReadFile("./snapshot/citylist_test.html")
		if err != nil {
			panic(err)
		}
		result := ParseCityList(content)
		const resultSize = 70
		Convey("result.requests length should equal with the given num", t, func() {
			So(len(result.Requests), ShouldEqual, resultSize)
		})

		Convey("result.items length should equal with the given num", t, func() {
			So(len(result.Items), ShouldEqual, resultSize)
		})
	})
}
