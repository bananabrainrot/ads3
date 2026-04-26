package main

import "fmt"

func mergeSort(arr []int, label string, indent string) []int {
	if len(arr) <= 1 {
		fmt.Printf("%s%s -> %v  (base case)\n", indent, label, arr)
		return arr
	}
	mid := len(arr) / 2
	fmt.Printf("%s%s -> left: %v | right: %v\n", indent, label, arr[:mid], arr[mid:])
	left := mergeSort(arr[:mid], "L", indent+"  ")
	right := mergeSort(arr[mid:], "R", indent+"  ")
	merged := merge(left, right)
	fmt.Printf("%sMerge %v + %v -> %v\n", indent, left, right, merged)
	return merged
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	arr := []int{64, 66, 23, 16, 53, 89, 15, 70, 48, 10}
	fmt.Println("Dataset:", arr)
	fmt.Println("\n--- Divide Phase (Recursion Tree) ---")
	sorted := mergeSort(arr, "root", "")
	fmt.Println("\nSorted result:", sorted)
}