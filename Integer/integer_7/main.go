package main

import "fmt"

func main() {

	var num int
	fmt.Println("Введите двузначное число ")
	fmt.Scan(&num)
	if num < 10 && num > 99 {
		fmt.Println("error")
	}
	first := num / 10
	second := num % 10
	sum := first + second
	mul := first * second
	fmt.Println(sum, mul)
}
