package main

import "fmt"

func main() {
	var a, b, n1, n2, n3 int
	fmt.Scan(&a, &b)

	n1 = a - b
	n2 = a + b
	n3 = a * b
	if n1 >= n2 && n1 >= n3 {
		fmt.Println(n1)
	} else if n2 >= n1 && n2 >= n3 {
		fmt.Println(n2)
	} else if n3 >= n1 && n3 >= n2 {
		fmt.Println(n3)
	}

}
