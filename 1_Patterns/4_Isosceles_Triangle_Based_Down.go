package main
import (
	"fmt"
)

var height int

func main() {
	fmt.Print("Enter the height of triangle: ")
	fmt.Scan(&height)
	for i := 1; i <= height; i++ {
		for j := 0; j <= height-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*(i)-1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
