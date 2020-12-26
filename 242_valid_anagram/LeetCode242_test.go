package leetcode242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestisAnagram(t *testing.T){
	tests:= []struct {
		name string 
		arg1 string
		arg2 string 
		want bool
	}{
		{
			"normal",
			"anagram",
			"nagaram",
			true,
		},
	}
	
	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T) {
			assert.Equal(t,tt.want,isAnagram(tt.arg1,tt.arg2))
		})
	}

}