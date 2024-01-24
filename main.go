package main

import (
	"fmt"

	"systementor.se/goapi0124/data"
)

func main() {
	emp := data.Employee{
		Id:   1,
		Age:  12,
		City: "Test",
		Namn: "Oliver",
	}
	fmt.Printf("Hello %v", emp)
}
