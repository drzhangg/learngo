package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//etcd lease
func main() {
	conf := clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}, DialTimeout: 5 * time.Second}

	client, err := clientv3.New(conf)

	//lease租约
	lease := clientv3.NewLease(client)
	//设置租约过期时间为3秒
	leaseGrant, err := lease.Grant(context.TODO(), 6)
	if err != nil {
		fmt.Println(err)
	}
	kv := clientv3.NewKV(client)
	//获取要续约的id
	leaseId := leaseGrant.ID

	//clientv3.LeaseKeepAliveResponse{}
	_, err = kv.Put(context.TODO(), "/test/key2", "测试租约", clientv3.WithLease(leaseId))
	if err != nil {
		fmt.Println(err)
	}

	ctx, _ := context.WithTimeout(context.TODO(), time.Duration(time.Second*5))
	//续租
	leaseKeepChan, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(leaseKeepChan)

}
