package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println(a)
	fmt.Println(b)
	c := a + b
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("hello, go")
}
