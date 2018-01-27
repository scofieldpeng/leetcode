package addTwoNumer

import "testing"

func Test_AddTwoNumber(t *testing.T) {
	l1_1 := &ListNode{Val: 2}
	l1_2 := &ListNode{Val: 4}
	l1_1.Next = l1_2
	l1_3 := &ListNode{Val: 3, Next: nil}
	l1_2.Next = l1_3
	l1_4 := &ListNode{Val: 5, Next: nil}
	l1_3.Next = l1_4

	l2_1 := &ListNode{Val: 5}
	l2_2 := &ListNode{Val: 6}
	l2_1.Next = l2_2
	l2_3 := &ListNode{Val: 4, Next: nil}
	l2_2.Next = l2_3
	l2_4 := &ListNode{Val: 7, Next: nil}
	l2_3.Next = l2_4

	res := AddTwoNumer(l1_1, l2_1)
	if res.Val == 7 {
		if res.Next.Val == 0 {
			if res.Next.Next.Val == 8 {
				if res.Next.Next.Next.Val == 2 {
					t.Log("ok")
				} else {
					t.Error("forth error:", res.Next.Next.Next.Val)
				}
			} else {
				t.Error("third error:", res.Next.Next.Val)
			}
		} else {
			t.Error("second error:", res.Next.Val)
		}
	} else {
		t.Error("first error:", res.Val)
	}
}
