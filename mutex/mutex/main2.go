package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m1 = make(map[int]uint64)
	lock1 sync.Mutex
)

type task1 struct {
	n int
}

func calc1(t *task1) {
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}

	fmt.Println(t.n,sum)
	lock1.Lock()
	m1[t.n] = sum
	lock1.Unlock()
}

func main() {
	for i := 0; i < 16; i++ {
		t1:=&task1{n:i}
		go calc1(t1)
	}

	time.Sleep(10 * time.Second)
	lock1.Lock()
	for k,v := range m1{
		fmt.Printf("%d != %v\n",k,v)
	}
	lock1.Unlock()
}
