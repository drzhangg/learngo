package main

import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

type Service struct {
	closeChan chan struct{}
	client    *clientv3.Client
	leaseID   clientv3.LeaseID
	key       string
	val       string
	wg        sync.WaitGroup
}

// NewService 构造一个注册服务函数
func NewService(addrs []string, key, val string) (*Service, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	service := &Service{
		closeChan: make(chan struct{}),
		client:    client,
		key:       key,
		val:       val,
	}
	return service, nil
}

// 开启注册
func (s *Service) Start(ttl int64) error {
	//申请续租时间
	leaseGrantResp, err := s.client.Grant(context.TODO(), ttl)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//获得租约id
	s.leaseID = leaseGrantResp.ID

	_, err = s.client.Put(context.TODO(), s.key, s.val, clientv3.WithLease(s.leaseID))
	if err != nil {
		fmt.Println(err)
		return err
	}

	ch, err := s.client.KeepAlive(context.TODO(), s.leaseID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("service start leaseID:%d,key:%s,val:%s", s.leaseID, s.key, s.val)
	s.wg.Add(1)
	defer s.wg.Done()

	for {
		select {
		case <-s.closeChan:
			return s.revoke()
		case <-s.client.Ctx().Done():
			return errors.New("service closed")
		case ka, ok := <-ch:
			if !ok {
				return s.revoke()
			} else {
				fmt.Printf("service start recv reply from service:%s,ttl:%d", s.key, ka.TTL)
			}
		}
	}
}

func (s *Service) revoke() error {
	return nil
}

func main() {

}
