package xlog

import (
	"fmt"
	"time"
)

const timeFormat = "01/02 15:04:05"

// Infof provides info level log statement
func Infof(format string, a ...interface{}) (int, error) {
	logTime := time.Now().Format(timeFormat)
	formatted := fmt.Sprintf("Info    %v: ", logTime)
	return fmt.Printf(formatted+format+"\n", a...)
}

// Verbosef provides verbose level log statement
func Verbosef(format string, a ...interface{}) (int, error) {
	logTime := time.Now().Format(timeFormat)
	formatted := fmt.Sprintf("Verbose %v: ", logTime)
	return fmt.Printf(formatted+format+"\n", a...)
}

// Warningf provides warning level log statment
func Warningf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(format, a...)
}

// Errorf provides error level log statement
func Errorf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(format, a...)
}
