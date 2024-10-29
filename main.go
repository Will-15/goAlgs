package main

import (
	"fmt"
)

/**
 * main demonstrates the usage of sorting algorithms, searching algorithms and graph algorithms.
 */
func main() {
	// Sorting algorithms
	arr := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("Original array:", arr)

	sortedBubble := BubbleSort(arr)
	fmt.Println("Bubble Sorted:", sortedBubble)

	sortedQuick := QuickSort(arr)
	fmt.Println("Quick Sorted:", sortedQuick)

	sortedMerge := MergeSort(arr)
	fmt.Println("Merge Sorted:", sortedMerge)

	// Searching algorithms
	target := 5
	index := BinarySearch(sortedQuick, target)
	fmt.Printf("Binary Search: Element %d found at index %d\n", target, index)

	index = LinearSearch(arr, target)
	fmt.Printf("Linear Search: Element %d found at index %d\n", target, index)

	// Graph algorithms
	graph := NewGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	fmt.Println("BFS Traversal:")
	graph.BFS(0)

	fmt.Println("DFS Traversal:")
	graph.DFS(0)
}
