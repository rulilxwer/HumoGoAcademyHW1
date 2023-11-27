package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите длину ребра куба :")
	fmt.Scan(&a)

	if a == 0 {
		fmt.Println("Введите ненулевое число")
	}

	{
		fmt.Println("OK")
	}
}
