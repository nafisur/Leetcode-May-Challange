package leetcode

/*
Problem statement

An image is represented by a 2-D array of integers, each integer representing the pixel value of the image (from 0 to 65535).

Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill, and a pixel value newColor, "flood fill" the image.

To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color as the starting pixel), and so on. Replace the color of all of the aforementioned pixels with the newColor.

At the end, return the modified image.

Input:
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]

*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	if image[sr][sc] == newColor {
		return image
	}

	fill(image, sr, sc, image[sr][sc], newColor)
	return image
}

func fill(image [][]int, i int, j int, color int, newColor int) {
	if i < 0 || i >= len(image) || j < 0 || j >= len(image[i]) || image[i][j] != color {
		return
	}

	image[i][j] = newColor
	fill(image, i+1, j, color, newColor)
	fill(image, i-1, j, color, newColor)
	fill(image, i, j+1, color, newColor)
	fill(image, i, j-1, color, newColor)
}
