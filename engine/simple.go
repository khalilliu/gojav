package engine

import (
	"fmt"
	"gojav/model"
	"gojav/persist"
)

var (
	startUrl       string
	TargetHasFound = false
	CurPage        = 1
	end            = false
	total          = 0
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
			persist.SaveItem(item.(model.Movie))
		}
	}


	//TargetHasFound = true
}


