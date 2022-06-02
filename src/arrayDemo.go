package main

import "fmt"

func modify(array [10]int) {
	array[0] = 10
	fmt.Println("In modify(), array value:", array)
}

func main() {
	array := [10]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("In main(), array value:", array)
}
