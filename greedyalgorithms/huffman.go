package greedyalgorithms

import (
	"container/heap"
	"fmt"
)

func getHuffmanInputs() HuffmanPrioQueue {
	fmt.Printf("Enter number of chars: ")
	var numChars int
	fmt.Scanf("%d", &numChars)
	var value string
	var freq int
	fmt.Printf("Enter the values and their frequencies: \n")
	prioQueue := make(HuffmanPrioQueue, numChars)
	for i := 0; i < numChars; i++ {
		fmt.Scanf("%s %d", &value, &freq)
		prioQueue[i] = &huffmanChar{value: value, freq: freq}

	}
	heap.Init(&prioQueue)
	return prioQueue
}

func printHuffmanCharEncoding(char *huffmanChar, links []int) {
	if char.value != "" {
		fmt.Printf("Value: %s Encoding: %v\n", char.value, links)
		return
	}
	printHuffmanCharEncoding(char.left, append(links, 0))
	printHuffmanCharEncoding(char.right, append(links, 1))
}

// HuffManTree create a huffman tree
func HuffManTree() {
	pq := getHuffmanInputs()
	for pq.Len() > 1 {
		charRight := heap.Pop(&pq).(*huffmanChar)
		charLeft := heap.Pop(&pq).(*huffmanChar)
		charRight.index = -1
		charLeft.index = -1
		newChar := &huffmanChar{
			value: "",
			freq:  charLeft.freq + charRight.freq,
			left:  charLeft,
			right: charRight,
		}
		heap.Push(&pq, newChar)
	}
	rootChar := heap.Pop(&pq).(*huffmanChar)
	printHuffmanCharEncoding(rootChar, []int{})
}
