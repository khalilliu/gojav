package cmd

import (
	"gojav/config"
	"gojav/engine"
	"gojav/parser"
	"gojav/persist"
	"gojav/scheduler"
	"gojav/utils"
)

var (
	startUrl       = config.BaseUrl
)

func Execute() {
	// 程序主入口

	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}

	e := &engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: config.Cfg.Parallel,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{
		Url:       utils.GetUrl(""),
		Type: engine.HTML,
		ParseFunc: parser.ParseMovieList,
	})

}


