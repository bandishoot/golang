package main

import "fmt"

func main() {
	m := map[uint16]string{
		1: "один",
		2: "два",
		3: "три",
	}
	fmt.Println(m[1]) // "один"
	fmt.Println(m[2]) // "два"
	fmt.Println(m[3]) // "три"
}
