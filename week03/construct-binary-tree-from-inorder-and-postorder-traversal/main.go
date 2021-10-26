package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度 O(n), n 为每个树中节点个数，因为需要遍历所有节点，所以时间复杂度为 O(n)。
// 空间复杂度 O(n), n 为每个树中节点个数，因为树高必是小于等于 n 的，而递归深度为树高，所以空间复杂度为 O(n)。
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 因为 postorder 长度和 inorder 相同，所以只要判断一个就好了
	if len(inorder) == 0 {
		return nil
	}

	// 先序的第一个节点为根节点
	root := postorder[len(postorder)-1]
	// 定义一个 idx，用作找到根节点在中序中的位置
	idx := 0
	for i := range inorder {
		if inorder[i] == root {
			idx = i
			break
		}
	}
	// 定义一颗子树，值为根节点
	tree := &TreeNode{
		Val: root,
	}
	// 通过根节点拆分中序遍历和后序遍历，找到左右子树的中序遍历和后序遍历
	// 递归实现子树
	tree.Left = buildTree(inorder[:idx], postorder[:idx])
	tree.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return tree
}
