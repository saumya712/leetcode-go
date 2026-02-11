func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		next := current.Next      // 1. save next
		current.Next = prev       // 2. reverse pointer
		prev = current            // 3. move prev
		current = next            // 4. move current
	}

	return prev
}
