package main

import (
	"container/list"
	"fmt"
)

func main() {
	//创建一个链表
	alist := list.New()
	fmt.Println("Size before:",alist.Len())

	//向链表中添加元素
	alist.PushBack("a")
	alist.PushBack("b")
	alist.PushFront("c")
	fmt.Println("Size after insert(push):",alist.Len())

	//遍历
	for e := alist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(string))
	}

	//移除元素
	alist.Remove(alist.Front())
	fmt.Println("Size after remove(pop):",alist.Len())
}
