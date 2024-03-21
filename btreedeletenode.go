package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	if elem == root.Data {
		return root
	} else if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return BTreeSearchItem(root.Right, elem)
	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	currentNode := root
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left != nil && root.Right != nil {
			min := BTreeMin(root.Right)
			root.Data = min.Data
			root.Right = BTreeDeleteNode(root.Right, min)
		} else {
			if root.Left == nil && root.Right == nil {
				root = nil
			} else if root.Right != nil {
				root = root.Right
			} else {
				root = root.Left
			}
		}
	}
	return root
}
