package utils

import (
	"fmt"
	"os"
)

func ErrHandle(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
