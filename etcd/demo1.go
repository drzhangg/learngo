package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	conf := clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}, DialTimeout: 5 * time.Second,}

	client, err := clientv3.New(conf)
	if err != nil {
		fmt.Println(err)
	}

	//lease := clientv3.NewLease(client)

	//kv := clientv3.NewKV(client)
	//leaseResp, err := lease.Grant(context.TODO(), 2)
	//if err != nil {
	//	fmt.Println(err)
	//}

	putResp, err := client.Put(context.TODO(), "/test/key1", "测试数据")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(putResp)

	//fmt.Println(putResp.PrevKv.Key)
	//fmt.Println(putResp.PrevKv.Value)

	getResp, err := client.Get(context.TODO(), "/test/key1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getResp)
	for _, v := range getResp.Kvs {
		fmt.Println(string(v.Key))
		fmt.Println(string(v.Value))
		fmt.Println(v.Lease)
		fmt.Println(v.Version)
	}

	//client.Get(context.TODO(),putResp.PrevKv.Key)
}
