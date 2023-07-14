package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")

	out, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
