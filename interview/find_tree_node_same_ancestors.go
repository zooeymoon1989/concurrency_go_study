package interview

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func FindTreeNodeSameAncestors(root *TreeNode , left , right *TreeNode ) *TreeNode{
	if root == nil || root.val == left.val || root.val == right.val{
		return root
	}

	leftNode := FindTreeNodeSameAncestors(root.left , left , right)
	rightNode := FindTreeNodeSameAncestors(root.right , left , right)

	if leftNode == nil && rightNode == nil {
		return nil
	}else if rightNode == nil {
		return leftNode
	}else if leftNode == nil{
		return rightNode
	}

	return root

}
