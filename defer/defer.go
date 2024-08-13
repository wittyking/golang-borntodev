package main

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("result :", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	}
}

func deferloop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("j =", i)
	}
}

func main() {
	// fmt.Println("Welcom to Calculator")
	// defer fmt.Println("End")
	// defer add(20, 10)
	// defer add(30, 10)
	// defer add(10, 10)
	loop()
	deferloop()
}
