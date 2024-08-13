package main

import "fmt"

var input string
var colorName = map[string]string{"blue": "#0000FF", "green": "#008000", "pink": "#FFC0CB", "yellow": "#FFFF00"}

func main() {
	// input := 2
	// switch input {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("invalid value")
	// }
	fmt.Scan(&input)
	fmt.Println(colorName[input])
}
