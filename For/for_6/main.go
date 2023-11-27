package main

import "fmt"

func main() {
	var A float64
	fmt.Scan(&A)
	if A <= 0 {
		fmt.Println("error")
	} else {
		for i := float64(1.2); i <= 2; i += float64(0.2) {
			fmt.Println(i, i*A)
		}
	}
}
