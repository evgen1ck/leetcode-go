package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers
func sumRootToLeaf(root *TreeNode) int {
	return search(root, 0)
}

func search(node *TreeNode, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum = currentSum*2 + node.Val

	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	return search(node.Left, currentSum) + search(node.Right, currentSum)
}
