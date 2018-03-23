package main

import (
	"crawler/engine"
	"crawler/date/parser"
	"crawler/scheduler"
)

func main() {
	//queue并发版
	e:= engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:100,
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc:parser.ParseCity,
	})
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFunc:parser.ParserCityList,
	//})

	//并发coroutine版
	//e:= engine.ConcurrentEngine{
	//	Scheduler:&scheduler.SimpleScheduler{},
	//	WorkerCount:100,
	//}
	//
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFunc:parser.ParserCityList,
	//})

	//单任务简单版
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFunc:parser.ParserCityList,
	//})
	//engine.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun/jinan",
	//	ParseFunc:parser.ParseCity,
	//})
}



