package lib

func Transpose(matrix [][]int) [][]int {
	transposed := make([][]int, len(matrix))
	for i := range transposed {
		transposed[i] = make([]int, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func Determinant(matrix [][]int) int {
	size := len(matrix)

	if size == 1 {
		return matrix[0][0]
	}

	if size == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	var det int
	for i := 0; i < size; i++ {
		sign := 1
		if i%2 != 0 {
			sign = -1
		}
		term := matrix[0][i] * sign * Determinant(Minor(matrix, 0, i))
		det += term
	}

	return det
}

func Multiply(A [][]int, B [][]int) [][]int {
	rowsA := len(A)
	colsA := len(A[0])
	colsB := len(B[0])

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return result
}

func MultiplyWithConstant(matrix [][]int, constant int) [][]int {
	result := make([][]int, len(matrix))
	for i := range result {
		result[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[i][j] = matrix[i][j] * constant
		}
	}

	return result
}

func Adjoint(matrix [][]int) [][]int {
	size := len(matrix)

	cofactors := make([][]int, size)
	for i := range cofactors {
		cofactors[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sign := 1
			if (i + j) % 2 != 0 {
				sign = -1
			}
			minor := Minor(matrix, i, j)
			cofactors[j][i] = sign * Determinant(minor)
		}
	}

	return cofactors
}

func Minor(matrix [][]int, row int, col int) [][]int {
	size := len(matrix)
	minor := make([][]int, size-1)
	for i := range minor {
		minor[i] = make([]int, size-1)
	}

	for i, r := range matrix {
		if i == row {
			continue
		}
		for j, val := range r {
			if j == col {
				continue
			}
			rowIdx := i
			if i > row {
				rowIdx--
			}
			colIdx := j
			if j > col {
				colIdx--
			}
			minor[rowIdx][colIdx] = val
		}
	}

	return minor
}
