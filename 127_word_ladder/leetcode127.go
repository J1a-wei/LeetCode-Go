package leetcode

import "math"

// DFS 直接超时了😓
func ladderLength(beginWord string, endWord string, wordList []string) (result int) {
	depth := 0
	if len(beginWord) != len(endWord) {
		return depth
	}
	used := make([]bool, len(wordList), len(wordList))

	var dfs func(int, string)
	dfs = func(depth int, check string) {
		if check == endWord {
			if result == 0 {
				result = depth
			} else {
				result = min(result, depth)
			}
			return
		}
		for index, v := range wordList {
			if !used[index] && isConvertWord(check, v) {

				used[index] = true
				dfs(depth+1, v)

				used[index] = false
			}
		}

	}
	dfs(1, beginWord)
	return

}

func min(result int, depth int) int {
	if result < depth {
		return result
	}
	return depth
}

func isConvertWord(str1, str2 string) bool {
	count := 0
	if len(str1) != len(str2) {
		return false
	}
	b1 := []byte(str1)
	b2 := []byte(str2)

	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			count++
		}
		if count >= 2 {
			return false
		}
	}

	if count == 1 {
		return true
	}
	return false
}

func ladderLengthBFS(beginWord, endWord string, wordList []string) int {
	wordMap, que, depth := getWordMap(wordList, beginWord), []string{beginWord}, 0

	for len(que) > 0 {
		depth++
		qlen := len(que)

		for i := 0; i < qlen; i++ {
			word := que[i]
			candidates := getCandidates(word)

			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					if candidate == endWord {
						return depth + 1
					}
					delete(wordMap, candidate)
					que = append(que, candidate)
				}

			}
		}
		que = que[qlen:]
	}
	return 0
}

func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)

	for i, word := range wordList {
		if word != beginWord {
			wordMap[word] = i
		}
	}
	return wordMap
}

func getCandidates(word string) []string {
	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a')+i) {
				res = append(res, word[:j]+string(byte(int('a')+i))+word[j+1:])
			}
		}
	}
	return res
}

func getCandidates2(beginWord string) []string {
	var res []string

	for i := 0; i < len(beginWord); i++ {
		for j := 0; j < 26; j++ {
			alpha := byte(int('a') + j)
			if beginWord[i] != alpha {
				res = append(res, beginWord[:i]+string(alpha)+beginWord[i+1:])
			}
		}
	}
	return res
}

// 官方题解: 广度优先 + 优化建图
func ladderLengthMain(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

// 官方题解 双向BFS
func ladderLengthMain2(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	distBegin := make([]int, len(wordId))
	for i := range distBegin {
		distBegin[i] = inf
	}
	distBegin[beginId] = 0
	queueBegin := []int{beginId}

	distEnd := make([]int, len(wordId))
	for i := range distEnd {
		distEnd[i] = inf
	}
	distEnd[endId] = 0
	queueEnd := []int{endId}

	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		q := queueBegin
		queueBegin = nil
		for _, v := range q {
			if distEnd[v] < inf {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, w := range graph[v] {
				if distBegin[w] == inf {
					distBegin[w] = distBegin[v] + 1
					queueBegin = append(queueBegin, w)
				}
			}
		}

		q = queueEnd
		queueEnd = nil
		for _, v := range q {
			if distBegin[v] < inf {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, w := range graph[v] {
				if distEnd[w] == inf {
					distEnd[w] = distEnd[v] + 1
					queueEnd = append(queueEnd, w)
				}
			}
		}
	}
	return 0
}

// LeetCode 玩家解 beat 100%
// 这代码真的太秀了！
func ladderLength100(beginWord string, endWord string, wordList []string) int {
	step := 0

	wordMap := make(map[string]int)

	for i, w := range wordList {
		wordMap[w] = i
	}

	startQueue := []string{beginWord}
	endQueue := []string{endWord}

	startUsed := make([]bool, len(wordList))
	endUsed := make([]bool, len(wordList))

	if i, ok := wordMap[endWord]; !ok {
		return 0
	} else {
		endUsed[i] = true //  endQueue 是从后往前走的
	}

	for len(startQueue) > 0 {
		step++
		l := len(startQueue) // BFS 遍历下一层的时候，go语言有好的处理方案，不需要重新new一个queue，通过切片截取完成新组"创建"

		for i := 0; i < l; i++ {
			chars := []byte(startQueue[i])
			for j := 0; j < len(chars); j++ {
				o := chars[j]

				for c := 'a'; c <= 'z'; c++ {
					chars[j] = byte(c)
					idw, ok := wordMap[string(chars)]

					if !ok || startUsed[idw] {
						continue
					}

					if endUsed[idw] { // 相撞
						return step + 1
					} else {
						startQueue = append(startQueue, string(chars))
						startUsed[idw] = true
					}
				}

				chars[j] = o
			}
		}
		startQueue = startQueue[l:]

		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
			startUsed, endUsed = endUsed, startUsed
		}
	}

	return 0
}
