package dynamicprogramming

import (
	"fmt"
	"math"
)

// Prices The Rod cut prices
var Prices = map[int64]int64{
	1:  1,
	2:  5,
	3:  8,
	4:  9,
	5:  10,
	6:  17,
	7:  17,
	8:  20,
	9:  24,
	10: 20,
}

func MemoizedCutRod() {
	var rodLength int64
	fmt.Scanf("%d", &rodLength)
	revenue := make([]int64, rodLength+1)
	var i int64
	for i = 0; i < rodLength+1; i++ {
		revenue[i] = -1
	}
	fmt.Printf("%d", memoizedCutRodAux(rodLength, revenue))
}

func memoizedCutRodAux(rodLength int64, revenue []int64) int64 {
	if revenue[rodLength] >= 0 {
		return revenue[rodLength]
	}
	var q int64
	if rodLength == 0 {
		q = 0
	} else {
		q = -1
		var i int64
		for i = 1; i <= rodLength; i++ {
			q = int64(math.Max(float64(q), float64(Prices[i]+memoizedCutRodAux(rodLength-i, revenue))))
		}
	}
	revenue[rodLength] = q
	return q
}

// NaieveCutRod Cut rod function without memoization
func NaieveCutRod() {
	var rodLength int64
	fmt.Scanf("%d", &rodLength)
	profit := getMaxProfit(rodLength)
	fmt.Printf("%d", profit)
}

func getMaxProfit(rodLength int64) int64 {
	if rodLength == 0 {
		return 0
	}
	var q int64 = -1
	var i int64
	for i = 1; i <= rodLength; i++ {
		q = int64(math.Max(float64(q), float64(Prices[i]+getMaxProfit(rodLength-i))))
	}
	return q
}
