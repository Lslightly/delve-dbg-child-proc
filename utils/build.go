package main

import (
	"fmt"
)

func compile() {
	fmt.Println("compile")
}

func main() {
	a := 1
	b := 2 + a
	_ = b
	compile()
}
