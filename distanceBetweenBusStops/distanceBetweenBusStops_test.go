package distanceBetweenBusStops

import "testing"

func Test_distanceBetweenBusStops(t *testing.T) {
	var asserts = [][]int{
		{1, 2, 3, 4},
		{0, 1, 1},
		{1, 2, 3, 4},
		{0, 2, 3},
		{1, 2, 3, 4},
		{0, 3, 4},
		{8, 11, 6, 7, 10, 11, 2},
		{0, 3, 25},
		{14, 13, 4, 7, 10, 17, 8, 3, 2, 13},
		{2, 9, 40},
	}

	for i := 0; i < len(asserts); i += 2 {
		res := distanceBetweenBusStops(asserts[i], asserts[i+1][0], asserts[i+1][1])
		if res != asserts[i+1][2] {
			t.Fatalf("input: %v,%v,%v,want:%v, get:%v", asserts[i], asserts[i+1][0], asserts[i+1][1], asserts[i+1][2], res)
		}
	}
}
