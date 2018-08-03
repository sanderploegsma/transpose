package main

import (
	"os"
)

var (
	delimiter = "|"
)

func main() {
	if err := Transform(os.Stdin, os.Stdout, delimiter); err != nil {
		panic(err)
	}
}
