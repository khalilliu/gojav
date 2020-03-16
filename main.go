package main

import (
	"gojav/engine"
	"gojav/parser"
	"gojav/scheduler"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: url,
	//	ParseFunc: parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:       url,
		ParseFunc: parser.ParseCityList,
	})
}
