package problems

// Solution to 994. Rotting Oranges
// Multi-source bfs algorithim

func orangesRotting(grid [][]int) int {
	cols, rows := len(grid[0]), len(grid)

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := make([][]int, 0)
	time := 0
	fresh := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	if fresh < 1 {
		return 0
	}

	for len(queue) > 0 {
		for range len(queue) {
			curr := queue[0]
			queue = queue[1:]

			for _, dir := range dirs {

				nx, ny := curr[0]+dir[0], curr[1]+dir[1]

				if nx < 0 || nx > rows-1 || ny < 0 || ny > cols-1 || grid[nx][ny] != 1 {
					continue
				}

				grid[nx][ny] = 2
				fresh--
				queue = append(queue, []int{nx, ny})
			}
		}
		time++
		if fresh == 0 {
			return time
		}
	}


	return -1
}

// Solution to 542. 01 Matrix
// Multi-source bfs algorithim

func updateMatrix(mat [][]int) [][]int {

	m, n, queue := len(mat), len(mat[0]), make([][]int, 0)

	for i, v := range mat {
		for j := range v {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else if mat[i][j] == 1 {
				mat[i][j] = -1
			}
		}
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}


	for len(queue) > 0 {

		curr := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {

			nx, ny := curr[0]+dir[0], curr[1]+dir[1]

			if nx < 0 || nx >= m || ny < 0 || ny >= n || mat[nx][ny] != -1{continue}

			mat[nx][ny] = mat[curr[0]][curr[1]] + 1

			queue = append(queue, []int{nx, ny})

		}

	}

	return mat
}