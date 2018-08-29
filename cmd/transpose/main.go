package main

import (
	"flag"
	"os"

	"github.com/sanderploegsma/transpose/pkg/transpose"
)

var (
	delimiter  = flag.String("d", "|", "Delimiter to use")
	trimSpaces = flag.Bool("t", false, "Set to true to trim whitespace")
)

func main() {
	flag.Parse()

	options := transpose.Options{
		Delimiter:  *delimiter,
		TrimSpaces: *trimSpaces,
	}

	if err := transpose.Transpose(os.Stdin, os.Stdout, options); err != nil {
		panic(err)
	}
}
