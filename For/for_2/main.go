package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	if A > B {
		fmt.Println("error")
	} else {
		for i := A; i <= B; i++ {
			fmt.Println(i)
		}
	}
}
