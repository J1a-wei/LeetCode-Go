package leetcode77

import "fmt"

//
// k 个数的组合
// 1...n中 k个数的组合
func combine(n int, k int) [][]int {
	var result [][]int = [][]int{}
	helper(n, k, 0, nil, &result)
	return result
}

func helper(n, k, i int, sub []int, result *[][]int) {
	if i == k {

		*result = append(*result, sub)
		return
	}
	if i == 0 {
		sub = []int{}
	}

	for x := n; x >= 1; x-- {
		comb := make([]int, len(sub)) // 这里一定记得重新开辟一个，然后进行拷贝，否则会出现数据错乱
		copy(comb, sub)
		comb = append(comb, x)
		helper(x-1, k, i+1, comb, result)
	}
}

// 闭包写法
func combine2(n int, k int) (ans [][]int) {
	tmp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝: tmp长度加上区间[cur,n]的长度小于k,不可能构造出长度为k的tmp, 进行了加速
		if len(tmp)+(n-cur+1) < k {
			return
		}

		// 记录答案
		if len(tmp) == k {
			comb := make([]int, k)
			copy(comb, tmp)
			ans = append(ans, comb)
			return
		}

		//考虑当前位置
		tmp = append(tmp, cur)
		dfs(cur + 1)

		// 不考虑当前位置
		tmp = tmp[:len(tmp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return
}

// 数学公式法递归
// n 中 取 k个，等于选了一个，从N-1中取 k-1个，或者没选这个，从n-1中，取K个
func combine3(n int, k int) (ans [][]int) {
	if n < k || k <= 0 {
		return
	}

	// 等于选了一个，从N-1中取 k-1个
	ans = combine3(n-1, k-1)
	if len(ans) == 0 {
		s := []int{}
		ans = append(ans, s)
	}

	for i, s := range ans {
		s = append(s, n) // 所有的组合加上当前选中的这个
		ans[i] = s       // 其他语言没必要加这句，但是Go语言必须要加
		// 目前这种写法，肯定发生了很多次拷贝...
	}

	ans = append(ans, combine3(n-1, k)...)
	return
}

// 技巧性太强了..不建议一开始花时间
func combine4(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	temp := []int{}
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			b := temp[j]+1 == temp[j+1]
			fmt.Println(b)
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}

func combine5(n, k int) (res [][]int) {
	var dfs func(int, []int)
	visted := make([]bool, n+1)
	dfs = func(count int, s []int) {
		if len(s) == k {
			res = append(res, append([]int{}, s...))
			return
		}

		for i := count; i <= n; i++ {
			if !visted[i] {
				visted[i] = true
				s = append(s, i)
				dfs(i+1, s)
				s = s[:len(s)-1]
				visted[i] = false
			}
		}

	}

	dfs(1, []int{})
	return
}
