package educativeingo

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func levelOrderTraversal(root *treeNode) [][]int {
	queue := []*treeNode{root}
	results := [][]int{{root.val}}

	for len(queue) > 0 {
		var level []int
		levLen := len(queue)
		for i := 0; i < levLen; i++ {
			top := queue[0]
			if top.left != nil {
				queue = append(queue, top.left)
				level = append(level, top.left.val)
			}
			if top.right != nil {
				queue = append(queue, top.right)
				level = append(level, top.right.val)
			}
			queue = queue[1:]
		}
		if level != nil {
			results = append(results, level)
		}
	}

	return results
}

func reverseLevelOrderTrav(root *treeNode) [][]int {
	queue := []*treeNode{root}
	results := [][]int{{root.val}}

	for len(queue) > 0 {
		var level []int
		levLen := len(queue)
		for i := 0; i < levLen; i++ {
			top := queue[0]
			if top.left != nil {
				queue = append(queue, top.left)
				level = append(level, top.left.val)
			}
			if top.right != nil {
				queue = append(queue, top.right)
				level = append(level, top.right.val)
			}
			queue = queue[1:]
		}
		if level != nil {
			results = append([][]int{level}, results...)
		}
	}

	return results
}
