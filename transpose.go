package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Options used by the Transposer
type Options struct {
	Delimiter  string
	TrimSpaces bool
}

// Transpose reads data from a given input, transposes it using the given options, and writes it to the given output
func Transpose(r io.Reader, w io.Writer, options Options) error {
	input, err := readLines(r)
	if err != nil {
		return err
	}

	wrapped := checkIfWrapped(input, options)
	if wrapped {
		input = clean(input, options)
	}

	matrix := split(input, options)
	transposed := transposeMatrix(matrix)
	result := join(transposed, options)

	if wrapped {
		result = wrap(result, options)
	}

	w.Write([]byte(strings.Join(result, "\n")))
	return nil
}

func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var output []string

	for {
		if !scanner.Scan() {
			break
		}
		output = append(output, strings.TrimSpace(scanner.Text()))
	}
	return output, scanner.Err()
}

// checkIfWrapped returns true iff all lines start and end with the delimiter
func checkIfWrapped(lines []string, options Options) bool {
	for _, line := range lines {
		if !strings.HasPrefix(line, options.Delimiter) || !strings.HasSuffix(line, options.Delimiter) {
			return false
		}
	}
	return true
}

// clean removes any delimiter prefixes, suffixes and leftover whitespace from each line
func clean(lines []string, options Options) []string {
	var output []string
	for _, line := range lines {
		line = strings.TrimPrefix(line, options.Delimiter)
		line = strings.TrimSuffix(line, options.Delimiter)
		if options.TrimSpaces {
			line = strings.TrimSpace(line)
		}
		output = append(output, line)
	}
	return output
}

// wrap adds the delimiter as prefix and suffix of each line
func wrap(lines []string, options Options) []string {
	var output []string
	for _, line := range lines {
		output = append(output, fmt.Sprintf("%s%s%s", options.Delimiter, line, options.Delimiter))
	}
	return output
}

// split splits each line using the delimiter
func split(lines []string, options Options) [][]string {
	var output [][]string
	for _, line := range lines {
		words := strings.Split(line, options.Delimiter)
		output = append(output, words)
	}
	return output
}

// join joins all cells in each row in the matrix using the delimiter, creating a slice of lines
func join(matrix [][]string, options Options) []string {
	var output []string
	for _, line := range matrix {
		output = append(output, strings.Join(line, options.Delimiter))
	}
	return output
}

func transposeMatrix(matrix [][]string) [][]string {
	output := make([][]string, len(matrix[0]))
	for x := range output {
		output[x] = make([]string, len(matrix))
	}
	for y, row := range matrix {
		for x, cell := range row {
			output[x][y] = cell
		}
	}
	return output
}
