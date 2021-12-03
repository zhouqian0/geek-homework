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

// 时间复杂度 O(n)，n 为 edges 长度，因为操作并查集的时间复杂度为O(α(n))，α(n) 为反阿克曼函数，近似常数。代码中使用了一个 for 循环（长度为 n）执行并查集的查找和合并，所以时间复杂度为O(α(n) * n)， 即 O(n)。
// 空间复杂度 O(n)，因为并查集需要一个长度等同 edges 的数组辅助计算，所以空间复杂度在 O(n)。
func findRedundantConnection(edges [][]int) []int {
	// 初始化并查集，节点下标从 1 开始（非 0），所以需要 +1。
	s := Construct(len(edges) + 1)
	// 遍历每个连接
	for _, e := range edges {
		// 如果当前的两个节点连通，就返回当前连接
		if s.Find(e[0]) == s.Find(e[1]) {
			return e
		}
		// 连通两个节点
		s.Join(e[0], e[1])
	}
	return nil
}
