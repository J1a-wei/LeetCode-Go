package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, m)

	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}

			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	steps := make([][]int, m)
	for i := 0; i < len(obstacleGrid); i++ {
		steps[i] = make([]int, n)
	}

	// 初始化第一列，只要碰到一个1，后面都无法走到
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		steps[i][0] = 1
	}

	// 初始化第一行，碰到一个1，后面都无法找到
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}

		steps[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				steps[i][j] = 0
			} else {
				steps[i][j] = steps[i-1][j] + steps[i][j-1]
			}
		}
	}

	return steps[m-1][n-1]
}
