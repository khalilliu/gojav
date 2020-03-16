package main

import (
	"gojav/engine"
	"gojav/parser"
	"gojav/persist"
	"gojav/scheduler"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: url,
	//	ParseFunc: parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 20,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:       url,
		ParseFunc: parser.ParseCityList,
	})
}
