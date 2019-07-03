package main

import "fmt"

var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}

func quickSelect(sli []int) []int {
	len := len(sli)
	if len <= 1 {
		return sli
	}

	base_num := sli[0]

	left_sli := []int{}
	right_sli := []int{}
	for i := 1; i < len; i++ {
		if base_num >sli[i] {
			left_sli = append(left_sli,sli[i])
		}else {
			right_sli = append(right_sli,sli[i])
		}
	}

	left_sli = quickSelect(left_sli)
	right_sli = quickSelect(right_sli)

	left_sli = append(left_sli,base_num)
	return append(left_sli,right_sli...)
}

func main() {
	res := quickSelect(sli)
	fmt.Println(res)

}
