package main

import "fmt"

type valueNode struct {
	row int
	col int
	val int
}

func main() {

	res := [11][11]int{}
	res[1][2] = 1
	res[2][3] = 2

	//for _, k := range res {
	//	for _, v := range k {
	//		fmt.Printf("%d\t", v)
	//	}
	//	fmt.Println()
	//}

	var valArr []valueNode

	for i, v := range res {
		for j, v1 := range v {
			if v1 != 0 {
				val := valueNode{
					col: i,
					row: j,
					val: v1,
				}
				valArr = append(valArr, val)
			}
		}
	}

	fmt.Println("当前稀疏数组")
	for _, va := range valArr {
		fmt.Printf("%d\t%d\t%d\n", va.col, va.row, va.val)
	}

	var spaArr [11][11]int
	for i,k2 := range valArr{
		if i != 0 {
			spaArr[k2.col][k2.row] = k2.val
		}
	}

	fmt.Println("稀疏数组 -> 原始数组:")
	for _, k := range spaArr {
		for _, v := range k {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}

}
