package main

import "fmt"

//获取指定个的数字并写入到numChan里
func getNum(numChan chan int, num int) {
	for i := 0; i < num; i++ {
		numChan <- i
	}
	//全部写完后关闭通道
	close(numChan)
}

func resNum(numChan chan int, resChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		flag = true

		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		//flag为true说明这时候读到的是素数
		if flag {
			//这时候写入数据，但是不关闭channel，因为可能当一个协程读完毕后，其他协程还在读，所以不能关闭
			resChan <- num
		}
	}
	exitChan <- flag
}

func main() {

	numChan := make(chan int, 100)
	resChan := make(chan int, 50)
	exitChan := make(chan bool, 4)

	go getNum(numChan, 100)

	for i := 0; i < 4; i++ {
		go resNum(numChan, resChan, exitChan)
	}

	//开启一个协程将读完数据的协程都释放掉
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//这时候说明所有的协程都读完了数据，这时候再关闭通道
		close(resChan)
	}()

	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
