package parser

import (
	"crawler/engine"
	"regexp"
	//"fmt"
)

var cityRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

var cityPageRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[0-9a-z]+/[0-9]+)">下一页</a>`)

func ParseCity(contents []byte) engine.ParseResult  {

	matches := cityRe.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,v := range matches {
		name := string(v[2])
		//result.Items = append(result.Items,"User"+ name)
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(v[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c,name)
			},
		})
	}

	matchess := cityPageRe.FindSubmatch(contents)

	for k,v := range matchess{
		if k != 0{
			result.Requests = append(result.Requests,engine.Request{
				Url:string(v),
				ParseFunc:ParseCity,
			})
		}
	}
	return result
}
