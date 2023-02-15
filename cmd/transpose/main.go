package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sanderploegsma/transpose/pkg/transpose"
)

var (
	version          = "dev"
	commit           = "none"
	date             = "unknown"
	versionOption    = flag.Bool("v", false, "Prints current version")
	delimiterOption  = flag.String("d", "|", "Delimiter to use")
	trimSpacesOption = flag.Bool("t", false, "Set to true to trim whitespace")
)

func main() {
	flag.Parse()

	if *versionOption {
		fmt.Printf("Transpose %s\nCommit: %s\nDate: %s", version, commit, date)
		os.Exit(0)
	}

	options := transpose.Options{
		Delimiter:  *delimiterOption,
		TrimSpaces: *trimSpacesOption,
	}

	if err := transpose.Transpose(os.Stdin, os.Stdout, options); err != nil {
		panic(err)
	}
}
