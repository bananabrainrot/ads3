package main

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	fmt.Printf("Pivot = arr[%d] = %d\n", low, pivot)
	fmt.Printf("Initial array: %v\n\n", arr)

	i := low
	for j := low + 1; j <= high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			fmt.Printf("arr[%d]=%d <= pivot(%d): swap arr[%d] and arr[%d] -> %v\n", j, arr[i], pivot, i, j, arr)
		} else {
			fmt.Printf("arr[%d]=%d >  pivot(%d): no swap\n", j, arr[j], pivot)
		}
	}

	arr[low], arr[i] = arr[i], arr[low]
	fmt.Printf("\nPlace pivot %d at final index %d -> %v\n", pivot, i, arr)
	return i
}

func main() {
	arr := []int{64, 66, 23, 16, 53, 89, 15, 70, 48, 10}
	fmt.Println("Dataset:", arr)
	fmt.Println("\n--- Partitioning Step (Lomuto Scheme, first element as pivot) ---\n")

	pivotIdx := partition(arr, 0, len(arr)-1)

	fmt.Printf("\nFinal pivot position: index %d\n", pivotIdx)
	fmt.Println("Elements <= pivot (left):", arr[:pivotIdx])
	fmt.Println("Pivot:", arr[pivotIdx])
	fmt.Println("Elements >  pivot (right):", arr[pivotIdx+1:])
}
