package main

import "fmt"

func hello() {
	fmt.Println("Hello witty")
}

func plus(value1 int, value2 int) {
	result := value1 + value2
	fmt.Println("result =", result)
}

func pluse3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	plus(1, 2)

	result := pluse3value(3, 5, 22)
	fmt.Println("result =", result)
}
