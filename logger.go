package simplog

import (
	"fmt"

	flag "github.com/cohix/simplflag"
)

var includeDebug bool

func init() {
	_, includeDebug = flag.CheckFlag("debug")
}

// LogError logs an error
func LogError(err error) {
	fmt.Printf("(E) %s\n", err.Error())
}

// LogWarn logs a warning
func LogWarn(msg string) {
	fmt.Printf("(W) %s\n", msg)
}

// LogInfo logs an information message
func LogInfo(msg string) {
	fmt.Printf("(I) %s\n", msg)
}

// LogDebug logs an information message
func LogDebug(msg string) {
	if includeDebug {
		fmt.Printf("(D) %s\n", msg)
	}
}

// LogTrace logs a function call and returns a function to be deferred marking the end of the function
func LogTrace(name string) func() {
	LogInfo(fmt.Sprintf("[trace] %s", name))
	return func() {
		LogInfo(fmt.Sprintf("%s completed", name))
	}
}
