package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число : ")
	fmt.Scan(&a)
	if a > 0 {
		fmt.Println(a + 1)
	} else {
		fmt.Println(a)
	}

}
