package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"strconv"
	"time"
)

var host = []string{
	"http://127.0.0.1:9200",
}

var client *elastic.Client

func init() {
	var err error
	if client, err = elastic.NewClient(elastic.SetURL(host...)); err != nil {
		fmt.Printf("create client failed, err: %v", err)
	}
}

func PingNode() {
	start := time.Now()

	info, code, err := client.Ping(host[0]).Do(context.TODO())
	if err != nil {
		fmt.Printf("ping es failed, err: %v", err)
	}

	duration := time.Since(start)
	fmt.Printf("cost time: %v\n", duration)
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

//校验index是否存在
func IndexExists(index ...string) bool {
	exists, err := client.IndexExists(index...).Do(context.TODO())
	if err != nil {
		fmt.Sprintf("%v\n", err)
	}
	return exists
}

func CreateIndex(index, mapping string) bool {

	result, err := client.CreateIndex(index).BodyString(mapping).Do(context.TODO())
	if err != nil {
		fmt.Printf("create index failed, err: %v\n", err)
	}
	return result.Acknowledged
}

func DelIndex(index ...string) bool {
	response, err := client.DeleteIndex(index...).Do(context.TODO())
	if err != nil {
		fmt.Printf("delete index failed, err: %v\n", err)
	}
	return response.Acknowledged
}

func Batch(index string, type_ string, datas ...interface{}) {
	bulkRequest := client.Bulk()
	for i, data := range datas {
		doc := elastic.NewBulkIndexRequest().Index(index).Type(type_).Id(strconv.Itoa(i)).Doc(data)
		bulkRequest = bulkRequest.Add(doc)
	}

	response, err := bulkRequest.Do(context.TODO())
	if err != nil {
		panic(err)
	}
	failed := response.Failed()
	iter := len(failed)
	fmt.Printf("error: %v, %v\n", response.Errors, iter)

}


//获取指定id的文档
func GetDoc(index, id string) []byte {
	temp := client.Get().Index(index).Id(id)
	get, err := temp.Do(context.TODO())
	if err != nil {
		panic(err)
	}

	if get.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get.Id, get.Version, get.Index, get.Type)
	}

	source,err := get.Source.MarshalJSON()
	if err != nil {
		fmt.Printf("byte convert string failed, err: %v", err)
	}
	return source
}

func TermQuery(index, type_, fileName, fileValue string) *elastic.SearchResult {
	query := elastic.NewTermQuery(fileName,fileValue)
	searchResult, err := client.Search().Index(index).Type(type_).Query(query).From(0).Size(10).Pretty(true).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("query cost %d millisecond.\n", searchResult.TookInMillis)

	return searchResult
}



func Search(index, type_ string) *elastic.SearchResult {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewMatchQuery("user", "Jame10"))
	boolQuery.Filter(elastic.NewRangeQuery("age").Gt("30"))
	searchResult, err := client.Search(index).
		Type(type_).Query(boolQuery).Pretty(true).Do(context.Background())
	if err != nil {
		panic(err)
	}

	return searchResult
}

func AggsSearch(index, type_ string) {

	minAgg := elastic.NewMinAggregation().Field("age")
	rangeAgg := elastic.NewRangeAggregation().Field("age").AddRange(0,30).AddRange(30,60).Gt(60)


	build := client.Search(index).Type(type_).Pretty(true)

	minResult, err := build.Aggregation("minAgg", minAgg).Do(context.Background())
	rangeResult, err := build.Aggregation("rangeAgg", rangeAgg).Do(context.Background())
	if err != nil {
		panic(err)
	}

	minAggRes, _ := minResult.Aggregations.Min("minAgg")
	fmt.Printf("min: %v\n", *minAggRes.Value)

	rangeAggRes, _ := rangeResult.Aggregations.Range("rangeAgg")
	for _, item := range rangeAggRes.Buckets {
		fmt.Printf("key: %s, value: %v\n", item.Key, item.DocCount)
	}

}


func main() {

}
