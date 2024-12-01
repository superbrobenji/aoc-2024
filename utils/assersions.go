package utils

import (
	"fmt"
	"os"
)

func Assert(assertion bool, message string) {
	if !assertion {
		fmt.Println(message)
		os.Exit(1)
	}
}
