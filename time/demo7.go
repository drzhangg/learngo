package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	yearStr := time.Now().Format("2006")
	monthStr := time.Now().Format("1")
	yearInt, _ := strconv.Atoi(yearStr)
	monthInt, _ := strconv.Atoi(monthStr)

	if monthInt == 1 {
		yearInt -= 1
		monthInt = 12
	} else {
		monthInt -= 1
	}

	if monthInt < 10 {
		monthStr = "0" + strconv.Itoa(monthInt)
	}

	fmt.Println(strconv.Itoa(yearInt) + "-" + monthStr)
}
