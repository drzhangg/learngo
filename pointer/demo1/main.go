package main

import "fmt"

func main() {

	var a int = 10

	fmt.Printf("a = %d\n",a)
	fmt.Printf("&a = %v\n",&a)

	//只有指针类型能保存变量的地址
	var p *int
	p = &a
	fmt.Printf("p = %v, &a = %v\n",p,&a)

	*p = 123
	fmt.Printf("*p = %v, a = %v\n",*p,a)
}
