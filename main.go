package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := []string{
		"ping -c 2 google.com",
		"ping -c 2 facebook.com",
	}

	for _, command := range c {
		cmd := exec.Command("bash", "-c", command)
		out, err := cmd.Output()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}
}
