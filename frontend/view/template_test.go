package view

import (
	. "github.com/smartystreets/goconvey/convey"
	"gojav/engine"
	viewModel "gojav/frontend/model"
	"gojav/model"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	Convey("template render successfully", t, func() {
		view := CreateSearchResultView("template.html")
		out, err := os.Create("template.test.html")

		page := viewModel.SearchResult{}
		page.Hits = 123

		item := engine.Item{
			Url:  "http://album.zhenai.com/u/1214814888",
			Type: "zhenai",
			Id:   "1214814888",
			Payload: model.Profile{
				Name:       "林YY",
				Marriage:   "未婚",
				Age:        "26岁",
				Xingzuo:    "魔羯座(12.22-01.19)",
				Height:     "165cm",
				Weight:     "50kg",
				Income:     "月收入:5-8千",
				Occupation: "职业技术教师",
				Education:  "高中及以下",
			},
		}

		page.Items = append(page.Items, item)
		err = view.Render(out, page)
		So(err, ShouldBeNil)
	})

}