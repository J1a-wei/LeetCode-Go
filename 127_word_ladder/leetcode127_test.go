package leetcode

import (
	"testing"
	"unsafe"
)

func TestLeetcode(t *testing.T) {
	worldList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "hit"
	endWrod := "cog"

	ret := ladderLengthBFS(beginWord, endWrod, worldList)
	t.Log(ret)

}

func TestCandidates(t *testing.T) {
	test := "hat"

	result := getCandidates2(test)
	t.Log(result)
}

func TestString(t *testing.T) {
	s := "HelloWorld"

	s1 := s[:len(s)]
	t.Log(s1)

	t.Log(unsafe.Pointer(&s1))
	t.Log(unsafe.Pointer(&s))
}
