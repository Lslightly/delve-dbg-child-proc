package main

import (
	"fmt"
	"os/exec"
)

func drive() {
	cmd := exec.Command("./utils/compile")
	bs, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

func main() {
	drive()
}
