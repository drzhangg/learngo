package main

import (
	"fmt"
	"os"
)

func main() {
	fi ,err := os.Stat("./src")
	if err == nil {
		fmt.Println(fi.IsDir())
	}

	err = os.RemoveAll("./path")
	fmt.Println(err)


}
