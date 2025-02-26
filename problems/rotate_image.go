package problems

import "fmt"

func RotateImage(matrix [][]int) {
	fmt.Println(matrix)
	n := len(matrix)

	// Transpose the matrix
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse each row
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}

	fmt.Println(matrix)
}
