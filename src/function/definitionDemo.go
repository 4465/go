package main

import (
	"errors"
	"fmt"
)

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("error message")
	}
	return a + b, nil
}

func main() {
	ret, error := Add(3, 6)
	fmt.Println(ret)
	fmt.Println(error)
}
