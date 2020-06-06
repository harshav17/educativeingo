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
	var currPath []int
	var allPaths [][]int
	findPath(root, sum, currPath, &allPaths)
	return allPaths
}

func findPath(root *treeNode, sum int, currPath []int, allPaths *[][]int) {
	if root == nil {
		return
	}

	currPath = append(currPath, root.val)

	if root.val == sum && root.left == nil && root.right == nil {
		*allPaths = append(*allPaths, currPath)
	} else {
		findPath(root.left, sum-root.val, currPath, allPaths)
		findPath(root.right, sum-root.val, currPath, allPaths)
	}

	currPath = currPath[:len(currPath)-1]
}
