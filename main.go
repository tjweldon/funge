package main

import (
	"fmt"
	"funge/internal/util"
)

func main() {
	s := util.NewStack[int]()
	fmt.Println(s)
	s.Push(100)
	fmt.Println(s)
	s.Push(200)
	fmt.Println(s)
	last := s.Pop()
	fmt.Println(last, s)
}
