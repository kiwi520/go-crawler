package persist

import (
	"log"
	"github.com/olivere/elastic"
	"context"
)

func ItemServer() chan  interface{} {
	out := make(chan interface{})

	go func() {
		itemCount :=0
		for {
			item:= <-out
			log.Printf("Item Server: get item #%d: %v",itemCount,item)
			itemCount++

			_,err :=save(item)
			if err != nil{
				log.Printf("Item Saver :error saving item %v: %v",item,err)
			}
		}
	}()

	return  out
}

func save(item interface{}) (id string,err error){
	client,err := elastic.NewClient(elastic.SetSniff(false))
	if err !=nil{
		return "",err
	}

	resp,err := client.Index().Index("profile").Type("zhenai").BodyJson(item).Do(context.Background())

	if err != nil{
		return "",err
	}
	return resp.Id,nil
}