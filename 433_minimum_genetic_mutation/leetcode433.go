package leetcode433

// DFS
func minMutation(start string, end string, bank []string) (res int) {
	var dfs func(string, int)
	bankVisited := make([]bool, len(bank), len(bank))
	dfs = func(check string, depth int) {
		if check == end {
			res = depth
			return
		}
		if depth == len(bank) {
			res = -1
			return
		}

		for index, value := range bank {
			if !bankVisited[index] && convertIsLegal(check, value) {
				bankVisited[index] = true
				dfs(value, depth+1)
				// 回溯，没访问这个元素
				bankVisited[index] = false
			}
		}

	}
	dfs(start, 0)
	return
}

// 检查是否合法
func convertIsLegal(s1, s2 string) bool {
	count := 0
	b1 := []byte(s1)
	b2 := []byte(s2)
	if len(b1) != len(b2) {
		return false
	}

	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			count++
		}
	}

	if count == 1 {
		return true
	}

	return false
}

/*
	BFS解法也许是本题的本意，通过bank逆推start感觉有点取巧了

	每个字母一共可以有四种变化 (A C G T)

	依次改动start的每个字母的可能性，如果bank当中有等于start改动后的，把元素推入到queue当中

*/
var mutationMap = map[uint8][3]string{
	'A': [...]string{"T", "G", "C"},
	'C': [...]string{"T", "G", "A"},
	'T': [...]string{"A", "G", "C"},
	'G': [...]string{"T", "A", "C"},
}

func indexOf(str string, bank []string) int {
	for i, s := range bank {
		if s == str {
			return i
		}
	}
	return -1
}

func minMutation2(start string, end string, bank []string) int {
	if indexOf(end, bank) == -1 {
		return -1
	}

	isUsed := make([]bool, len(bank))
	queue := []string{start}
	count := 0

	for len(queue) != 0 {
		l := len(queue)

		for i := 0; i < l; i++ {
			curr := queue[i]
			if curr == end {
				return count
			}

			for j := 0; j < len(curr); j++ {
				for _, s := range mutationMap[curr[j]] {
					if idx := indexOf(curr[:j]+s+curr[j+1:], bank); idx != -1 && !isUsed[idx] {
						queue = append(queue, bank[idx])
						isUsed[idx] = true
					}
				}
			}

		}
		count++
		queue = queue[1:]

	}
	return -1
}
