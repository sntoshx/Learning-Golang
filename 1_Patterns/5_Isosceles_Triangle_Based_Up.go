package main
import (
	"fmt"
)
var height int

func main() {
	fmt.Print("Enter the height of triangle: ")
	fmt.Scan(&height)
	for i := 0; i < height; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*(height-i)-1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
