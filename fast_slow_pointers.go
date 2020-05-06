package educativeingo

type node struct {
	value int
	next  *node
}

//the hare and the tortoise

func hasCycle(root *node) bool {
	slow := root
	fast := root
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if slow == fast {
			return true
		}
	}
	return false
}

func startOfCycle(root *node) *node {
	slow := root
	fast := root
	length := 0
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if slow == fast {
			length = calculateLength(slow)
			break
		}
	}

	pointer1 := root
	pointer2 := root
	for length > 0 {
		pointer2 = pointer2.next
		length--
	}

	for pointer1 != pointer2 {
		pointer1 = pointer1.next
		pointer2 = pointer2.next
	}
	return pointer1
}

func calculateLength(start *node) int {
	current := start
	count := 1
	for start.next != current {
		start = start.next
		count++
	}
	return count
}

func isHappyNumber(num int) bool {
	slow := findSquareSum(num)
	fast := findSquareSum(slow)
	for slow != fast {
		slow = findSquareSum(slow)
		fast = findSquareSum(findSquareSum(fast))
	}
	return slow == 1
}

func findSquareSum(num int) int {
	sum := 0
	digit := 0
	for num > 0 {
		digit = num % 10
		sum += digit * digit
		num /= 10
	}
	return sum
}