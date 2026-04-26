package main

import "fmt"

func findMinMax(arr []int, low, high int, indent string) (int, int) {
	if low == high {
		fmt.Printf("%sBase case [%d]: single element %d\n", indent, low, arr[low])
		return arr[low], arr[low]
	}
	if high == low+1 {
		fmt.Printf("%sBase case [%d,%d]: two elements %d, %d\n", indent, low, high, arr[low], arr[high])
		if arr[low] >= arr[high] {
			return arr[low], arr[high]
		}
		return arr[high], arr[low]
	}
	mid := (low + high) / 2
	fmt.Printf("%sSplit [%d..%d] -> left [%d..%d], right [%d..%d]\n", indent, low, high, low, mid, mid+1, high)
	lMax, lMin := findMinMax(arr, low, mid, indent+"  ")
	rMax, rMin := findMinMax(arr, mid+1, high, indent+"  ")
	max, min := lMax, lMin
	if rMax > max {
		max = rMax
	}
	if rMin < min {
		min = rMin
	}
	fmt.Printf("%sCombine [%d..%d]: max=%d, min=%d\n", indent, low, high, max, min)
	return max, min
}

func naiveMinMax(arr []int) (int, int, int) {
	max, min := arr[0], arr[0]
	comparisons := 0
	for i := 1; i < len(arr); i++ {
		comparisons++
		if arr[i] > max {
			max = arr[i]
		}
		comparisons++
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min, comparisons
}

func main() {
	arr := []int{64, 66, 23, 16, 53, 89, 15, 70, 48, 10}
	n := len(arr)
	fmt.Println("Dataset:", arr)

	fmt.Println("\n--- D&C Approach ---")
	max, min := findMinMax(arr, 0, n-1, "")
	fmt.Printf("\nMax = %d, Min = %d, Range = %d\n", max, min, max-min)
	dcComparisons := 3*(n/2) - 2
	fmt.Printf("D&C comparisons: ~3n/2 - 2 = %d\n", dcComparisons)

	fmt.Println("\n--- Naive Iterative Approach ---")
	nMax, nMin, naiveComp := naiveMinMax(arr)
	fmt.Printf("Max = %d, Min = %d, Range = %d\n", nMax, nMin, nMax-nMin)
	fmt.Printf("Naive comparisons: 2(n-1) = %d\n", naiveComp)

	fmt.Printf("\nEfficiency: D&C uses %d fewer comparisons than naive (%d vs %d)\n", naiveComp-dcComparisons, dcComparisons, naiveComp)
}