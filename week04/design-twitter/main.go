package main

// 整体空间复杂度为 O(n),n 为出现的用户的数量，主要内存消耗在保存关注信息和堆文信息的无序集合上，二者都和涉及的人数有关。
// 具体时间复杂度见函数注释。

// ListNode 定义一个节点。
type ListNode struct {
	TimeIdx int       // 时间索引
	TweetId int       // 推文 ID
	Next    *ListNode // 下一个节点地址
	Pre     *ListNode // 上一个节点地址
}

// Deque 定义一个双端队列，打算每个队列除去队头队尾外保存十条推文。
type Deque struct {
	Head *ListNode // 头节点
	Tail *ListNode // 尾节点
	Size int       // 长度
}

// Heap 自定义实现一个大顶堆，时间大的为堆顶。
type Heap struct {
	cache []*ListNode // 用数组来存放堆内元素
}

// NewHeap 创建一个堆。
func NewHeap() *Heap {
	return &Heap{
		cache: make([]*ListNode, 1), // 第一个元素空着（为了方便之后对父节点，子节点的计算）
	}
}

// Push 往堆中 push 一个节点。时间复杂度 logn。
func (h *Heap) Push(node *ListNode) {
	// 保存节点
	h.cache = append(h.cache, node)
	// 堆化
	h.heapifyUp()
}

// heapifyUp 向上堆化。时间复杂度 logn。
func (h *Heap) heapifyUp() {
	p := len(h.cache) - 1
	// 因为数组中存放堆元素是从下标 1 开始的
	for p > 1 {
		// 父节点下标索引
		fa := p / 2
		// 此时已经满足堆性质
		if h.cache[fa].TimeIdx >= h.cache[p].TimeIdx {
			break
		}

		// 因为父节点的值小于子节点，所以需要交换父子节点
		h.cache[fa], h.cache[p] = h.cache[p], h.cache[fa]
		// 继续向上堆化，直至满足堆性质
		p = fa
	}
}

// Pop 从堆顶取出时间最近的节点。时间复杂度 logn。
func (h *Heap) Pop() *ListNode {
	// 取堆顶
	top := h.cache[1]
	// 将堆顶和堆尾元素互换，然后删除堆尾
	h.cache[1], h.cache[len(h.cache)-1] = h.cache[len(h.cache)-1], h.cache[1]
	h.cache = h.cache[:len(h.cache)-1]

	// 向下堆化
	h.heapifyDown()
	return top
}

// heapifyDown 向下堆化。时间复杂度 logn。
func (h *Heap) heapifyDown() {
	p := 1
	// 2*p 为左孩子下标索引
	for 2*p < len(h.cache) {
		leftChild := 2 * p
		rightChild := 2*p + 1
		child := leftChild
		// 如果存在右节点，并且右节点的时间更新，则准备父节点和右孩子的交换
		if rightChild < len(h.cache) &&
			h.cache[rightChild].TimeIdx > h.cache[child].TimeIdx {
			child = rightChild
		}

		// 如果此时已经满足堆性质，即父节点的时间新于所有孩子的时间
		if h.cache[p].TimeIdx >= h.cache[child].TimeIdx {
			break
		}

		// 交换父子节点
		h.cache[p], h.cache[child] = h.cache[child], h.cache[p]
		// 继续向下堆化
		p = child
	}
}

// Size 取堆长度。
func (h *Heap) Size() int {
	return len(h.cache) - 1
}

type Twitter struct {
	Time    int                      // 时间索引，通过对其累加判断推文时间顺序
	Follows map[int]map[int]struct{} // 关注列表，key：用户 ID，val：关注列表，用集合方便查找
	Tweets  map[int]*Deque           // 推文集合，key：用户 ID，val：对应用户推文，双端队列保存
}

func Constructor() Twitter {
	return Twitter{
		Time:    0,
		Follows: make(map[int]map[int]struct{}),
		Tweets:  make(map[int]*Deque),
	}
}

// PostTweet 发送推文。时间复杂度 O(1)，只是通过双端队列和无序集合保存了用户 id 和推文 id 的关系，期间所有设计逻辑都是 O(1) 复杂度。
func (this *Twitter) PostTweet(userId int, tweetId int) {
	// 如果双端队列未初始化
	if _, ok := this.Tweets[userId]; !ok {
		deque := &Deque{
			Head: &ListNode{},
			Tail: &ListNode{},
			Size: 0,
		}
		deque.Head.Next = deque.Tail
		deque.Tail.Pre = deque.Head
		this.Tweets[userId] = deque
	}

	// 如果队列已满，因为题意最多只需要保存每个用户的十条推文，多了无用。
	if this.Tweets[userId].Size == 10 {
		// 移除堆尾
		this.Tweets[userId].Tail.Pre.Pre.Next = this.Tweets[userId].Tail
		this.Tweets[userId].Tail.Pre = this.Tweets[userId].Tail.Pre.Pre
		this.Tweets[userId].Size--
	}

	// 更新时间信息
	this.Time++
	// 打包一个节点
	newNode := &ListNode{
		TimeIdx: this.Time,
		TweetId: tweetId,
	}
	// 更新双端队列
	this.Tweets[userId].Head.Next.Pre = newNode
	newNode.Next = this.Tweets[userId].Head.Next
	this.Tweets[userId].Head.Next = newNode
	// 更新计数
	this.Tweets[userId].Size++
}

// GetNewsFeed 获取单用户相关所有推文 id。
// 时间复杂度 O(1)，因为涉及操作都是常数复杂度。
func (this *Twitter) GetNewsFeed(userId int) []int {
	//  用一个双端队列数组保存该用户所有关注的人的推文，其长度为关注列表长度加上自身（所以加 1）
	lists := make([]*Deque, 0, len(this.Follows[userId])+1)
	for followeeId := range this.Follows[userId] {
		if _, ok := this.Tweets[followeeId]; ok {
			lists = append(lists, this.Tweets[followeeId])
		}
	}
	// 再额外保存子集的推文信息
	if _, ok := this.Tweets[userId]; ok {
		lists = append(lists, this.Tweets[userId])
	}

	// 如果没有任何推文，结束函数
	if len(lists) == 0 {
		return nil
	}

	// 因为只取十条
	ans := make([]int, 0, 10)
	// 初始化一个堆
	h := NewHeap()
	// 往堆中添加元素，每次复杂度 logn，n 最大 10，所以常数时间复杂度操作
	for _, q := range lists {
		h.Push(q.Head.Next)
	}
	// 判断堆中是否还有元素，并且推文数量不超过 10 条
	for h.Size() > 0 && len(ans) < 10 {
		// 取堆顶，也就是最新的推文,取堆顶复杂度 O(1)，但是之后涉及向下堆化，所以复杂度 logn，n 不超过 10
		top := h.Pop()
		// 更新结果
		ans = append(ans, top.TweetId)
		// 如果堆顶元素之后还有节点，并且不是 tail 节点（因为 tail 节点的 timeIdx 是 0）
		if top.Next != nil && top.Next.TimeIdx != 0 {
			// 继续往堆添加元素，复杂度 logn，n 不超过 10
			h.Push(top.Next)
		}
	}
	return ans
}

// Follow 关注。因为只涉及哈希操作，所以时间复杂度为 O(1)。
func (this *Twitter) Follow(followerId int, followeeId int) {
	// 初始化无序集合
	if _, ok := this.Follows[followerId]; !ok {
		this.Follows[followerId] = make(map[int]struct{})
	}

	// 用无序集合保存关注和被关注人的信息，方便查找操作
	this.Follows[followerId][followeeId] = struct{}{}
}

// Unfollow 取关。因为只涉及哈希操作，所以时间复杂度为 O(1)（不考虑 map 恰好缩容）。
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.Follows[followerId]; !ok {
		return
	}
	// 更新无序集合。
	delete(this.Follows[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
