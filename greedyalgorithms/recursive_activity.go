package greedyalgorithms

import "fmt"

func getInputs() ([]int, []int, int) {
	var numActivities int
	fmt.Printf("Enter number of activities: ")
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
	return s, f, numActivities
}

// RecursiveActivitySelection recursively selects activity
func RecursiveActivitySelection() {
	s, f, numActivities := getInputs()
	selectedActivities := recursiveActivitySelector(s, f, 0, numActivities)
	fmt.Printf("%v\n", selectedActivities)
}

// IterativeActivitySelection iteratively select activities
func IterativeActivitySelection() {
	s, f, numActivities := getInputs()
	selectedActivities := iterativeActivitySelector(s, f, numActivities)
	fmt.Printf("%v\n", selectedActivities)
}

func iterativeActivitySelector(s []int, f []int, n int) []int {
	activities := []int{1}
	k := 1
	for m := 2; m <= n; m++ {
		if s[m] > f[k] {
			activities = append(activities, m)
			k = m
		}
	}
	return activities
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
