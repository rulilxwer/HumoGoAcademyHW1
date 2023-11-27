package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var result = 0
	if a > b {
		fmt.Println("error")
	} else {
		for i := a; i <= b; i++ {
			result += i
		}
		fmt.Println(result)
	}
}
