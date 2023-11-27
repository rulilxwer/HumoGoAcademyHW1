package main

import "fmt"

func main() {
	var K, N int
	fmt.Scan(&K, &N)
	if N <= 0 {
		fmt.Println("error")
	} else {
		for i := 1; i <= N; i++ {
			fmt.Println(K)
		}
	}
}
