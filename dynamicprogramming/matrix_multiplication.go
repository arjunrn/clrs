package dynamicprogramming

import "fmt"

// MatrixMultiplicationCost Calculate optimumcost of matrix chain multiplication
func MatrixMultiplicationCost() {
	var numMatrices int
	fmt.Printf("Enter number of matrices: ")
	fmt.Scanf("%d", &numMatrices)
	fmt.Printf("Enter matrix dimensions: ")
	dimensions := make([]int, numMatrices+1)
	for i := 0; i <= numMatrices; i++ {
		fmt.Scanf("%d", &dimensions[i])
	}
	m, s := matrixChainOrder(dimensions)
	fmt.Printf("Optimum Multiplication Cost: %d\n", m[0][numMatrices-1])
	fmt.Printf("Parens Order: %s\n", optimalParens(s, 0, numMatrices-1))
}

func matrixChainOrder(p []int) ([][]int, [][]int) {
	n := len(p) - 1
	m := make([][]int, n)
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		m[i][i] = 0
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			m[i][j] = -1
			for k := i; k <= j-1; k++ {
				q := m[i][k] + m[k+1][j] + p[i]*p[k+1]*p[j+1]
				if m[i][j] == -1 || q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}
	return m, s
}

func optimalParens(s [][]int, i int, j int) string {
	if i == j {
		return fmt.Sprintf("M%d", i)
	}
	return fmt.Sprintf("( %s  %s )", optimalParens(s, i, s[i][j]), optimalParens(s, s[i][j]+1, j))

}
