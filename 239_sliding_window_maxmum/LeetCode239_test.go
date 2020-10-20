package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindowQueue(t *testing.T) {

	tests := []struct {
		name string
		arg1 []int
		arg2 int
		want []int
	}{
		{
			name: "slidingWindowMaxmum",
			arg1: []int{1, 3, -1, -3, 5, 3, 6, 7},
			arg2: 3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
	}

	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T){
			assert.Equal(t,tt.want,maxSlidingWindowQueue(tt.arg1,tt.arg2))
		})
	}

}


func TestMaxSlidingWindowForce(t *testing.T) {

	tests := []struct {
		name string
		arg1 []int
		arg2 int
		want []int
	}{
		{
			name: "slidingWindowMaxmum",
			arg1: []int{1, 3, -1, -3, 5, 3, 6, 7},
			arg2: 3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
	}

	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T){
			assert.Equal(t,tt.want,maxSlidingWindowForce(tt.arg1,tt.arg2))
		})
	}

}