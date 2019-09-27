package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	config := clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}, DialTimeout: 5 * time.Second}

	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建租约
	lease := clientv3.NewLease(client)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	leaseGrant, err := lease.Grant(ctx, 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	//租约id
	leaseId := leaseGrant.ID

	//续租
	leaseKeepResp, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		fmt.Println(err)
		return
	}

	//开启协程来监听
	go listenLeaseChan(leaseKeepResp)

	//获取处理业务kv
	kv := clientv3.NewKV(client)

	// 重点：事务操作
	// 开启事务
	txn := kv.Txn(ctx)
	txn.If(clientv3.Compare(clientv3.CreateRevision("/test/key4"), "=", 0)).Then(clientv3.OpPut("/test/key4", "xxx", clientv3.WithLease(leaseId))).Else(clientv3.OpGet("/test/key4")) //否则抢锁失败

	//提交事务
	txnResp, err := txn.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}

	//判断是否抢到锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用：", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}
	fmt.Println("处理任务")

	//终止租约
	defer lease.Revoke(ctx, leaseId)

	time.Sleep(10 * time.Second)
}

func listenLeaseChan(leaseRespChan <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case leaseKeepResp := <-leaseRespChan:
			if leaseKeepResp == nil {
				fmt.Println("租约失效了")
				goto END
			} else {
				fmt.Println("租约id为：", leaseKeepResp.ID)
			}
		}
	}
END:
}
