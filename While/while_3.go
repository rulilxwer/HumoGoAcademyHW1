package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	del := 0

	for n >= k {
		n -= k
		del++
	}

	fmt.Println(del)
	fmt.Println(n)

}
