package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	switch root != nil {
	case elem == root.Data:
		return root
	case elem < root.Data:
		return BTreeSearchItem(root.Left, elem)
	default:
		return BTreeSearchItem(root.Right, elem)
	}
}
