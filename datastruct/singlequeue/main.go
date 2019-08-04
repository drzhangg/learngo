package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxsize int //队列最大长度
	array   [5]int
	front   int //队列首部
	rear    int //队列尾部
}

func (this *Queue) AddQueue(val int) (error) {

	//先判断队列是否已满
	if this.rear == this.maxsize-1 {
		return errors.New("queue full")
	}

	//如果没满
	this.rear++
	//如果没满，当前队尾长度+1，并把数据放到当前位置
	this.array[this.rear] = val

	return nil
}

func (this *Queue) GetQueue() (val int, err error) {

	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}

	this.front++
	val = this.array[this.front]
	return val, nil
}

func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")

	//this.front 不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t\n", i, this.array[i])
	}
}

func main() {

	queue := &Queue{
		maxsize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 表示退出队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数：")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {

				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出一个数=", val)
			}

		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}

	}
}
