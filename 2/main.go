package main

import "fmt"

func fibonacci(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
		fmt.Print(a)
		fmt.Print(",")

	}
	fmt.Println("")
	return a
}

func main() {
	fibonacci(10)
}
