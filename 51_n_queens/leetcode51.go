package leetcode51

func solveNQueens(n int) [][]string {
	// 思考点1： dia1 dia2 为什么是2n-1， 这里实际上是斜对角线的个数，对于一个对称棋牌而言，对角线的个数就是2 * n - 1
	// 棋牌就是row，index值和value 分别代表x 和 y，前面三个走限制流程
	col, dia1, dia2, row, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, [][]string{}

	putQueen(n, 0, &col, &dia1, &dia2, &row, &res)
	return res
}

// index的值实际上控制的是y的值
func putQueen(n int, index int, col *[]bool, dia1 *[]bool, dia2 *[]bool, row *[]int, res *[][]string) {
	if index == n {
		*res = append(*res, generateBoard(n, row))
		return
	}
	for i := 0; i < n; i++ {
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i] = true
			(*dia2)[index-i+n-1] = true
			putQueen(n, index+1, col, dia1, dia2, row, res)

			(*col)[i] = false
			(*dia1)[index+i] = false
			(*dia2)[index-i+n-1] = false
			*row = (*row)[:len(*row)-1]
		}
	}
	return
}

func generateBoard(n int, row *[]int) []string {
	board := []string{}
	res := ""
	for i := 0; i < n; i++ {
		res += "."
	} // 生成一行全部为点的字符串

	for i := 0; i < n; i++ {
		board = append(board, res)
	} // 把整个棋谱都先设置成 点

	for i := 0; i < n; i++ {
		tmp := []byte(board[i])
		tmp[(*row)[i]] = 'Q'
		board[i] = string(tmp)

	}
	return board
}
