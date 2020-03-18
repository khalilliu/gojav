package engine

import (
	"fmt"
	"gojav/config"
	"log"
	"strconv"
)

var (
	startUrl       string
	TargetHasFound = false
	curPage        = 1
	end            = false
)

type SimpleEngine struct {
}

func (e *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			fmt.Printf("msg: 获取%s失败, err: %+v", r.Url, err)
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
	//TargetHasFound = true
}

func GetStartUrl() string {
	startUrl = config.BaseUrl
	if config.Cfg.Search != "" {
		startUrl = fmt.Sprintf("%s%s/%s", config.BaseUrl, config.SearchRoute, config.Cfg.Search)
	}
	if curPage != 1 {
		startUrl += "/page/%d" + strconv.Itoa(curPage)
	}
	return startUrl
}
