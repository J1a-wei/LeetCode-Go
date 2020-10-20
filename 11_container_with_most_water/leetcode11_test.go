package leetcode11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T){
	tests := []struct{
		name string
		arg []int
		want int
	}{
		{
			"Double Pointer",
			[]int{1,8,6,2,5,4,8,3,7},
			49,
		},
		{
			"Force soulution",
			[]int{1,8,6,2,5,4,8,3,7},
			49,
		},
	}
	
	for _,tt:=range tests{
		t.Run(tt.name,func(t *testing.T){
			assert.Equal(t,tt.want,maxArea(tt.arg))
		})
	}

}