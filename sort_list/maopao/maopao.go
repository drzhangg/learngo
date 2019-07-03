package main

import "fmt"

var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}

func maopao(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	res := maopao(sli)
	fmt.Println(res)
}
