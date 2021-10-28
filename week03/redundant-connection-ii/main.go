package main

// 时间复杂度为 O(n^2)，n 为节点数量，主要耗时操作逆序遍历 edges 时，针对每一个元素做了 dfs，而 dfs 的耗时和树深度有关，树深度不超过 n，嵌套后时间复杂度为 O(n^2)。
// 空间复杂度 O(n+m)，m 为 edges 长度，因为借助了三个辅助变量，一个数组统计入度（消耗 n），一个数组统计 dfs 时走过的元素（消耗 n），一个邻接表保存图（消耗n+m），所以总空间复杂度为 O(n+m)。
func findRedundantDirectedConnection(edges [][]int) []int {
	// 查找节点个数
	n := 0
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		n = max(n, max(u, v))
	}

	// 定义魔改的邻接表，用 map 是为了之后的查找方便
	to := make([]map[int]struct{}, n+1)
	for i := range to {
		to[i] = make(map[int]struct{})
	}

	// 统计入度
	in := make([]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		in[v]++
		to[u][v] = struct{}{}
	}

	// 用作统计深度遍历过的元素
	var visited map[int]struct{}
	// 用深度遍历来判断是否能通过一个节点递归完所有节点
	var dfs func(root int)
	dfs = func(root int) {
		// 记录当当前节点
		visited[root] = struct{}{}
		// 通过邻接表递归相邻节点
		for node := range to[root] {
			// 如果已经访问过了就跳过
			if _, ok := visited[node]; ok {
				continue
			}
			dfs(node)
		}
	}
	// 辅助检查当前的邻接表能否构成树
	checkValid := func() bool {
		// 定义一个 root，因为有效节点编号从 1 开始
		root := 0
		// 找一个入度为 0 的节点充当根节点
		for i, v := range in {
			if v == 0 {
				root = i
			}
			// 如果一个点入度大于 1， 必定构成环
			if v > 1 {
				return false
			}
		}

		// 如果未找到 root
		if root == 0 {
			return false
		}

		// 初始化统计数组，统计 dfs 能访问到元素（每次判断都要重新初始化）
		visited = make(map[int]struct{})
		// 通过 dfs 从 root 点开始遍历，查看是否能遍历到所有元素
		dfs(root)
		return len(visited) == n
	}

	// 因为要返回最后的边，所以从后往前
	for i := len(edges) - 1; i >= 0; i-- {
		// 尝试删掉当前这条边能否构成树
		u, v := edges[i][0], edges[i][1]
		in[v]--
		delete(to[u], v)

		// 检查当前是否构成树
		if checkValid() {
			return edges[i]
		}

		// 还原现场
		in[v]++
		to[u][v] = struct{}{}
	}
	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
