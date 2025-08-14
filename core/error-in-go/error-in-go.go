package main

import (
	"errors"
	"fmt"
)

// Go doesn't return exceptions,
// it returns error value explicitly

// means you have to check if error or=ccured or not?

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("test error")
	}
	return a / b, nil
}

func main() {
	op, err := divide(4, 0)
	fmt.Printf("op:", op)
	fmt.Println("err", err)
}
