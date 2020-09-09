// problem: https://leetcode.com/problems/minimum-time-visiting-all-points/
package minTimeToVisitAllPoints

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	// 如果x1=x2,那么直接移动abs(x1-x2)就是两点之间的最短距离
	// 如果y1=y2,那么直接移动abs(y1-y2)就是两点之间最短的距离
	// x1!=x2 and y1!=y2
	// (x1,y1)先挪动到距离(x2,y2)最近的点，而这个最近的点就是先斜着走到最后一个，然后在横着/竖着走到终点即可
	// 其实提炼出来后，永远都是看两个点横向和纵向那个远，那个就是最短路径
	length := len(points)
	steps := 0
	for i := 0; i < length-1; i++ {
		steps += int(math.Max(math.Abs(float64(points[i][0]-points[i+1][0])), math.Abs(float64(points[i][1]-points[i+1][1]))))
	}

	return steps
}
