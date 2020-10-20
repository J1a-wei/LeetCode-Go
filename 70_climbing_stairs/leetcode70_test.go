package leetcode70

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T){
	tests := []struct{
		name string
		arg int
		want int
	}{
		{
			"climb 4 steps",
			4,
			5,
		},
		{
			"climb 1 step",
			1,
			1,
		},
		{
			"climb 2 step",
			2,
			2,
		},
	}

	for _,tt := range tests{
		t.Run(tt.name,func(t *testing.T) {
			assert.Equal(t,tt.want,climbStairs(tt.arg))
		})
	}
}