package parser

import (
	"learngo/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	
	result := ParseCityList(contents)

	const resultSize  = 470
	expectedUrls := []string{"http://www.zhenai.com/zhenghun/aba","http://www.zhenai.com/zhenghun/akesu","http://www.zhenai.com/zhenghun/alashanmeng"}
	expectedCities := []string{"阿坝","阿克苏","阿拉善盟"}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests;but had %d",resultSize,len(result.Items) )
	}
	for i,url := range expectedUrls{
		if result.Requests[i].Url != url {
			t.Errorf("expected url %d: %s; but was %s",i,url,result.Requests[i].Url)
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests;but had %d",resultSize,len(result.Requests))
	}
	for i,city := range expectedCities{
		if result.Items[i].(string) != city {
			t.Errorf("expected url %d: %s; but was %s",i,city,result.Items[i].(string))
		}
	}
}
