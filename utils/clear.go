package utils

import (
	"os"
	"os/exec"
)

func WindowsClear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
