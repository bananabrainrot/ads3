package main

import "fmt"

func linearSearch(arr []int, target int) (int, int) {
	comparisons := 0
	for i, v := range arr {
		comparisons++
		fmt.Printf("Comparison %2d: arr[%d] = %2d == %d? %v\n", comparisons, i, v, target, v == target)
		if v == target {
			return i, comparisons
		}
	}
	return -1, comparisons
}

func main() {
	arr := []int{64, 66, 23, 16, 53, 89, 15, 70, 48, 10}
	target := 70

	fmt.Println("Dataset:", arr)
	fmt.Printf("Searching for T1 = %d\n\n", target)

	idx, comparisons := linearSearch(arr, target)

	fmt.Printf("\nResult: found at index %d\n", idx)
	fmt.Printf("Comparisons made: %d\n", comparisons)
	fmt.Printf("\nAverage-case time complexity: O(n/2) simplified to O(n)\n")
	fmt.Printf("For n=%d elements: expected average comparisons = %d\n", len(arr), len(arr)/2)
}