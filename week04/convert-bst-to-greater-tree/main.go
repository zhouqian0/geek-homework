package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度 O(n)。n 为树节点个数，因为需要遍历每个节点累加和，而节点个数为 n，所以时间复杂度为 O(n)。
// 空间复杂度 O(n)。因为采用了 dfs 深度优先遍历，所以空间消耗和栈深度相关，而栈深度和树深相关，树深最坏为 n，最好 logn，按照最坏情况，空间复杂度为 O(n)。
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	// 使用逆中序遍历，在遍历的同时计算前缀和，并将和赋值给正在进行遍历的节点
	// 因为二叉搜索树的中序遍历结果为递增数组，而题目要求的结果恰好等于这个递增数组逆序之后的前缀和
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}

	dfs(root)
	return root
}
