package main

import "fmt"

//写数据
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeData = %d\n", i)
	}
	//数据写完后关闭channel
	close(intChan)
}

//读数据
func readData(intChan chan int, exitChan chan bool, i int) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}

		fmt.Printf("第%d个协程读到数据，readData 读到数据= %v\n", i, v)
	}

	exitChan <- true
	close(exitChan)
}

func main() {

	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	//for i := 0; i < 3; i++ {
	go writeData(intChan)
	//}

	for i := 1; i <= 3; i++ {
		go readData(intChan, exitChan, i)
	}

	for {
		v, ok := <-exitChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
