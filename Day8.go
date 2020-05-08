package leetcode

/*
Problem statement

You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point.
 Check if these points make a straight line in the XY plane.

 Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true

Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
Output: false

Constraints:

2 <= coordinates.length <= 1000
coordinates[i].length == 2
-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
coordinates contains no duplicate point.

slope = (y2-y1)/(x2-x1) == (y3-y1)/(x3-x1)
to avoid divide by 0 error, I'll be using:
(y2-y1)*(x3-x1)==(x2-x1)*(y3-y1)
*/

func checkStraightLine(coordinates [][]int) bool {

	if len(coordinates) < 2 {
		return true
	}

	for i := 2; i < len(coordinates); i++ {
		if ((coordinates[1][1] - coordinates[0][1]) * (coordinates[i][0] - coordinates[0][0])) != ((coordinates[1][0] - coordinates[0][0]) * (coordinates[i][1] - coordinates[0][1])) {
			return false
		}
	}
	return true

}
