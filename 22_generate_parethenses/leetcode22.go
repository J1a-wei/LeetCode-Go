package leetcode22

func generateParenthesis(n int) []string {
	ret := []string{}
	generate(n, 0, 0, "", &ret)
	return ret
}

/*
	规则： 1. 只能生成 '(' ')'
		  2. 生成'(' 只要小于n，即可生成
		  3. 生成')' 小于n

*/
func generate(n, left, right int, s string, ret *[]string) {
	// 终止条件
	if left == n && right == n {
		*ret = append(*ret, s)
		return
	}

	if left < n {
		// s = s + "("    BUG 写法 ,让我卡了好久
		generate(n, left+1, right, s+"(", ret)
	}

	if left > right {
		// 按照这样写就ok
		generate(n, left, right+1, s+")", ret)
	}

}
