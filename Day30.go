package leetcode

/*
K Closest Points to Origin
Solution
We have a list of points on the plane.  Find the K closest points to the origin (0, 0).

(Here, the distance between two points on a plane is the Euclidean distance.)

You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)

Example 1:

Input: points = [[1,3],[-2,2]], K = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].
Example 2:

Input: points = [[3,3],[5,-1],[-2,4]], K = 2
Output: [[3,3],[-2,4]]
(The answer [[-2,4],[3,3]] would also be accepted.)
*/

import (
	"math"
	"sort"
)

type Points [][]int

func kClosest(points [][]int, K int) [][]int {

	sort.Sort(Points(points))
	return points[0:K]
}

func (p Points) Len() int {
	return len(p)
}

func (p Points) Swap(x, y int) {
	p[x], p[y] = p[y], p[x]
}

func (p Points) Less(x, y int) bool {
	x1, y1, x2, y2 := float64(p[x][0]), float64(p[x][1]), float64(p[y][0]), float64(p[y][1])

	return math.Sqrt(math.Pow(x1, 2)+math.Pow(y1, 2)) < math.Sqrt(math.Pow(x2, 2)+math.Pow(y2, 2))
}
