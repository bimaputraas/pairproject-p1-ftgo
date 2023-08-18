package cli

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearCli() {
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
		fmt.Println("Tidak dapat membersihkan layar pada sistem ini.")
	}
}
