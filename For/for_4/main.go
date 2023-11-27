package main

import "fmt"

func main() {
	var A float32
	fmt.Scan(&A)
	if A <= 0 {
		fmt.Println("error")
	} else {
		for i := 1; i <= 10; i++ {
			fmt.Println(i, float32(i)*A)
		}
	}
}
