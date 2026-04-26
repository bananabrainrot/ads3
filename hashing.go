package main

import "fmt"

const M = 7

func main() {
	arr := []int{64, 66, 23, 16, 53, 89, 15, 70, 48, 10}
	table := make([]int, M)
	for i := range table {
		table[i] = -1
	}

	fmt.Println("Dataset:", arr)
	fmt.Printf("Hash Table size M = %d, h(k) = k mod %d\n\n", M, M)

	for _, k := range arr {
		h := k % M
		fmt.Printf("Insert %d: h(%d) = %d mod %d = %d", k, k, k, M, h)

		if table[h] == -1 {
			table[h] = k
			fmt.Printf("  -> placed at index %d\n", h)
		} else {
			fmt.Printf("  -> COLLISION with %d, probing...\n", table[h])
			for step := 1; step < M; step++ {
				probe := (h + step) % M
				if table[probe] == -1 {
					table[probe] = k
					fmt.Printf("   step %d: index %d is free -> placed at index %d\n", step, probe, probe)
					break
				}
				fmt.Printf("   step %d: index %d occupied by %d\n", step, probe, table[probe])
			}
		}
	}

	fmt.Println("\n--- Final Hash Table ---")
	for i, v := range table {
		if v == -1 {
			fmt.Printf("  [%d]: empty\n", i)
		} else {
			fmt.Printf("  [%d]: %d  (h=%d)\n", i, v, v%M)
		}
	}
}
