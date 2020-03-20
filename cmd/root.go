package cmd

import (
	"fmt"
	"gojav/config"
	"gojav/engine"
	"gojav/parser"
	"gojav/utils"
	"time"
)

var (
	startUrl       = config.BaseUrl
)

func Execute() {
	// 程序主入口
	e := &engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:       utils.GetUrl(""),
		ParseFunc: parser.ParseMovieList,
	})
	for {
		time.Sleep(time.Second)
		if engine.TargetHasFound {
			fmt.Println("抓取完成")
			break
		}
	}
}


