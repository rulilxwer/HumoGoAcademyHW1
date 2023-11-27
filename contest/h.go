package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a <= 0 {
		fmt.Println("error")
	} else if a <= 3 {
		fmt.Println("elementary")
	} else if a >= 4 && a <= 6 {
		fmt.Println("average")
	} else if a >= 7 && a <= 9 {
		fmt.Println("sufficient")
	} else if a >= 10 && a <= 12 {
		fmt.Println("high")
	} else {
		fmt.Println("error")
	}

}
