package simplog

import (
	"fmt"
	"time"

	flag "github.com/cohix/simplflag"
)

var includeDebug bool

func init() {
	_, includeDebug = flag.CheckFlag("debug")
}

// LogError logs an error
func LogError(err error) {
	fmt.Printf("%s (E) %s\n", time.Now(), err.Error())
}

// LogWarn logs a warning
func LogWarn(msg string) {
	fmt.Printf("%s (W) %s\n", time.Now(), msg)
}

// LogInfo logs an information message
func LogInfo(msg string) {
	fmt.Printf("%s (I) %s\n", time.Now(), msg)
}

// LogDebug logs an information message
func LogDebug(msg string) {
	if includeDebug {
		fmt.Printf("%s (D) %s\n", time.Now(), msg)
	}
}

// LogTrace logs a function call and returns a function to be deferred marking the end of the function
func LogTrace(name string) func() {
	LogInfo(fmt.Sprintf("[trace] %s began", name))
	return func() {
		LogInfo(fmt.Sprintf("[trace] %s completed", name))
	}
}
