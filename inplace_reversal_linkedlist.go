package educativeingo

func reverseASubList(root *node, p int, q int) *node {
	var prev *node
	current := root
	for i := 0; root != nil && i < p-1; i++ {
		prev = current
		current = current.next
	}

	//i is pointing to p, reverse it until q
	lastNodeOfFirst := prev
	lastNodeOfSubList := current
	for i := 0; current != nil && i < q-p+1; i++ {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	if lastNodeOfFirst != nil {
		lastNodeOfFirst.next = prev
	} else {
		//this means prev == 1
		root = prev
	}

	lastNodeOfSubList.next = current

	return root
}

func reverseEveryKSublist(root *node, k int) *node {
	var prev *node
	current := root

	for current != nil {
		lastNodeOfFirst := prev
		lastNodeOfSubList := current

		for i := 0; current != nil && i < k; i++ {
			next := current.next
			current.next = prev
			prev = current
			current = next
		}

		if lastNodeOfFirst != nil {
			lastNodeOfFirst.next = prev
		} else {
			root = prev
		}
		lastNodeOfSubList.next = current
		prev = lastNodeOfSubList
	}

	return root
}

func reverseAlternatingKSublist(root *node, k int) *node {
	var prev *node
	current := root

	for current != nil {
		lastNodeOfFirst := prev
		lastNodeOfSubList := current

		for i := 0; current != nil && i < k; i++ {
			next := current.next
			current.next = prev
			prev = current
			current = next
		}

		if lastNodeOfFirst != nil {
			lastNodeOfFirst.next = prev
		} else {
			root = prev
		}

		lastNodeOfSubList.next = current
		for j := 0; current != nil && j < k; j++ {
			prev = current
			current = current.next
		}
	}
	return root
}

func rotateALinkedList(root *node, rotations int) *node {
	lastNode := root
	listLen := 1
	for lastNode.next != nil {
		lastNode = lastNode.next
		listLen++
	}
	lastNode.next = root
	rotations %= listLen
	skipLen := listLen - rotations
	rotNode := root
	for i := 0; i < skipLen-1; i++ {
		rotNode = rotNode.next
	}

	root = rotNode.next
	rotNode.next = nil
	return root
}
