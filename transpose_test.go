package main

import (
	"bytes"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFilterEmpty(t *testing.T) {
	input := [][]string {
		{ "", ""},
		{"a", ""},
		{"", "b"},
		{"", ""},
		{"a", "b"},
		{"", ""},	
	}
	expected := [][]string {
		{"a", ""},
		{"", "b"},
		{"a", "b"},
	}
	assert.Equal(t, expected, FilterEmpty(input))
}

func TestTranspose(t *testing.T) {
	input := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
	}
	expected := [][]string{
		{"a", "d"},
		{"b", "e"},
		{"c", "f"},
	}
	assert.Equal(t, expected, Transpose(input))
}

func TestTrimAll(t *testing.T) {
	input := []string { "   a", "b   ", "  c ", "d" }
	expected := []string{"a", "b", "c", "d"}
	assert.Equal(t, expected, TrimAll(input))
}

func TestJoinLines(t *testing.T) {
	input := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
	}
	expected := []string{
		"a,b,c",
		"d,e,f",
	}
	assert.Equal(t, expected, JoinLines(input, ","))
}

func TestSplitLines(t *testing.T) {
	input := []string{
		"a,b,c",
		"d,e,f",
	}
	expected := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
	}
	assert.Equal(t, expected, SplitLines(input, ","))
}

func TestReadLines(t *testing.T) {
	input := "a\nb\nc"
	expected := []string{"a","b","c"}
	output, err := ReadLines(bytes.NewBufferString(input))
	assert.Nil(t, err)
	assert.Equal(t, expected, output)
}

func TestTransform(t *testing.T) {
	input := "|Header1|Header2|Header3|\n|Value1|Value2|Value3|"
	expected := "|Header1|Value1|\n|Header2|Value2|\n|Header3|Value3|"

	var output bytes.Buffer
	err := Transform(bytes.NewBufferString(input), &output, "|")
	assert.Nil(t, err)
	assert.Equal(t, expected, string(output.Bytes()))
}