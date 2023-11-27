package main

import "fmt"

func main() {

	var num int
	fmt.Println("Введите трехзначное число :")
	fmt.Scan(&num)
	if num < 100 && num > 999 {
		fmt.Println("error")
	}
	first := num / 100

	fmt.Println(first)
}
