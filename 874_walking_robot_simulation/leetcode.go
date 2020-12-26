package leetcode

import "fmt"

var dirs = [][]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func robotSim(commands []int, obstacles [][]int) int {
	obstaclesMap := map[string]bool{}

	for _, v := range obstacles {
		str := fmt.Sprintf("%v %v", v[0], v[1])
		obstaclesMap[str] = true
	}

	maxFunc := func(n, m int) int {
		if n > m {
			return n
		}
		return m
	}
	x, y, d, ans := 0, 0, 0, 0

	for _, v := range commands {
		if v == -1 { // 顺时针转
			d++
			if d == 4 {
				d = 0
			}
		} else if v == -2 { // 逆时针转
			d--
			if d == -1 {
				d = 3
			}
		} else {
			for i := 0; i < v; i++ {
				key := fmt.Sprintf("%v %v", x+dirs[d][0], y+dirs[d][1])

				if _, ok := obstaclesMap[key]; !ok {
					x += dirs[d][0]
					y += dirs[d][1]
				}
			}
		}
		ans = maxFunc(ans, x*x+y*y)
	}
	return ans
}
