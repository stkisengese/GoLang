package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root} // to store level order traversal
	level := 0
	for queue != nil {
		level++
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[len(queue):] // remove processed nodes from th equeue
	}
	return level
}
