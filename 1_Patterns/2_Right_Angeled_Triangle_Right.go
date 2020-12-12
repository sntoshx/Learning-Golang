package main

import (
	"fmt"
)

var height int

func main() {
	fmt.Print("Enter the height of the triangle: ")
	fmt.Scanln(&height)
	for i := 1; i <= height; i++ {
		for j := 0; j < height-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
