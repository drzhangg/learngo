package main

import "fmt"

func main() {
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	for k, v := range team {
		fmt.Println(k, v)
	}
}
