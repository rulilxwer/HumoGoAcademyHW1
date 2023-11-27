package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Println("Введите двузначное число :")
	fmt.Scan(&num)

	first := num / 10
	second := num % 10

	swapped := second*10 + first

	fmt.Println(swapped)
}
