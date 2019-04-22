package parser

import (
	"fmt"
	"learngo/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	metches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range metches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
		fmt.Printf("City: %s,URL: %s\n", m[2], m[1])
	}
	return result

}
