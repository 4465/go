package main

import (
	"fmt"
)

func search() {
	str := "Hello, 世界"
	n := len(str)
	// 字节数组方式遍历
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	//unicode字符遍历
	fmt.Println("=======================")
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}

func main() {
	var str string
	str = "Hello World"
	ch := str[0]

	fmt.Printf("The length of \" %s \" is %d. \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	x := "hello"
	y := "123"
	fmt.Println(x + y)
	fmt.Println("=================================")
	search()
}
