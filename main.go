package main

import "fmt"

func main() {
	condition := true
	x := 1
	if condition {
		x := 2
		fmt.Println("x =", x)
	}
	fmt.Println("x =", x)
}
