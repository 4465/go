package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,......"}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,......"}

	person, ok := personDB["12345"]

	if ok {
		fmt.Println("Yes", person.Name)
	} else {
		fmt.Println("No")
	}
}
