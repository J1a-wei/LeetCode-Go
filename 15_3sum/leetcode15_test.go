package leetcode15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3Sum(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want [][]int
	}{
		{	"3sum",
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				[]int{-1, 0, 1},
				[]int{-1, -1, 2},
			},
		},
	}

	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T) {
			assert.Equal(t,tt.want,threeSum(tt.arg))
		})
	}

}
