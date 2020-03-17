package engine

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
	Scheduler   Scheduler // 调度器
	WorkerCount int       // worker的数量
	ItemChan    chan Item // 接收最终结果的chan
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	co.Println("starting...")
	//in := make(chan Request)
	out := make(chan ParseResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//从out中获取result，传入ItemChan，对于request，就继续分配
	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
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
