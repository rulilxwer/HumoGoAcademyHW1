package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	var c int
	b = a / 100
	c = a % 10
	if a < 100 {
		fmt.Println("number must be between 100 and 999")
	} else if a > 999 {
		fmt.Println("number must be between 100 and 999")
	} else if b > c {
		fmt.Println(b)
	} else if b < c {
		fmt.Println(c)
	} else {

	}

}
