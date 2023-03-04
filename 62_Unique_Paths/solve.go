package main

func UniquePaths(m int, n int) int {

	//Preprocessing, init matrix
	var result [][]int = make([][]int, m)
	for index := range result {
		result[index] = make([]int, n)
	}

	// Solve
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j] + result[i][j-1]
			}
		}
	}
	return result[m-1][n-1]

}
