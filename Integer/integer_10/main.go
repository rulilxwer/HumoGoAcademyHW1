package main

import "fmt"

func main() {

	var num int
	fmt.Println("Введите трехзначное число :")
	fmt.Scan(&num)
	if num < 100 && num > 999 {
		fmt.Println("error")
	}
	second := (num / 10) % 10
	third := num % 10
	fmt.Println(third, second)
}
