package tree

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree Node

func NewBinaryTree(data int) *BinaryTree {
	return &BinaryTree{
		Data: data,
		Left: nil,
		Right: nil,
	}
}


