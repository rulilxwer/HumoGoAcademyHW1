package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		if a[i]%2 == 0 {
			fmt.Printf()
		}
	}
}
