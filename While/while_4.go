package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	pow := 0
	for n/3 == 0 {
		fmt.Println(true)
		break
	}
}
