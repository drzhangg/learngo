package main

import (
	"fmt"
	"time"
)

type Slice []int

func NewSilce() Slice {
	return make(Slice,0)
}

func (s *Slice) Add(i int) *Slice {
	*s = append(*s,i)
	fmt.Print(i)
	return s
}

func main() {
	s := NewSilce()
	defer s.Add(1).Add(2)
	time.Sleep(time.Second)
	s.Add(3)
}
