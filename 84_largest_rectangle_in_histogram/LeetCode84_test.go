package leetcode84

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleAreaForce(t *testing.T) {
	// type args struct{
	// 	heights []int
	// }
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "largestRectangleArea",
			args: []int{2, 1, 5, 6, 2, 3},
			want: 10,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, largestRectangleAreaForce(tt.args))
		})
	}

}

func TestLargestRectangleAreaForce2(t *testing.T) {
	// type args struct{
	// 	heights []int
	// }
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "largestRectangleArea",
			args: []int{2, 1, 5, 6, 2, 3},
			want: 10,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, largestRectangleAreaForce2(tt.args))
		})
	}

}

func TestLargestRectangleAreaStack(t *testing.T) {
	// type args struct{
	// 	heights []int
	// }
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "largestRectangleArea",
			args: []int{2, 1, 5, 6, 2, 3},
			want: 10,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, largestRectangleAreaForceStack(tt.args))
		})
	}

}