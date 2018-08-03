package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Transform reads data from a given input, transforms it using the given delimiter, and writes it to the given output
func Transform(r io.Reader, w io.Writer, delim string) error {
	input, err := ReadLines(r)
	if err != nil {
		return err
	}

	matrix := SplitLines(input, delim)
	transposed := Transpose(matrix)
	filtered := FilterEmpty(transposed)
	result := JoinLines(filtered, delim)

	var output []string
	for _, line := range result {
		output = append(output, fmt.Sprintf("%s%s%s", delim, line, delim))
	}

	w.Write([]byte(strings.Join(output, "\n")))
	return nil
}

// ReadLines reads all lines from the provided io.Reader until EOF or until an error occurs
func ReadLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var output []string

	for {
		if !scanner.Scan() {
			break
		}
		output = append(output, scanner.Text())
	}
	return TrimAll(output), scanner.Err()
}

// SplitLines takes a slice of lines and splits each line using the provided delimiter
func SplitLines(lines []string, delim string) [][]string {
	var output [][]string
	for _, line := range lines {
		words := strings.Split(line, delim)
		output = append(output, TrimAll(words))
	}
	return output
}

// JoinLines takes a slice of slices and joins each inner slice using the provided delimiter
func JoinLines(lines [][]string, delim string) []string {
	var output []string
	for _, line := range lines {
		output = append(output, strings.Join(line, delim))
	}
	return output
}

// TrimAll takes a slice of strings and trims each element.
func TrimAll(input []string) []string {
	var output []string
	for _, line := range input {
		output = append(output, strings.TrimSpace(line))
	}
	return output
}

// Transpose takes a slice of slices and transposes it.
func Transpose(input [][]string) [][]string {
	output := make([][]string, len(input[0]))
	for x := range output {
		output[x] = make([]string, len(input))
	}
	for y, row := range input {
		for x, cell := range row {
			output[x][y] = cell
		}
	}
	return output
}

// FilterEmpty omits all inner slices that contain only empty strings
func FilterEmpty(input [][]string) [][]string {
	var output [][]string
	for _, row := range input {
		var empty = true
		for _, cell := range row {
			if cell != "" {
				empty = false
			}
		}
		if !empty {
			output = append(output, row)
		}
	}
	return output
}
