package main

import "fmt"

//将int定义为MyInt类型
type MyInt int

//为MyInt添加IsZero()方法
func (m MyInt) IsZero() bool {
	return m == 0
}

//为MyInt添加Add()方法
func (m MyInt) Add(other int) int {
	return other + int(m)
}

func main() {

	var b MyInt

	fmt.Println(b.IsZero())

	b = 1
	fmt.Println(b.Add(2))
	fmt.Println(b.IsZero())
	fmt.Println(b)
}
