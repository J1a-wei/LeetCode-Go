package leetcode

import "testing"

func TestMap(t *testing.T) {
	m := map[string]bool{}

	changeMap(m)

	t.Log(m)
}

func changeMap(m map[string]bool) {
	m["123"] = true
}

/*
"red"
"tax"
["ted","tex","red","tax","tad","den","rex","pee"]

"hit"
"cog"
["hot","dot","dog","lot","log","cog"]

*/
func TestLeetcode(t *testing.T) {
	begin := "a"
	end := "c"
	wordList := []string{"a", "b", "c"}

	result := findLadders(begin, end, wordList)
	t.Log(result)
}
