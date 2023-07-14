package main

import (
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func main() {
	getOS := []string{
		"cat /etc/os-release | grep 'PRETTY_NAME'",

		"cat /etc/os-release | grep 'UBUNTU_CODENAME'",
	}

	border := strings.Repeat("#", 30)
	color.Green(border)

	for _, cmd := range getOS {
		c := exec.Command("bash", "-c", cmd)
		out, err := c.Output()

		if err != nil {
			color.Red("Error: %d", err)
		}

		color.Yellow(string(out))
	}

	color.Green(border)
}
