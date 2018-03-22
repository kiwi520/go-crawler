package parser

import (
	"regexp"
	"crawler/engine"
	"fmt"
)

const cityListRe =`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	//limit := 10
	for _,v := range matches {
		fmt.Println(string(v[1]))
		result.Items = append(result.Items, string(v[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(v[1]),
			ParseFunc: ParseCity,
		})
		//limit--
		//if limit ==0{
		//	break
		//}
	}
	return result
}