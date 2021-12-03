package main

// 并查集模版
type DisjointSet struct {
	fa []int
}

func Construct(n int) *DisjointSet {
	s := DisjointSet{fa: make([]int, n)}
	for i := 0; i < n; i++ {
		s.fa[i] = i
	}
	return &s
}
func (s *DisjointSet) Find(x int) int {
	if s.fa[x] != x {
		s.fa[x] = s.Find(s.fa[x])
	}
	return s.fa[x]
}
func (s *DisjointSet) Join(x, y int) {
	x, y = s.Find(x), s.Find(y)
	if x != y {
		s.fa[x] = y
	}
}

// 时间复杂度 O(m*n)，m, n 分别为 grid 的行数和列数，因为操作并查集的时间复杂度为O(α(n))，α(n) 为反阿克曼函数，近似常数。代码中使用了一个 for 循环（该循环遍历了所有数组元素）执行并查集的查找和合并，所以时间复杂度为O(α(n)*m*n)， 即 O(m*n)。
// 空间复杂度 O(m*n)，因为并查集需要一个长度等同 grid 元素数量的数组辅助计算，所以空间复杂度在 O(m*n)。
func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	// 将二维坐标转为一维
	num := func(i, j int) int {
		return i*n + j
	}

	// 方向数组
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}
	// 初始化并查集
	d := Construct(m * n)
	// 初始化答案
	ans = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 跳过海水
			if grid[i][j] == '0' {
				continue
			}

			// 将当前点坐标转为一维
			key := num(i, j)
			// 计算当前点父节点
			fa := d.Find(key)
			// 更新岛屿总数，暂时将每个 1 记做一个岛屿，每一次有效的合并都对其减一
			// 有效的合并，指的是合并两个不同父节点的相邻的 1
			ans++
			for k := 0; k < 4; k++ {
				// 向四周移动
				nx, ny := i+dx[k], j+dy[k]
				// 跳过越界和海水的情况
				if nx < 0 || nx >= m || ny < 0 || ny >= n ||
					grid[nx][ny] == '0' {
					continue
				}
				// 如果 next 的点的父节点和当前点的父节点不一致，就将其和并（有效合并）
				keyn := num(nx, ny)
				if d.Find(keyn) != fa {
					d.Join(keyn, key)
					// 开始将所有的 1 判断为独立岛屿，每次有效的合并，就说明将两个不同父节点的连通岛屿合并在一起
					ans--
				}
			}
		}
	}
	return
}
