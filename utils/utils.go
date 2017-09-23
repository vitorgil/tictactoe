package utils

import (
	"os"
	"os/exec"	
)


// IsEnter checks whether the given character is the ENTER key
// Tested on Windows PowerShell.
// Does not really work well, pressing arrows will also evaluate to 'x\00'
func IsEnter(c rune) bool {
	return c == '\x00'
}

// ClearConsole is a thin wrapper to execute a exec.Command
// with "cls", clearing out the console.
func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
