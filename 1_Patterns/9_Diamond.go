package main

import (
	"fmt"
)

var width int

func main() {
	fmt.Print("Enter the width of the diamond: ")
	fmt.Scan(&width)
	for i := 0; i < width-1; i++ {
		for j := 0; j < width - i-1; j++ {
			fmt.Print(" ")
		}
		for j:=0; j<i+1; j++{
			fmt.Print(" *")
		}
		fmt.Println("")
	}
	for i := 0; i < width; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j:=0; j<width-i; j++{
			fmt.Print(" *")
		}
		fmt.Println("")
	}
}
