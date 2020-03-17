package parser

import (
	"github.com/bitly/go-simplejson"
	"gojav/engine"
	"log"
	"regexp"
)

var (
	cityLimit = 0
 	cityListReg= `<script>window.__INITIAL_STATE__=(.+);\(function`
)

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListReg)
	match := re.FindSubmatch(contents)

	all := parseJsonCityList(match[1])
	result := engine.ParseResult{}
	i := 0
	for _, c := range all {
		//result.Items = append(result.Items, string(c[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       c[1],
			ParseFunc: ParseCity,
		})
		i++
		if cityLimit != 0 && i >= cityLimit {
			break
		}
	}
	return  result
}


func parseJsonCityList(cityListStr []byte) (cityList [][]string) {
	resp, err := simplejson.NewJson(cityListStr)
	if err != nil {
		log.Println("解析json失败...")
		return
	}
	infos, _ := resp.Get("cityListData").Get("cityData").Array()

	for _,v := range infos {
		if each_map, ok := v.(map[string]interface{}); ok {
			map2 := each_map["cityList"]
			for _, v2 := range map2.([]interface{}) {
				if data, ok := v2.(map[string]interface{}); ok {
					var dataPair [] string
					cityName := data["linkContent"].(string)
					cityUrl := data["linkURL"].(string)
					dataPair = append(dataPair, cityName, cityUrl)
					cityList = append(cityList, dataPair)
				}
			}
		}
	}
	return
}