package main

import (
	"errors"
	"fmt"
)

//定义除数为0的错误
var errDivisionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0,errDivisionByZero
	}

	return dividend / divisor,nil
}

func main() {

	fmt.Println(div(1,0))
}
