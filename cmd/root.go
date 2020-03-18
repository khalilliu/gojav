package cmd

import (
	"gojav/config"
	"gojav/engine"
	"gojav/parser"
)

var (
	startUrl       = config.BaseUrl
)

func Execute() {
	// 程序主入口
	e := &engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:       engine.GetStartUrl(),
		ParseFunc: parser.ParseMovieList,
	})
	//for {
	//	time.Sleep(time.Second)
	//	if engine.TargetHasFound {
	//		fmt.Println("抓取完成")
	//		return
	//	}
	//}
}


