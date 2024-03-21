package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0 // Empty tree has 0 levels
	}
	queue := []*TreeNode{root} // Queue to store nodes for level-order traversal
	level := 0                 // Current level counter
	for len(queue) > 0 {
		level++
		for i := 0; i < len(queue); i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[len(queue):]
	}
	return level
}
