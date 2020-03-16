package scheduler

import "gojav/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func(){
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler)  Run() {
	s.workerChan = make(chan engine.Request)
}