package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var sum float64
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}
	fmt.Println(sum)
}
