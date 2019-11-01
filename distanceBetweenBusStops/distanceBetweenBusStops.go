// source: https://leetcode.com/problems/distance-between-bus-stops/
// solution: https://algs.home.pjf.im/distance-between-bus-stops/
package distanceBetweenBusStops

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	var (
		res1    = 0
		res2    = 0
		current = 0
		length  = len(distance)
	)

	if start == destination {
		return 0
	}

	// 确保start比destination小
	if start > destination {
		start, destination = destination, start
	}

	// 正向遍历和反向遍历，然后看两个值那个值最少

	current = start
	for ; current != destination; current = (current + 1) % length {
		res1 += distance[current]
	}

	current = destination
	for ; current != start; current = (current + 1) % length {
		res2 += distance[current]
	}

	if res1 <= res2 {
		return res1
	}

	return res2
}
