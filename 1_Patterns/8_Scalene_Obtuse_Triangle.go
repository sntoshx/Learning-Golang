package main
import("fmt")
var height int
func main() {
	fmt.Print("Enter a height: ")
	fmt.Scan(&height)
	for i:=0; i<height; i++ {
		for j:=0; j<i; j++{
			fmt.Print("  ")
		}
		for j:=0; j<i; j++{
			fmt.Print(" *")
		}
		fmt.Println("")
	}
}
