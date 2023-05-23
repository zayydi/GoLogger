package utils

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		color.New(color.FgRed).Println("Unsupported platform. Unable to clear the command line.")
	}
}
