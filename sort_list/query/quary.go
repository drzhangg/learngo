package main

import "fmt"

var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}

func selectSort(sli []int) []int {
	len := len(sli)

	for i := 0; i < len -1; i++ {
		//先假设最小值的位置
		k := i
		for j := i + 1; j < len; j++ {
			if sli[k] > sli[j] {
				k = j
			}
		}

		if k != i {
			sli[k], sli[i] = sli[i], sli[k]
		}
	}
	return sli
}

func main() {
	res := selectSort(sli)
	fmt.Println(res)
}
