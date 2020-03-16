package parser

import (
	"gojav/engine"
	"regexp"
)

const (
	cityLimit = 20
 	cityListReg= `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`
)

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListReg)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	i := 0
	for _, c := range all {
		result.Items = append(result.Items, string(c[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(c[1]),
			ParseFunc: ParseCity,
		})
		i++
		if i >= cityLimit {
			break
		}
	}
	return  result
}
