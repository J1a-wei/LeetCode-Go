package leetcode22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "test",
			arg:  3,
			want: 5,
		},
	}

	for _,tt := range tests {
		assert.Equal(t,tt.want,len(generateParenthesis(tt.arg)))
	}


}
