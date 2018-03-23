package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	//queue版不需要此方法
	//ConfigureMasterWorkerChan(chan Request)
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}


func (e *ConcurrentEngine) Run(sends ...Request){
	//queue版
	out := make(chan ParseResult)
	e.Scheduler.Run()
    //并发coroutine版
	//in := make(chan Request)
	//out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)

	for i:= 0; i<e.WorkerCount ;i++  {
		//queue并发版
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)

		//并发coroutine版
		//createWorker(in, out)
	}

	for _,r := range sends{
		e.Scheduler.Submit(r)
	}

	itemCount:= 0
	for {
		result:= <-out
		for _,item := range  result.Items {
			log.Printf("Got item #%d: %v",itemCount,item)
			itemCount ++
		}

		for _, request := range  result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

//queue并发版
func createWorker(in chan Request,out chan ParseResult, ready ReadyNotifier)  {
	//in := make(chan Request)
	go func() {
		for {
			ready.WorkerReady(in)
			request := <- in
			result,err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

//并发coroutine版
//func createWorker(in chan  Request,out chan ParseResult)  {
//	go func() {
//		for {
//			request := <- in
//			result,err := worker(request)
//			if err != nil {
//				continue
//			}
//			out <- result
//		}
//	}()
//}