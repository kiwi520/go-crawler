package parser

import (
"testing"
"crawler/fetcher"

	"fmt"
)

func TestParserCity(t *testing.T) {
	contents ,err := fetcher.Fetch("./citylist_test_data.html")

	if err != nil {
		panic(err.Error())
	}

	result :=ParseCity(contents)

	fmt.Printf("%s\n",result)
	//const resultSize  = 470
	//expectedUrls :=[]string{
	//	"http://www.zhenai.com/zhenghun/aba",
	//	"http://www.zhenai.com/zhenghun/akesu",
	//	"http://www.zhenai.com/zhenghun/alashanmeng",
	//}
	//expectedCities :=[]string{
	//	"City 阿坝","City 阿克苏","City 阿拉善盟",
	//}
	//
	//if len(result.Requests) != resultSize {
	//	t.Errorf("测试结果应为 %d resquests; 但是只获取的 %d ",resultSize,len(result.Requests))
	//}
	//
	//for i,url :=range expectedUrls{
	//	if result.Requests[i].Url !=url {
	//		t.Errorf("expected url #%d: %s; 而不是 %s",i,url,result.Requests[i].Url)
	//	}
	//}
	//
	//if len(result.Items) != resultSize {
	//	t.Errorf("测试结果应为 %d resquests; 但是只获取的 %d ",resultSize,len(result.Items))
	//}
	//
	//for i,city :=range expectedCities{
	//	if result.Items[i].(string) !=city {
	//		t.Errorf("expected url #%d: %s; 而不是 %s",i,city,result.Items[i].(string))
	//	}
	//}
}
