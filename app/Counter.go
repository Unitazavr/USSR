package main

func CountMiddleValue(matrix [][]int) float64 {
	var answer = 0.0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			num := float64(matrix[i][j])
			answer += num
		}
	}
	return answer / float64(len(matrix)*len(matrix[0]))
}
