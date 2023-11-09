package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	fmt.Println("P =", 2 * (a + b))

	fmt.Println("S =", a * b)
}
