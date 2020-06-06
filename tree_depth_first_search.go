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

//FindSumOfPathNumbers of all the numbers represented by all paths
func FindSumOfPathNumbers(root *treeNode) int {
	return findSumPath(root, 0)
}

func findSumPath(root *treeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = sum*10 + root.val

	if root.left == nil && root.right == nil {
		return sum
	}

	return findSumPath(root.left, sum) + findSumPath(root.right, sum)
}

//FindPathWGivenSequence finds the path with a given sequence
func FindPathWGivenSequence(root *treeNode, seq []int) bool {
	if root == nil {
		return false
	}

	currItem := seq[0]
	seq = seq[1:]

	if currItem != root.val {
		return false
	}

	if len(seq) == 0 && root.left == nil && root.right == nil {
		return true
	}

	return FindPathWGivenSequence(root.left, seq) || FindPathWGivenSequence(root.right, seq)
}

//CountPathsOfSum counts all the paths that adds up to a particular sum
func CountPathsOfSum(root *treeNode, sum int) int {
	var currPath []int
	return findPathsOfSum(root, sum, currPath)
}

func findPathsOfSum(root *treeNode, sum int, currPath []int) int {
	if root == nil {
		return 0
	}

	currPath = append(currPath, root.val)

	pathCount, pathSum := 0, 0

	for i := len(currPath) - 1; i >= 0; i-- {
		pathSum += currPath[i]

		if pathSum == sum {
			pathCount++
		}
	}

	pathCount += findPathsOfSum(root.left, sum, currPath)
	pathCount += findPathsOfSum(root.right, sum, currPath)

	currPath = currPath[1:]

	return pathCount
}

func findDiameter(root *treeNode) int {
	return 0
}
