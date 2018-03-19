package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult  {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,v := range matches {
		name := string(v[2])
		result.Items = append(result.Items,"User"+ name)
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(v[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c,name)
			},
		})
	}
	return result
}