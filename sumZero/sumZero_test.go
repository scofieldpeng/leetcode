package sumZero

import "testing"

func Test_sumZero(t *testing.T) {
	inputs := []int{
		5, 3, 1, 4, 9, 80, 123, 592, 29,
	}
	for _, input := range inputs {
		res := sumZero(input)
		if !isSumZero(res) {
			t.Fatalf("\ninput: %v\n,get: %v\n", input, res)
		}
	}
}

func isSumZero(data []int) bool {
	sum := 0
	for _, v := range data {
		sum += v
	}

	return sum == 0
}
