package main

import "fmt"

func isMatch(s string, p string) bool {
	memo := map[string]bool{}
	var dp func(int, int) bool
	dp = func(i, j int) bool {
		m, n := len(s), len(p)
		if j == n {
			return i == m
		}
		if i == m {
			if (n-j)%2 == 1 {
				return false
			}
			for ; j+1 < n; j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true

		}
		// record state which is by key  now
		key := fmt.Sprint(i) + "," + fmt.Sprint(j)
		// memo 的key, 要用坐标表示key，有肯能存在值相等但是key不同的value
		
		if res, ok := memo[key]; ok {
			return res
		}

		res := false

		if s[i] == p[j] || p[j] == '.' {
			if j < n-1 && p[j+1] == '*' {
				// 匹配到*,,可能出现0次或者多次
				res = dp(i, j+2) || dp(i+1, j)
			} else {
				// 匹配到. , 1:1匹配即可
				res = dp(i+1, j+1)
			}
		} else {
			// 匹配0次
			if j < n-1 && p[j+1] == '*' {
				res = dp(i, j+2)
			} else {
				res = false
			}
		}

		memo[key] = res
		return res
	}

	return dp(0, 0)

}
func main() {
	s := "aa"
	p := "a"
	fmt.Println(isMatch(s, p))
}
