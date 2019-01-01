package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("被除数不能为0")
	}else {
		result = a / b
	}
	return
}

func main() {

	result,err := MyDiv(2,0)
	if err != nil {
		fmt.Println("err = ",err)
	}else {
		fmt.Println("result = ",result)
	}
}
