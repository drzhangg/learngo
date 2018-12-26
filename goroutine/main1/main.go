package main

import (
	"fmt"
	"sync"
	"time"
)

var(
	m = make(map[int]uint64)
	lock sync.Mutex		//申明一个互斥锁
)

type task struct {
	n int
}

func calc(t *task) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error...")
			return
		}
	}()

	var sum uint64 = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}

	lock.Lock()		//写全局数据加互斥锁
	m[t.n] = sum
	lock.Unlock()	//解锁

}

func main() {
	for i := 0; i < 10; i++ {
		t := &task{n:i}
		go calc(t)	//goroutine来执行任务
	}

	time.Sleep(time.Second)	//goroutine异步，所以等一秒到任务完成

	lock.Lock()
	for k,v := range m{
		fmt.Printf("%d != %v\n",k,v)
	}
	fmt.Println(len(m))
	lock.Unlock()
}
