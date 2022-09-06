package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Printf("a = %d, b = %d\n", a, b)
	var c int = 0
	c = a
	a = b
	b = c
	fmt.Printf("a = %d, b = %d\n", a, b)
}
