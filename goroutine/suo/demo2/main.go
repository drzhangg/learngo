package main

import (
	"fmt"
	"sync"
)

var(
	//逻辑中使用的某个变量
	count int

	//与变量对应的使用互斥锁
	countGuard sync.RWMutex
)

func GetCount() int {
	//锁定
	countGuard.RLock()

	//在函数退出时解除锁定
	defer countGuard.RUnlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func main() {
	SetCount(1)

	fmt.Println(GetCount())
}
