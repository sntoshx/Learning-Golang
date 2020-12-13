package main
import("fmt")
var height int
func main() {
	fmt.Print("Enter a height: ")
	fmt.Scan(&height)
	for i:=1; i<height+1; i++ {
		for j:=0; j<i; j++{
			fmt.Print(" ")
		}
		for j:=0; j<i; j++{
			fmt.Print(" *")
		}
		fmt.Println("")
	}
}
