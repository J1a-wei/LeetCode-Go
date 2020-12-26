package leetcode

/*


给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：

每次转换只能改变一个字母。
转换后得到的单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回一个空列表。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]
示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: []

解释:endWord "cog" 不在字典中，所以不存在符合要求的转换序列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-ladder-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/

// 本题weiwei大佬给的题解是bfs + dfs. 先挖一个小坑，为什么在bfs的过程当中，不同时就把满足题意的path记录下来呢？

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := map[string]bool{}
	res := [][]string{}
	for _, word := range wordList {
		wordSet[word] = false
	}
	if _, contains := wordSet[endWord]; !contains {
		return res
	}

	// step1 使用广度优先遍历得到后继节点列表
	successor := map[string]map[string]bool{}

	found := bfs(beginWord, endWord, wordSet, successor)
	if !found {
		return res
	}

	// step2. 基于后继节点列表 successors，使用回溯算法得到所有最短路径列表
	path := []string{beginWord}
	dfs(beginWord, endWord, successor, path, &res)
	return res
}

func dfs(beginWord string, endWord string, successor map[string]map[string]bool, path []string, res *[][]string) {
	if beginWord == endWord {
		newPath := make([]string, len(path))
		copy(newPath, path)
		*res = append(*res, newPath)
	}
	if _, ok := successor[beginWord]; !ok {
		return
	}
	successorWords := successor[beginWord]

	for k := range successorWords {
		path = append(path, k)
		dfs(k, endWord, successor, path, res)
		path = path[:len(path)-1]
	}

}

func bfs(beginWord string, endWord string, wordSet map[string]bool, successor map[string]map[string]bool) bool {
	queue := []string{beginWord}

	// 记录访问过的单词
	visted := map[string]bool{beginWord: true}

	found := false

	wordLen := len(beginWord)

	// 当前层访问过的节点，当前层全部遍历完成后，再添加到总的visited集合里
	nextLevelVisited := map[string]bool{}

	for len(queue) != 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			currentWord := queue[0]
			queue = queue[1:]
			currentWordByte := []byte(currentWord)

			for j := 0; j < wordLen; j++ {
				originByte := currentWordByte[j]

				for k := 'a'; k <= 'z'; k++ {
					if currentWordByte[j] == byte(k) {
						continue
					}
					currentWordByte[j] = byte(k)
					nextWord := string(currentWordByte)

					if _, ok := wordSet[nextWord]; ok && !visted[nextWord] {
						if endWord == nextWord {
							found = true
						}

						// 避免下层元素重复加入队列，注意是队列
						if _, ok := nextLevelVisited[nextWord]; !ok {
							queue = append(queue, nextWord)
							nextLevelVisited[nextWord] = true
						}

						if v, ok := successor[currentWord]; ok {
							v[nextWord] = true
						} else {
							newSet := map[string]bool{}
							newSet[nextWord] = true
							successor[currentWord] = newSet
						}
					}

				}
				currentWordByte[j] = originByte

			}
		}

		if found {
			break
		}

		for k := range nextLevelVisited {
			visted[k] = true
		}
		nextLevelVisited = map[string]bool{}

	}
	return found
}

func findLadders233(beginWord string, endWord string, wordList []string) [][]string {
	result, graph, queue, set := [][]string{}, map[string][]string{}, []string{beginWord}, map[string]bool{}
	for _, v := range wordList {
		set[v] = true
	}
	if _, ok := set[endWord]; !ok {
		return result
	}

	// 先通过bfs把graph绘制出来
	nextLevelMap := map[string]bool{}
	for len(queue) != 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			curWord := queue[i]
			curByte := []byte(curWord)
			graphSlice := []string{}

			for j := 0; j < len(curByte); j++ {
				originByte := curByte[j]
				for k := 'a'; k <= 'z'; k++ {
					curByte[j] = byte(k)
					newStr := string(curByte)
					if _, ok := set[newStr]; ok {
						graphSlice = append(graphSlice, newStr)
						if _, ok := nextLevelMap[newStr]; !ok {
							queue = append(queue, newStr)
							nextLevelMap[newStr] = true
						}
						// delete(set,newStr)

					}
				}
				curByte[j] = originByte
			}
			graph[curWord] = graphSlice

		}

		queue = queue[qLen:]
		for k := range nextLevelMap {
			delete(set, k)
		}
		nextLevelMap = map[string]bool{}

	}

	path := []string{beginWord}
	var dfs func(string, string, []string)
	dfs = func(beginWord string, endWord string, path []string) {
		if beginWord == endWord {
			newPath := make([]string, len(path))
			copy(newPath, path)
			result = append(result, newPath)
			return
		}

		nextWords := graph[beginWord]
		for _, v := range nextWords {
			path = append(path, v)
			dfs(v, endWord, path)
			path = path[:len(path)-1]
		}

	}
	dfs(beginWord, endWord, path)

	return result
}
