package main

import (
	"fmt"
	"os"

	"github.com/arjunrn/cormen/dynamicprogramming"
)

func main() {
	fmt.Printf("1. Naieve Cut Rod\n")
	fmt.Printf("2. Memoized Cut Rod\n")
	fmt.Printf("3. Matrix Multiplication Optimum Cost\n")
	fmt.Println("4. Longest Common Subsequence")
	fmt.Printf("Enter choice: ")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		dynamicprogramming.NaieveCutRod()
	case 2:
		dynamicprogramming.MemoizedCutRod()
	case 3:
		dynamicprogramming.MatrixMultiplicationCost()
	case 4:
		dynamicprogramming.LCS()
	default:
		dynamicprogramming.MemoizedCutRod()
		fmt.Printf("Invalid Choice\n")
		os.Exit(1)
	}
}
