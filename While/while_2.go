package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	count := 0
	if A < B {
		fmt.Println("error")
	} else {
		for A >= B {
			A -= B
			count++
		}
		fmt.Println("Количество отрезков B, размещенных на отрезке A:", count)
	}

}
