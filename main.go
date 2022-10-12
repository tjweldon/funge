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
	t := *&s
	t.Push(300)
	fmt.Println(t)
	fmt.Println(s)
}
