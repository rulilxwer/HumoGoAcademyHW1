package main

import "fmt"

func main() {

	var sizeOfArr, numArr, count int
	fmt.Scan(&sizeOfArr)
	arr := make([]int, sizeOfArr)
	for i := 0; i < sizeOfArr; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&numArr)
	for i := 0; i < sizeOfArr; i++ {
		if arr[i] == numArr {
			count++
		}
	}
	fmt.Println(count)
}
