package leetcode49

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		arg  []string
		want [][]string
	}{
		{
			arg: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				[]string{"bat"},
				[]string{"nat", "tan"},
				[]string{"ate", "eat", "tea"},
			},
		},
	}

	for _,tt := range tests {
		t.Run("test",func(t *testing.T) {
			assert.Equal(t,tt.want,groupAnagrams(tt.arg))
		})
	}
}
