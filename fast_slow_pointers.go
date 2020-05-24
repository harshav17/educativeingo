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

func findMiddle(head *node) *node {
	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func reverse(head *node) *node {
	var prev *node
	for head != nil {
		next := head.next
		head.next = prev
		prev = head
		head = next
	}
	return prev
}

func isPalindrome(head *node) bool {
	middle := findMiddle(head)
	headSecondHalf := reverse(middle)
	for head != nil && headSecondHalf != nil {
		if head.value != headSecondHalf.value {
			break
		}
		head = head.next
		headSecondHalf = headSecondHalf.next
	}
	if head == nil || headSecondHalf == nil {
		return true
	}
	return false
}

func rearrangeLinkedList(head *node) {
	//find middle
	middle := findMiddle(head)

	//reverse second half
	headSecondHalf := reverse(middle)
	headFirstHalf := head

	for headFirstHalf != nil && headSecondHalf != nil {
		temp := headFirstHalf.next
		headFirstHalf.next = headSecondHalf
		headFirstHalf = temp

		temp = headSecondHalf.next
		headSecondHalf.next = headFirstHalf
		headSecondHalf = temp
	}

	if headFirstHalf != nil {
		headFirstHalf.next = nil
	}
}

func findCycleInCircularArray(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		isFwd := arr[i] > 0
		slow := i
		fast := i

		for {
			slow = findNextIndex(arr, isFwd, slow)
			fast = findNextIndex(arr, isFwd, fast)
			if fast != -1 {
				fast = findNextIndex(arr, isFwd, fast)
			}

			if slow == -1 || fast == -1 || slow == fast {
				break
			}
		}

		if slow != -1 && slow == fast {
			return true
		}

	}
	return false
}

func findNextIndex(arr []int, isFwd bool, currIdx int) int {
	//find direction and check if reversed
	dir := arr[currIdx] >= 0
	if isFwd != dir {
		return -1
	}

	//move to next index, if negative rotate
	nxtIdx := (currIdx + arr[currIdx]) % len(arr)
	if nxtIdx < 0 {
		nxtIdx += len(arr)
	}

	//single cycle
	if nxtIdx == currIdx {
		nxtIdx = -1
	}
	return nxtIdx
}
