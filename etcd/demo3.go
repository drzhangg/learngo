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

	//获取租约
	lease := clientv3.NewLease(client)
	leaseGrant, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println(err)
	}

	kv := clientv3.NewKV(client)
	_, err = kv.Put(context.TODO(), "/test/key3", "租约请求", clientv3.WithLease(leaseGrant.ID))
	if err != nil {
		fmt.Println(err)
	}

	leaseKeepResp, err := lease.KeepAlive(context.TODO(), leaseGrant.ID)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			select {
			case leaseKeepResp := <-leaseKeepResp:
				if leaseKeepResp == nil {
					fmt.Println("租约停止")
					break
				} else {
					fmt.Println("收到租约应答：", leaseKeepResp.ID)
				}
			}
		}
	}()

	go func() {
		watchChan := client.Watch(context.TODO(), "/test/key3", clientv3.WithPrevKV())
		for watch := range watchChan {
			for _, e := range watch.Events {
				fmt.Printf("type:%v kv:%v  prevKey:%v \n ", e.Type, string(e.Kv.Key), e.PrevKv)
			}
		}
	}()

}
