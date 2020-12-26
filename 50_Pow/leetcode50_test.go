package leetcode50

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {
	x := 2.0
	n := -3
	res := myPow(x, n)
	fmt.Println(res)
	res2 := myPow2(x, n)
	fmt.Println(res2)
}
