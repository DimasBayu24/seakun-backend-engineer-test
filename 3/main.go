package main

import (
	"fmt"
)

func main() {
	multiply(4)
}

func multiply(num int) {
	// fmt.Println(num)
	a := num
	for i := num - 1; i > 0; i-- {
		a = a * i
		if i == 1 {
			fmt.Println(a)
		}

	}
}
