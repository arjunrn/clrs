package greedyalgorithms

import "fmt"

// RecursiveActivitySelection recursively selects activity
func RecursiveActivitySelection() {
	fmt.Printf("Enter number of activities: ")
	var numActivities int
	fmt.Scanf("%d", &numActivities)
	s := make([]int, numActivities+1)
	f := make([]int, numActivities+1)
	s[0] = 0
	f[0] = 0
	fmt.Printf("Enter start and end times\n")
	var startTime, endTime int
	for i := 1; i <= numActivities; i++ {
		fmt.Scanf("%d %d", &startTime, &endTime)
		s[i] = startTime
		f[i] = endTime
	}
	selectedActivities := recursiveActivitySelector(s, f, 0, numActivities)
	fmt.Printf("%v\n", selectedActivities)
}

func recursiveActivitySelector(s []int, f []int, k int, n int) []int {
	m := k + 1
	for m <= n && s[m] < f[k] {
		m++
	}
	if m <= n {
		return append([]int{m}, recursiveActivitySelector(s, f, m, n)...)
	}
	return []int{}
}
