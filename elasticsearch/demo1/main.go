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

func main() {

}
