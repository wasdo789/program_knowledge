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
