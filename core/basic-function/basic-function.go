package main

import (
	"fmt"
)

func add(a int, b int) (ret int) {
	return a + b
}

func main() {
	fmt.Println("I AM BASIC addition FUNCTION!")
	var c int = add(3, 2)
	fmt.Println(c)
}
