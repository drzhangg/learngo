package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	conf := clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}, DialTimeout: 5 * time.Second}

	client, err := clientv3.New(conf)
	if err != nil {
		fmt.Println(err)
	}

	//创建租约
	lease := clientv3.NewLease(client)

	//获取租约id
	leaseGrant, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println(err)
	}

	//获取kv
	kv := clientv3.NewKV(client)

	//续租
	leaseKeepResp, err := lease.KeepAlive(context.TODO(), leaseGrant.ID)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		select {
		case leaseKeepChan := <-leaseKeepResp:
			if leaseKeepChan == nil {
				fmt.Println("租约失效")
				break
			} else {
				fmt.Println("收到续租请求：", leaseKeepChan.ID)
			}
		}
	}()

	//put操作
	putResp, err := kv.Put(context.TODO(), "/test/key4", "hello", clientv3.WithLease(leaseGrant.ID))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("写入成功：", putResp.Header.Revision)

	//watch操作

	//get操作
	for {
		getResp, err := kv.Get(context.TODO(), "/test/key4")
		if err != nil {
			fmt.Println(err)
		}

		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}

		fmt.Println("kv还没过期：", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}
