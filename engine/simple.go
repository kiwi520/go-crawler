package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(sends ...Request)  {
	var requests []Request

	for _,r :=range sends{
		requests = append(requests,r)
	}

	for len(requests) >0  {
		r := requests[0]
		requests = requests[1:]

		//log.Printf("Fetching %s",r.Url)
		//body,err := fetcher.Fetch(r.Url)
		//
		//if err != nil{
		//	log.Printf("Fetcher :err Fetch")
		//	continue
		//}
		//
		//ParseResult := r.ParseFunc(body)

		ParseResult,err := worker(r)
		if err != nil{
			continue
		}
		requests = append(requests,ParseResult.Requests...)

		for _,item := range ParseResult.Items {
			log.Println(item)

		}
	}
}

func worker(r Request)(ParseResult,error){
	log.Printf("Fetching %s",r.Url)
	body,err := fetcher.Fetch(r.Url)

	if err != nil{
		log.Printf("Fetcher :err Fetch")
		return  ParseResult{},err
	}

	return r.ParseFunc(body),nil
}