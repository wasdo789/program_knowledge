package leetcode

func bfsnums(grid [][]byte, i, j int) {
	row := len(grid)
	col := len(grid[0])
	if i < 0 || i >= row || j < 0 || j >= col || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	bfsnums(grid, i, j-1)
	bfsnums(grid, i, j+1)
	bfsnums(grid, i-1, j)
	bfsnums(grid, i+1, j)

}
func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	cnt := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != '1' {
				continue
			}
			bfsnums(grid, i, j)
			cnt++
		}
	}
	return cnt
}

type Point struct {
	Row int
	Col int
}

func searchboard(row int, col int, board [][]byte) {
	if row >= 0 && row < len(board) && col >= 0 && col < len(board[0]) && board[row][col] == 'O' {
		board[row][col] = 'A'
		searchboard(row-1, col, board)
		searchboard(row+1, col, board)
		searchboard(row, col-1, board)
		searchboard(row, col+1, board)
	}
}

func solve(board [][]byte) {
	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])
	if col == 0 {
		return
	}
	for i := 0; i < row; i++ {
		// dfs board[i][0]
		searchboard(i, 0, board)
		// dfs board[i][col-1]
		if col-1 != 0 {
			searchboard(i, col-1, board)
		}

	}
	for j := 1; j < col-1; j++ {
		// dfs board[0][j]
		searchboard(0, j, board)
		// dfs board[row-1][j]
		if row-1 != 0 {
			searchboard(row-1, j, board)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}
