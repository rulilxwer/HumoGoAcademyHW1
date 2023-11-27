package main

import "fmt"

func main() {

	var number int
	fmt.Println("введите число")
	fmt.Scan(number)

	var tens int = number / 10

	var ones int = number % 10

	fmt.Println(tens, ones)
}
