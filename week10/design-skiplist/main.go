package main

import "math/rand"

const maxLevel = 11 // maxLevel 索引层最大层高

// node 跳表节点。
type node struct {
	data     int             // 节点数据
	forwards [maxLevel]*node // 每层前进的指针，比如 forwards[3] 表示当前节点在第三级索引上的下一个节点的内容
	maxLevel int             // 层高
}

// Skiplist 跳表。
type Skiplist struct {
	head   *node // 带头链表
	levels int   // 当前层数
	length int   // 链表长度
}

// Constructor 初始化跳表。
func Constructor() Skiplist {
	return Skiplist{
		head:   &node{},
		levels: 1,
		length: 0,
	}
}

// Search 查找节点是否存在。
func (this *Skiplist) Search(target int) bool {
	if this.length == 0 {
		return false
	}

	cur := this.head
	// 从最高层往下依次遍历查找元素
	for i := this.levels - 1; i >= 0; i-- {
		for cur.forwards[i] != nil && cur.forwards[i].data < target {
			// 将当前元素替换成第 i 层的下一个元素
			// 因为 target 大于当前元素，所以必定在位置上排列在当前元素之后
			cur = cur.forwards[i]
		}
	}

	// 判断元素是否在跳表中
	return cur.forwards[0] != nil && cur.forwards[0].data == target
}

// Add 往跳表添加元素。
func (this *Skiplist) Add(num int) {
	// 随机生成一个层数，将到该层数为止的索引层以及原链表都插入新节点
	level := this.randomLevel()
	// 构建新节点
	newNode := &node{
		data:     num,
		maxLevel: level,
		forwards: [maxLevel]*node{},
	}
	// 存放待插入节点的前一个节点的指针信息切片
	// 因为链表的插入需要记录带插入节点的上一个节点信息，并更改其 next 指针
	update := make([]*node, level)
	for i := range update {
		update[i] = newNode
	}

	cur := this.head
	// 从最大层开始查找，找到前一节点，移动到下层再开始查找
	for i := this.levels - 1; i >= 0; i-- {
		// 在当前索引层查找最后一个小于 num 的节点
		for cur.forwards[i] != nil && cur.forwards[i].data < num {
			cur = cur.forwards[i]
		}

		// 找到当前索引层最靠近带插入节点位置并且值小于待插入节点的节点信息
		if i < level {
			update[i] = cur
		}
	}

	// 将待插入节点信息更新至随机生成的索引层之前的所有索引层中
	for i := 0; i < level; i++ {
		// 每次重新生成一个指针地址（tmp），防止所有元素粘连
		tmp := newNode
		tmp.forwards[i] = update[i].forwards[i]
		update[i].forwards[i] = tmp
	}

	// 更新层高
	if this.levels < level {
		this.levels = level
	}
	// 更新跳表原始链表长度
	this.length++
}

// Erase 抹除跳表节点。
func (this *Skiplist) Erase(num int) (exist bool) {
	// 保存所有待删除节点的前一个节点的指针信息（因为链表删除元素需要修改待删除节点的前一个节点的 next 指针信息）
	// 和插入类似的处理思路，将所有待更新指针统一处理
	update := make([]*node, this.levels)

	cur := this.head
	// 从最高层开始查找待删除节点信息
	for i := this.levels - 1; i >= 0; i-- {
		for cur.forwards[i] != nil && cur.forwards[i].data < num {
			cur = cur.forwards[i]
		}
		update[i] = cur
	}

	// 目标元素确实在存在
	if cur.forwards[0] != nil && cur.forwards[0].data == num {
		// 通过 update 切片中存储的待删除节点的上一个节点的指针信息，判断索引层中是否包含待删除节点信息，如果有，更新索引层
		for i := this.levels - 1; i >= 0; i-- {
			if update[i].forwards[i] != nil && update[i].forwards[i].data == num {
				update[i].forwards[i] = update[i].forwards[i].forwards[i]
			}
		}
		// 更新跳表原始链表长度
		this.length--
		exist = true
	}

	// 清理层高
	for this.levels > 1 && this.head.forwards[this.levels] == nil {
		this.levels--
	}
	return
}

// randomLevel 随机生成一个层数，用作插入节点。
func (this *Skiplist) randomLevel() int {
	level := 1
	for rand.Int()&1 == 1 && level < maxLevel-1 {
		level++
	}
	return level
}
