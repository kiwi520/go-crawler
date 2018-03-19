package main

import (
	"crawler/engine"
	"crawler/date/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parser.ParserCityList,
	})
	//engine.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun/jinan",
	//	ParseFunc:parser.ParseCity,
	//})
}



