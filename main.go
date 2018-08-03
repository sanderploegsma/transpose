package main

import (
	"flag"
	"os"
)

var (
	delimiter = flag.String("d", "|", "Delimiter to use")
)

func main() {
	flag.Parse()

	if err := Transform(os.Stdin, os.Stdout, *delimiter); err != nil {
		panic(err)
	}
}
