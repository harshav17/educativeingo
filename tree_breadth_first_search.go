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

func findLevelAverages(root *treeNode) []float32 {
	var results []float32
	queue := []*treeNode{root}

	for len(queue) > 0 {
		lvlSum := 0
		levLen := len(queue)

		for i := 0; i < levLen; i++ {
			top := queue[0]
			lvlSum += top.val

			if top.left != nil {
				queue = append(queue, top.left)
			}
			if top.right != nil {
				queue = append(queue, top.right)
			}
			queue = queue[1:]
		}

		lvlAvg := float32(lvlSum) / float32(levLen)
		results = append(results, lvlAvg)
	}

	return results
}

func findMinimumDepth(root *treeNode) int {
	queue := []*treeNode{root}
	minDepth := 0
	for len(queue) > 0 {
		minDepth++
		lvlLen := len(queue)
		for i := 0; i < lvlLen; i++ {
			top := queue[0]

			if top.left == nil || top.right == nil {
				return minDepth
			}

			if top.left != nil {
				queue = append(queue, top.left)
			}
			if top.right != nil {
				queue = append(queue, top.right)
			}
			queue = queue[1:]
		}
	}
	return -1
}

func findSuccessor(root *treeNode, key int) *treeNode {
	queue := []*treeNode{root}
	for len(queue) > 0 {
		top := queue[0]

		if top.left != nil {
			queue = append(queue, top.left)
		}
		if top.right != nil {
			queue = append(queue, top.right)
		}

		queue = queue[1:]

		if key == top.val {
			break
		}
	}
	return queue[0]
}
