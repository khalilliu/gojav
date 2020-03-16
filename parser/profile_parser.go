package parser

import (
	"github.com/bitly/go-simplejson"
	"gojav/engine"
	"gojav/model"
	"log"
	"regexp"
)

const (
	JsonDataReg = `<script>window.__INITIAL_STATE__=(.+);\(function`
)

func ParseProfile(content []byte, name string) engine.ParseResult {
	re := regexp.MustCompile(JsonDataReg)
	match := re.FindSubmatch(content)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		data := match[1]
		user := parseUser(data)
		user.Name = name
		//fmt.Println(user)
		result.Items = append(result.Items, user)
	}
	return result
}

func parseUser(json []byte) model.User {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}
	infos, err := res.Get("objectInfo").Get("basicInfo").Array()

	var user model.User
	for k, v := range infos {
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				user.Marriage = e
			case 1:
				//年龄:47岁，我们可以设置int类型，所以可以通过另一个json字段来获取
				user.Age = e
			case 2:
				user.Xingzuo = e
			case 3:
				user.Height = e
			case 4:
				user.Weight = e
			case 6:
				user.Income = e
			case 7:
				user.Occupation = e
			case 8:
				user.Education = e
			}
		}
	}
	return user
}
