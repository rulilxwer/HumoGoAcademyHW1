package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите значение a")
	fmt.Scan(&a)
	var b int
	fmt.Println("Введите значение b")
	fmt.Scan(&b)
	if a > b {
		fmt.Println(a, b)
	} else {
		fmt.Println(b, a)
	}
}
