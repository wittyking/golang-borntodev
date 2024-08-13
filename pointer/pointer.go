package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}

func zeropointer(ivalue *int) {
	*ivalue = 0
}

func main() {
	i := 1
	fmt.Println("i=", i)

	zerovalue(i)
	fmt.Println("i from function zerovalue =", i)

	zeropointer(&i)
	fmt.Println("i value function zeropointer =", i)
	fmt.Println("i address function zeropointer =", &i)
}
