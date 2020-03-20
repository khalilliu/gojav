package engine

import (
	"fmt"
)

type SimpleEngine struct {
	Scheduler   Scheduler //调度器
	WorkerCount int       // worker的数量
	ItemChan    chan Item // 接收最终结果的chan
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
			e.ItemChan <- item
		}
	}


	//TargetHasFound = true
}


