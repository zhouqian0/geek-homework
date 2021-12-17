package main

type segmentTree struct {
	max    []int
	change []int
	update []bool
}

func NewsegmentTree(size int) *segmentTree {
	N := size + 1
	return &segmentTree{
		max:    make([]int, N<<2),
		change: make([]int, N<<2),
		update: make([]bool, N<<2),
	}
}

func (seg *segmentTree) pushUp(rt int) {
	seg.max[rt] = Max(seg.max[rt<<1], seg.max[rt<<1|1])
}

// ln表示左子树元素结点个数，rn表示右子树结点个数
func (seg *segmentTree) pushDown(rt, ln, rn int) {
	if !seg.update[rt] {
		return
	}

	seg.update[rt<<1] = true
	seg.update[rt<<1|1] = true
	seg.change[rt<<1] = seg.change[rt]
	seg.change[rt<<1|1] = seg.change[rt]
	seg.max[rt<<1] = seg.change[rt]
	seg.max[rt<<1|1] = seg.change[rt]
	seg.update[rt] = false
}

func (seg *segmentTree) Update(L, R, C, l, r, rt int) {
	if L <= l && r <= R {
		seg.update[rt] = true
		seg.change[rt] = C
		seg.max[rt] = C
		return
	}

	mid := (l + r) >> 1
	seg.pushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		seg.Update(L, R, C, l, mid, rt<<1)
	}
	if R > mid {
		seg.Update(L, R, C, mid+1, r, rt<<1|1)
	}
	seg.pushUp(rt)
}

func (seg *segmentTree) query(L, R, l, r, rt int) int {
	if L <= l && r <= R {
		return seg.max[rt]
	}
	mid := (l + r) >> 1
	seg.pushDown(rt, mid-l+1, r-mid)
	left := 0
	right := 0
	if L <= mid {
		left = seg.query(L, R, l, mid, rt<<1)
	}
	if R > mid {
		right = seg.query(L, R, mid+1, r, rt<<1|1)
	}
	return Max(left, right)
}

func index(positions [][]int) map[int]int {
	pos := make(map[int]bool)
	for _, arr := range positions {
		pos[arr[0]] = true
		pos[arr[0]+arr[1]-1] = true
	}

	tmp := make([]int, len(pos))
	index := 0
	for key, _ := range pos {
		tmp[index] = key
		index++
	}
	sort.Ints(tmp)

	mp := make(map[int]int)
	count := 0
	for _, value := range tmp {
		count++
		mp[value] = count
	}
	return mp
}

func fallingSquares(positions [][]int) []int {
	mp := index(positions)
	N := len(mp)
	seg := NewsegmentTree(N)
	max := 0
	res := make([]int, 0)
	// 每落一个正方形，收集一下，所有东西组成的图像，最高高度是什么
	for _, arr := range positions {
		L := mp[arr[0]]
		R := mp[arr[0]+arr[1]-1]
		height := seg.query(L, R, 1, N, 1) + arr[1]
		max = Max(max, height)
		res = append(res, max)
		seg.Update(L, R, height, 1, N, 1)
	}
	return res
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
