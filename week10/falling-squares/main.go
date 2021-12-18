package main

import "sort"

type SegmentTree struct {
	max    []int
	change []int
	update []bool
}

// NewSegmentTree  创建一颗线段树。
func NewSegmentTree(size int) *SegmentTree {
	// 因为下标从一开始，所以 + 1
	l := 4 * (size + 1)
	return &SegmentTree{
		max:    make([]int, l),
		change: make([]int, l),
		update: make([]bool, l),
	}
}

// pushUp 计算当前节点的 max。
func (s *SegmentTree) pushUp(p int) {
	s.max[p] = max(s.max[p*2], s.max[p*2+1])
}

// pushDown 更新子节点。
func (s *SegmentTree) pushDown(p int) {
	// 如果当前节点无需更新
	if !s.update[p] {
		return
	}

	// 将当前节点的属性更新到子节点
	s.update[p*2] = true
	s.update[p*2+1] = true
	s.change[p*2] = s.change[p]
	s.change[p*2+1] = s.change[p]
	s.max[p*2] = s.change[p]
	s.max[p*2+1] = s.change[p]
	s.update[p] = false
}

// Update 更新
func (s *SegmentTree) Update(L, R, H, l, r, p int) {
	if L <= l && r <= R {
		s.update[p] = true
		s.change[p] = H
		s.max[p] = H
		return
	}

	mid := (l + r) / 2
	// 更新子节点
	s.pushDown(p)
	if L <= mid {
		s.Update(L, R, H, l, mid, p*2)
	}
	if R > mid {
		s.Update(L, R, H, mid+1, r, p*2+1)
	}
	// 更新当前节点
	s.pushUp(p)
}

// Query 查询线段树。
func (s *SegmentTree) Query(L, R, l, r, p int) int {
	// 如果当前的左右区间正好在合法范围内
	if L <= l && r <= R {
		return s.max[p]
	}
	mid := (l + r) / 2
	s.pushDown(p)
	// 分别计算左右子树的高度
	lh, rh := 0, 0
	if L <= mid {
		lh = s.Query(L, R, l, mid, p*2)
	}
	if R > mid {
		rh = s.Query(L, R, mid+1, r, p*2+1)
	}
	return max(lh, rh)
}

// index 坐标压缩。
func index(positions [][]int) map[int]int {
	// 坐标去重
	mp := make(map[int]bool)
	for _, info := range positions {
		left, sideLen := info[0], info[1]
		mp[left], mp[left+sideLen-1] = true, true
	}

	// 坐标排序
	tmp := make([]int, 0, len(mp))
	for key := range mp {
		tmp = append(tmp, key)
	}
	sort.Ints(tmp)

	// 坐标压缩
	pos, cnt := make(map[int]int), 0
	for _, v := range tmp {
		cnt++
		pos[v] = cnt
	}
	return pos
}

// 时间复杂度 O(nlogn)，n 为 positions 的长度，也就是线段树的时间复杂度。
// 空间复杂度为 O(n)，因为需要借助 n 的常数倍的辅助空间辅助计算，所以空间复杂度为 O(n)。
func fallingSquares(positions [][]int) (ans []int) {
	mp := index(positions)
	n := len(mp)
	seg := NewSegmentTree(n)
	maxH := 0
	// 计算每个正方形落下后的最大高度
	for _, info := range positions {
		left, sideLen := info[0], info[1]
		l, r := mp[left], mp[left+sideLen-1]
		h := seg.Query(l, r, 1, n, 1) + sideLen
		maxH = max(maxH, h)
		ans = append(ans, maxH)
		seg.Update(l, r, h, 1, n, 1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
