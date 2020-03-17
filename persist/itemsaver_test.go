package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	. "github.com/smartystreets/goconvey/convey"
	"gojav/engine"
	"gojav/model"
	"testing"
)

func TestSave(t *testing.T) {

	Convey("save item return id should be right", t, func() {
		profile := engine.Item{
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
		err := save(profile)

		So(err, ShouldBeNil)

		client, _ := elastic.NewClient(elastic.SetSniff(false))
		resp, err := client.Get().Index("datint_profile").Type(profile.Type).Id(profile.Id).Do(context.Background())

		var actual engine.Item
		err = json.Unmarshal([]byte(*resp.Source), &actual)
		if err != nil {
			panic(err)
		}
		fmt.Println(actual)
		So(actual.Url, ShouldEqual, profile.Url)
	})
}
