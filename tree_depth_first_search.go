package educativeingo

//HasPath that add up to a given sum
func HasPath(root *treeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.val == sum && root.left == nil && root.right == nil {
		return true
	}

	return HasPath(root.left, sum-root.val) || HasPath(root.right, sum-root.val)
}

//FindAllPaths that add up to the given sum
func FindAllPaths(root *treeNode, sum int) [][]int {
	currPath := make([]int, 10)
	allPaths := make([][]int, 10)
	findPath(root, sum, currPath, allPaths)
	return [][]int{}
}

func findPath(root *treeNode, sum int, currPath []int, allPaths [][]int) {
	if root == nil {
		return
	}

	currPath = append(currPath, root.val)

	if root.val == sum && root.left == nil && root.right == nil {
		allPaths = append(allPaths, currPath)
	} else {

	}
}
