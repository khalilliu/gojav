package parser

import (
	"gojav/engine"
	"regexp"
)

const (
	cityReg    = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
	cityUrlReg = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)">`
)

func ParseCity(contents []byte) engine.ParseResult {
	cityRe := regexp.MustCompile(cityReg)
	cityUrlReg := regexp.MustCompile(cityUrlReg)

	all := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, c := range all {
		url := string(c[1])
		name := string(c[2])
		//fmt.Println("用户url：", url)
		//result.Items = append(result.Items, "User:"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParseFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, url, name)
			},
		})
	}

	nextPage := cityUrlReg.FindAllSubmatch(contents, -1)
	for _, c := range nextPage {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(c[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
