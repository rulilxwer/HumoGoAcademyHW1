package main

import "fmt"

func main() {
	var a string

	fmt.Scan(&a)
	if a == "a" || a == "e" || a == "i" || a == "o" || a == "u" {
		fmt.Println("Vowel")
	} else {
		fmt.Println("CONSONANT")
	}

}
