package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	. "github.com/smartystreets/goconvey/convey"
	"gojav/model"
	"testing"
)

func TestSave(t *testing.T) {

	Convey("save item return id should be right", t, func() {
		profile := model.User{
			Name:       "林YY",
			Marriage:   "未婚",
			Age:        "26岁",
			Xingzuo:    "魔羯座(12.22-01.19)",
			Height:     "165cm",
			Weight:     "50kg",
			Income:     "月收入:5-8千",
			Occupation: "职业技术教师",
			Education:  "高中及以下",
		}
		id, err := save(profile)

		So(err, ShouldBeNil)

		client, _ := elastic.NewClient(elastic.SetSniff(false))
		resp, err := client.Get().Index("datint_profile").Type("zhenai").Id(id).Do(context.Background())

		var actual model.User
		err = json.Unmarshal([]byte(*resp.Source), &actual)
		if err != nil {
			panic(err)
		}

		So(profile.Name, ShouldEqual, actual.Name)
	})
}
