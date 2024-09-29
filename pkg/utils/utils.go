package utils

import (
	"fmt"
	"log"
)

func PrintError(err error, message string) bool {
	if err != nil {
		fmt.Printf("%s:%v", message, err)
		return true
	}
	return false
}

func LogFatalError(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%v", message, err)

	}
}
