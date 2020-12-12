package main

import (
	"fmt"
)

var height int

func main() {
	fmt.Print("Enter the height of the triangle: ")
	fmt.Scan(&height)
	for i := 1; i <= height; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
	for i := height - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}

}
