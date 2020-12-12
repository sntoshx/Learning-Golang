package main

import (
	"fmt"
)

var l, b int

func main() {
	fmt.Print("Enter length and breadth: ")
	fmt.Scan(&l, &b)
	for i := 0; i < b; i++ {
		for j := 0; j < l; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
