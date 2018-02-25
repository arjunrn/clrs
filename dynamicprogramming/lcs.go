package dynamicprogramming

import (
	"fmt"
	"strings"
)

// LCS Longest Common Subsequence
func LCS() {
	var x, y string
	fmt.Printf("Enter First Sequence: ")
	fmt.Scanf("%s", &x)
	fmt.Printf("Enter Second Sequence: ")
	fmt.Scanf("%s", &y)
	x = strings.ToUpper(x)
	y = strings.ToUpper(y)
	first := []byte(x)
	second := []byte(y)
	c, b := lcsLength(first, second)
	fmt.Printf("Longest Subsequnce Length: %d\n", c[len(first)][len(second)])
	fmt.Printf("Longest Subsequence: %s\n", getLCS(b, first, len(first)-1, len(second)-1))
}

type previousNode int

const (
	upLeft previousNode = iota
	left
	up
)

func lcsLength(x []byte, y []byte) ([][]int, [][]previousNode) {
	m := len(x)
	n := len(y)
	c := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		c[i] = make([]int, n+1)
	}
	b := make([][]previousNode, m)
	for i := 0; i < m; i++ {
		b[i] = make([]previousNode, n)
	}

	for i := 0; i < m; i++ {
		c[i][0] = 0
	}
	for j := 0; j < n; j++ {
		c[0][j] = 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i-1][j-1] = upLeft
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i-1][j-1] = up
			} else {
				c[i][j] = c[i][j-1]
				b[i-1][j-1] = left
			}
		}
	}
	return c, b
}

func getLCS(b [][]previousNode, x []byte, i, j int) string {
	if i < 0 || j < 0 {
		return ""
	}
	if b[i][j] == upLeft {
		return fmt.Sprintf("%s %c", getLCS(b, x, i-1, j-1), rune(x[i]))
	} else if b[i][j] == up {
		return getLCS(b, x, i-1, j)
	}
	return getLCS(b, x, i, j-1)
}
