package leetcode77

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		arg1 int
		arg2 int
		want [][]int
	}{
		{
			"test",
			4,
			2,
			[][]int{
				[]int{2, 4},
				[]int{3, 4},
				[]int{2, 3},
				[]int{1, 2},
				[]int{1, 3},
				[]int{1, 4},
			},
		},
	}

	for _, test := range tests {
		t.Run("debug", func(t *testing.T) {
			assert.Equal(t, test.want, combine5(test.arg1, test.arg2))

		})
	}
}

func TestSlice(t *testing.T) {
	ans := [][]int{}
	s1 := []int{1}
	s2 := []int{2}

	ans = append(ans, s1)
	ans = append(ans, s2)

	for i, v := range ans {
		v = append(v, 3)
		ans[i] = v
		fmt.Println(v)
	}
	for _, v := range ans {
		fmt.Println(v)
	}

}
