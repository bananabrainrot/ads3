package main

import "fmt"

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func buildMaxHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func printHeap(arr []int) {
	fmt.Println("Array view:", arr)
	fmt.Println("Tree view (index: value):")
	for i, v := range arr {
		parent := (i - 1) / 2
		if i == 0 {
			fmt.Printf("  [%d] %d  (root)\n", i, v)
		} else {
			fmt.Printf("  [%d] %d  (parent: [%d]=%d)\n", i, v, parent, arr[parent])
		}
	}
}

func main() {
	arr := []int{64, 66, 23, 16, 53, 89, 15, 70, 48, 10}
	fmt.Println("Original dataset:", arr)

	buildMaxHeap(arr)
	fmt.Println("\n--- After Heapification (Max-Heap) ---")
	printHeap(arr)

	n := len(arr)
	extracted := arr[0]
	arr[0], arr[n-1] = arr[n-1], arr[0]
	heapify(arr, n-1, 0)

	fmt.Println("\n--- After First Extract-Max ---")
	fmt.Println("Extracted value:", extracted)
	fmt.Println("Heap after extraction:", arr[:n-1])
	fmt.Println("Full array state:    ", arr)
}