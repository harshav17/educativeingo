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
