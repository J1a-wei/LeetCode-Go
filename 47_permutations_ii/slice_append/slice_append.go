package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := []int{}
	fmt.Printf("%v\n", unsafe.Pointer(&s))
	for i := 0; i < 100; i++ {
		s = append(s, i)
		fmt.Printf("len(%d),cap(%d),%v，&v\n", len(s), cap(s), unsafe.Pointer(&s), &s)
	}
}
