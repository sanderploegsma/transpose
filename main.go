package main

import (
	"flag"
	"os"
)

var (
	delimiter  = flag.String("d", "|", "Delimiter to use")
	trimSpaces = flag.Bool("t", false, "Set to true to trim whitespace")
)

func main() {
	flag.Parse()

	options := Options{
		Delimiter:  *delimiter,
		TrimSpaces: *trimSpaces,
	}

	if err := Transpose(os.Stdin, os.Stdout, options); err != nil {
		panic(err)
	}
}
