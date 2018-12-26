package main

import "learngo/oop/employee"

func main() {
	e := employee.New("Sam","Adolf",30,20)
	e.LeavesRemaining()
}