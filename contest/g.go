package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 10 {
		fmt.Println("number must be between 10 and 99")
	} else if a > 99 {
		fmt.Println("number must be between 10 and 99")
	} else {
		var b int
		var c int
		b = a / 10
		c = a % 10
		fmt.Println(b, c)
	}
}
