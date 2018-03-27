package persist

import (
	"testing"
	"crawler/models"
	"github.com/olivere/elastic"
	"context"
	//"encoding/json"
	"fmt"
	"encoding/json"
)

func TestServe(t *testing.T) {
	profile := models.Profile{
		Gender:"女",
		Age:24,
		Height:166,
		Weight:0,
		Marriage:"未婚",
		Education:"大学本科",
		Income:"3001-5000元",
		Occupation:"小学教师",
		Xinzuo:"狮子座",
		Hukou:"山东济南",
		House:"--",
		Car:"未购车",
	}
	elasticId,err :=save(profile)

	if err != nil{
		panic(err.Error())
	}

	client ,err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil{
		panic(err.Error())
	}

	result,err:=client.Get().Index("profile").Type("zhenai").Id(elasticId).Do(context.Background())

	if err != nil{
		panic(err.Error())
	}

	var act  models.Profile
	err =json.Unmarshal([]byte(*result.Source),&act)
	if err != nil{
		panic(err.Error())
	}

	fmt.Println(act)
}