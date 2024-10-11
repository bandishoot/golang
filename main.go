package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scan(&size)
	slice := make([]int, size)
	for i := 0; i < len(slice); i++ {
		fmt.Scan(&slice[i])
	}
	for i := 0; i < len(slice); i++ {
		if slice[i]%3 == 0 {
			fmt.Print(slice[i], " ")
		}
	}
}
