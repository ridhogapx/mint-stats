package main

import (
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	getOS := exec.Command("bash", "-c", "cat /etc/os-release | grep 'UBUNTU_CODENAME' ")

	out, err := getOS.Output()

	if err != nil {
		color.Red("Erro: %d", err)
	}

	color.Yellow(string(out))

}
