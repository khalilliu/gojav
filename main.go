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

	itemChan, err := persist.ItemSaver("datint_profile")

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 5,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{
		Url:       url,
		ParseFunc: parser.ParseCityList,
	})
}
