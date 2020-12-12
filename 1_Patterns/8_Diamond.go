package main

import (
	"fmt"
)

var height int

func main() {
	fmt.Print("Enter the width of the diamond: ")
	fmt.Scan(&height)
	for i := 0; i < height-1; i++ {
		for j := 0; j < height - i-1; j++ {
			fmt.Print(" ")
		}
		for j:=0; j<i+1; j++{
			fmt.Print(" *")
		}
		fmt.Println("")
	}
	for i := 0; i < height; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j:=0; j<height-i; j++{
			fmt.Print(" *")
		}
		fmt.Println("")
	}
}
