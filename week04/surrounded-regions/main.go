package main

// 时间复杂度 O(m*n)， m 为矩阵行数，n 为矩阵列数，因为使用了 bfs 广度优先搜索，所以会遍历所有的矩阵元素，而矩阵元素数量为 O(m*n)，所以时间复杂度为 O(m*n)。
// 空间复杂度 O(m*n)， m 为矩阵行数，n 为矩阵列数，因为定义了一个队列用作 bfs 广度优先搜索，队列的容量最大不会超过矩阵元素数量 m*n，所以空间复杂度为 O(m*n)。
func solve(board [][]byte) {
	// 定义方向数组
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	m, n := len(board), len(board[0])
	// 定义一个队列，队列内放置所有的未被包围的字符 O 的 x，y 坐标
	q := make([][]int, 0, m*n)
	// 从矩阵的四周，将所有的字符 O 标记为字符 A
	// 并将标记后的字符 A 加入队列
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			board[i][0] = 'A'
			q = append(q, []int{i, 0})
		}
		if board[i][n-1] == 'O' {
			board[i][n-1] = 'A'
			q = append(q, []int{i, n - 1})
		}
	}
	for i := 1; i < n-1; i++ {
		if board[0][i] == 'O' {
			board[0][i] = 'A'
			q = append(q, []int{0, i})
		}
		if board[m-1][i] == 'O' {
			board[m-1][i] = 'A'
			q = append(q, []int{m - 1, i})
		}
	}

	// 遍历队列，找到所有未被包围的字符 O 的附近的所有同类，并将其替换为字符 A
	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		x, y := front[0], front[1]
		// 往四个方向 bfs
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			// 是否已经越界，或者是否遇到 'X' 或已经被遍历过的 'O' （被遍历过的未被包围的 O 都会被更新成 A）
			if nx < 0 || nx >= m || ny < 0 || ny >= n || board[nx][ny] != 'O' {
				continue
			}

			// 标记未被包围的坐标，并将其加入队列，往四周 bfs
			board[nx][ny] = 'A'
			q = append(q, []int{nx, ny})
		}
	}

	// 遍历完整矩阵，将所有余下的字符 O 替换成字符 X，余下的字符 A 替换成字符 O
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
	return
}
