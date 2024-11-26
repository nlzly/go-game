package color

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

func PrintRed(input string) {
	fmt.Println(Red, input, Reset)
}

func PrintGreen(input string) {
	fmt.Println(Green, input, Reset)
}
