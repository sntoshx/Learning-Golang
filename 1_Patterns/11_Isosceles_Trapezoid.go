package main
import("fmt")
var height int
func main() {

	fmt.Print("Enter a height for the trapezoid: ")
	fmt.Scan(&height)
	for i:=0; i<height; i++ {
		 for j:=0; j<i; j++ {
			 fmt.Print(" ")
		 }
		 for j:=0; j<height-i+(2*height); j++{
			 fmt.Print(" *")
		 }
		 fmt.Println("")
	}
}