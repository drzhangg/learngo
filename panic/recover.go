package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {

	func(){
		defer func() {
			err := recover()
			switch err.(type) {
			case runtime.Error:
				fmt.Println("Runtime error:",err)
			default:
				fmt.Println("User error:",err)
			}
		}()
		panic("强行宕机")
	}()

	fmt.Println("after panic")
	fmt.Printf("%c\n",filepath.Separator)
	fmt.Printf("%c\n",filepath.ListSeparator)
	fmt.Printf(filepath.Abs("./"))

	fmt.Println(filepath.IsAbs("./"))
}
