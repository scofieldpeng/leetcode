package addTwoNumer

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumer(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{Next: nil}
	first := &ListNode{Next: nil}
	prevExceedTen := false
	i := 0
	for {
		current := &ListNode{Next: nil}
		if prevExceedTen {
			current.Val = 1
			prevExceedTen = false
		}
		current.Val = current.Val + l1.Val + l2.Val
		if current.Val >= 10 {
			current.Val = current.Val - 10
			prevExceedTen = true
		}
		// 链的操作
		prev.Next = current
		prev = current
		// 把头记录下来
		if i == 0 {
			first = prev
			// 这里主要是防止只有一位的时候，next不为nil，最后变成了一个无限循环链
			prev.Next = nil
		}
		i++

		// 判断是否还有链表
		// 如果最后一次计算的时候超过10,那么再计算一次
		if l1.Next == nil && l2.Next == nil && !prevExceedTen {
			break
		}
		// 这里加上判断是为了l1,l2长度链表不等时依然可以运行
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}
		if l2.Next == nil {
			l2.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return first
}
