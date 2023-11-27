package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a <= 0 {
		fmt.Println("error")
	} else if a <= 9 {
		fmt.Println("1")
	} else if a <= 99 {
		fmt.Println("2")
	} else if a <= 999 {
		fmt.Println("3")
	} else {
		fmt.Println("more than 3")
	}

}
