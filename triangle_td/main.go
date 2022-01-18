/*
class Solution {
  public int minimumTotal(List<List<Integer>> triangle) {
    int[][] dp=new int[triangle.size()][triangle.size()];
    return recur(triangle, 0, 0, dp);
  }

  public int recur(List<List<Integer>> triangle, int row, int col, int[][] dp) {
    if (row == triangle.size()) {return 0;}

    // Key line: to avoid duplicate calculation
    if (dp[row][col] != 0) {return dp[row][col];}

    int curr = triangle.get(row).get(col);
    int first = recur(triangle, row + 1, col, dp);
    int second = recur(triangle, row + 1, col + 1, dp);
    dp[row][col] = Math.min(first, second) + curr;
    return dp[row][col];
  }
}
*/

package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	dp := buildEmptyCache(len(triangle))
	return recur(triangle, 0, 0, dp)
}

func buildEmptyCache(size int) [][]int {
	emptyCache := make([][]int, size)

	for i := 0; i < size; i++ {
		emptyCache[i] = make([]int, size)

		for j := 0; j < size; j++ {
			emptyCache[i][j] = 0
		}
	}

	return emptyCache
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func recur(triangle [][]int, row int, col int, dp [][]int) int {
	if row == len(triangle) {
		return 0
	}

	if dp[row][col] != 0 {
		return dp[row][col]
	}

	var curr int = triangle[row][col]
	var first int = recur(triangle, row+1, col, dp)
	var second int = recur(triangle, row+1, col+1, dp)
	dp[row][col] = min(first, second) + curr
	return dp[row][col]
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
