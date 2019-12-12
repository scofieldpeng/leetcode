package findUnsortedSubarray

import "sort"

// sort大法
func findUnsortedSubarray(nums []int) int {
	var data = make([]int, len(nums), len(nums))
	copy(data, nums)
	sort.Ints(data)

	begin := 0
	end := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != data[i] {
			begin = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != data[i] {
			end = i
			break
		}
	}

	if end == begin {
		return 0
	}

	return end - begin + 1
}

// stack法
func findUnsortedSubarray2(nums []int) int {
	var stack = NewStack()
	var (
		leftIndex  = len(nums) - 1
		rightIndex = 0
	)
	// 找出最左侧那个下降的波谷
	for k, v := range nums {
		for ; stack.Length() > 0 && nums[stack.Peek()] > v; {
			leftIndex = min(stack.Pop(), leftIndex)
		}
		stack.Push(k)
	}

	stack.Clear()

	// 找出最右侧的那个上升的波谷
	for i := len(nums) - 1; i >= 0; i-- {
		for ; stack.Length() > 0 && nums[stack.Peek()] < nums[i]; {
			rightIndex = max(stack.Pop(), rightIndex)
		}
		stack.Push(i)
	}

	if leftIndex == rightIndex {
		return 0
	}

	return rightIndex - leftIndex + 1
}

type Stack struct {
	length int
	data   []int
}

func NewStack() *Stack {
	return &Stack{length: 0, data: make([]int, 0),}
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) Peek() int {
	if s.length > 0 {
		return s.data[s.length-1]
	}

	return 0
}

func (s *Stack) Valley() int {
	if s.length > 0 {
		return s.data[0]
	}

	return 0
}

func (s *Stack) Pop() int {
	if s.length < 1 {
		return 0
	}
	var last = s.Peek()
	s.data = s.data[0 : s.length-1]
	s.length--

	return last
}

func (s *Stack) Push(data int) {
	s.length++
	s.data = append(s.data, data)
}

func (s *Stack) Clear() {
	s.length = 0
	s.data = make([]int, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
