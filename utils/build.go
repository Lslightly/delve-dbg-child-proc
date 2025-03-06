package main

import (
	"fmt"
	"time"
)

func compile() {
	fmt.Println("compile")
}

func main() {
	a := 1
	b := 2 + a
	_ = b
	compile()
	time.Sleep(5 * time.Second)
}
