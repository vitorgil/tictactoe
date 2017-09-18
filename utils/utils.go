package utils

// IsEnter checks whether the given character is the ENTER key
// Tested on Windows PowerShell.
// Does not really work well, pressing arrows will also evaluate to 'x\00'
func IsEnter(c rune) bool {
	return c == '\x00'
}
