package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	ret := []*Node{}

	n := &Node{}

	ret = append(ret, n.Children...)
	fmt.Println(ret)
}
