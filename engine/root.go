package engine

import (
	"fmt"
	"gojav/config"
	"gojav/parser"
	"strconv"
	"time"
)

var (
	TargetHasFound = false
	startUrl string
	curPage     = 1
	end         = false

)


func Execute() {
	// 程序主入口
	e := &SimpleEngine{
	}
	e.Run(Request{
		Url: getStartUrl(),
		ParseFunc: parser.ParseMovieList,
	})
	for {
		time.Sleep(time.Second)
		if TargetHasFound {
			fmt.Println("抓取完成")
			return
		}
	}
}


func getStartUrl() string {
	if config.Cfg.Search != "" {
		startUrl = fmt.Sprintf("%s%s/%s", config.BaseUrl, config.SearchRoute, config.Cfg.Search)
	}
	if curPage != 1 {
		startUrl += "/page/%d" + strconv.Itoa(curPage)
	}
	return startUrl
}

