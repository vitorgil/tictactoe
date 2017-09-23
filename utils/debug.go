package utils

/* Set of functions to help debugging. Actual debugger to 
*  be added in the future.
*  Use these functions anywhere necessary. Content will only 
*  be written to output if inDebug is true, which can be set 
*  with ActivateDebug().
*/


import (
	"fmt"
	"github.com/eiannone/keyboard"
)

var inDebug bool
func ActivateDebug(state bool) {
	inDebug = state
}

func Debug(format string, a ...interface{}) {

	if !inDebug {
		return
	}
	
	fmt.Printf(format, a)
}

func DebugString(str string) {
	if !inDebug {
		return
	}

	fmt.Print(str)
}

func WaitInput() {
	if !inDebug {
		return
	}
	keyboard.GetSingleKey()
}
