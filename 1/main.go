package main

import "fmt"

func main() {
	star(4)
}

func star(num int) {
	a := (num * 2) - 1
	b := a - num
	for i := 0; i < a; i++ {
		if i < num {
			for j := 0; j <= i; j++ {
				fmt.Print("*")
				if i == j {
					fmt.Println("")
				}
			}
		} else {
			for k := b; k > 0; k-- {
				fmt.Print("*")

			}
			fmt.Println("")
			b--
		}

	}

}
