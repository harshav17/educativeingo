package educativeingo

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func levelOrderTraversal(root *treeNode) [][]int {
	queue := []*treeNode{root}
	var results [][]int

	for len(queue) > 0 {
		var level []int
		levLen := len(queue)
		for i := 0; i < levLen; i++ {
			top := queue[0]
			level = append(level, top.val)
			if top.left != nil {
				queue = append(queue, top.left)
			}
			if top.right != nil {
				queue = append(queue, top.right)
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
	var results [][]int

	for len(queue) > 0 {
		var level []int
		levLen := len(queue)
		for i := 0; i < levLen; i++ {
			top := queue[0]
			level = append(level, top.val)
			if top.left != nil {
				queue = append(queue, top.left)
			}
			if top.right != nil {
				queue = append(queue, top.right)
			}
			queue = queue[1:]
		}
		if level != nil {
			results = append([][]int{level}, results...)
		}
	}

	return results
}

func zigzagLevelOrderTrav(root *treeNode) [][]int {
	queue := []*treeNode{root}
	var results [][]int
	flip := true

	for len(queue) > 0 {
		var level []int
		levLen := len(queue)
		for i := 0; i < levLen; i++ {
			top := queue[0]

			if flip {
				level = append(level, top.val)
			} else {
				level = append([]int{top.val}, level...)
			}

			if top.left != nil {
				queue = append(queue, top.left)
			}
			if top.right != nil {
				queue = append(queue, top.right)
			}
			queue = queue[1:]
		}
		if level != nil {
			results = append(results, level)
		}
		flip = !flip
	}
	return results
}
