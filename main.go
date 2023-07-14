package main

import (
	"os/exec"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

func main() {
	getOS := exec.Command("cat", "/etc/os-release")

	border := strings.Repeat("=", 15)
	msg := border + " Linux Ubuntu Based Information " + border

	out, err := getOS.Output()

	if err != nil {
		color.Red("Error: %v", err)
	}

	color.Green(msg)
	color.Yellow(string(out))
	color.Green("%v ", strings.Repeat("=", utf8.RuneCountInString(msg)))
}
