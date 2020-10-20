package leetcode283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T){
	tests := []struct{
		name string 
		arg []int
		want []int
	}{
		{
			"FastSlowPorinter",
			[]int{0,1,0,3,12},
			[]int{1,3,12,0,0},
		},
	}

	for _,tt := range tests{
		t.Run(tt.name,func(t *testing.T){
			moveZeroes(tt.arg)
			assert.Equal(t,tt.want,tt.arg)
		})
	}
}