package parser

import (
	"testing"
	"crawler/fetcher"
	"crawler/models"
)

func TestParseProfile(t *testing.T) {
	contents ,err := fetcher.Fetch("http://album.zhenai.com/u/1254739352")
	//
	//contents ,err := ioutil.ReadFile("./profile_test_data.html")

	if err!= nil{
		panic(error.Error)
	}

	result := ParseProfile(contents)

	if len(result.Items) !=1{
		t.Errorf("items 条数应为 1 但是它是 %v",len(result.Items))
	}

	profile := result.Items[0].(models.Profile)

	expected := models.Profile{
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
	if profile != expected{
		t.Errorf("expected %v; but was %v",expected,profile)
	}
}