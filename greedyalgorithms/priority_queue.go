package greedyalgorithms

type huffmanChar struct {
	value string
	freq  int
	index int
	left  *huffmanChar
	right *huffmanChar
}

type HuffmanPrioQueue []*huffmanChar

func (pq HuffmanPrioQueue) Len() int {
	return len(pq)
}

func (pq HuffmanPrioQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq HuffmanPrioQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *HuffmanPrioQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*huffmanChar)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *HuffmanPrioQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// func (pq *HuffmanPrioQueue) update(char *huffmanChar, value string, freq int) {
// 	char.value = value
// 	char.freq = freq
// 	heap.Fix(pq, char.index)
// }
