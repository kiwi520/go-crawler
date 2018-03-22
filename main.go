package main

import (
	"crawler/engine"
	"crawler/date/parser"
	"crawler/scheduler"
)

func main() {
	e:= engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:100,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parser.ParserCityList,
	})
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFunc:parser.ParserCityList,
	//})
	//engine.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun/jinan",
	//	ParseFunc:parser.ParseCity,
	//})
}



