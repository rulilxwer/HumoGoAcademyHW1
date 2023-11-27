package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	if A < B {
		fmt.Println("error")
	} else {
		for A >= B {
			A -= B
		}
		fmt.Println("Длина незанятой части отрезка A:", A)
	}

}
