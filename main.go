package main

import (
	"gojav/engine"
	"gojav/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url: url,
		ParseFunc: parser.ParseCityList,
	})
}