package engine

import (
	"fmt"
	"github.com/fatih/color"
)

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ConcurrentEngine struct {
	Scheduler   Scheduler //调度器
	WorkerCount int       // worker的数量
	ItemChan    chan Item // 接收最终结果的chan
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	fmt.Println(color.CyanString("Starting..."))

	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				// 处理Item
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}

		// check if engine is finished

	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
