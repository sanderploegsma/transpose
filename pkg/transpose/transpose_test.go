package transpose

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	var cases = []struct {
		delimiter string
		input     string
		output    string
	}{
		{
			delimiter: "|",
			input:     "|Header1|Header2|Header3|\n|Value1|Value2|Value3|",
			output:    "|Header1|Value1|\n|Header2|Value2|\n|Header3|Value3|",
		},
		{
			delimiter: "|",
			input:     "Header1|Header2|Header3\nValue1|Value2|Value3",
			output:    "Header1|Value1\nHeader2|Value2\nHeader3|Value3",
		},
		{
			delimiter: ",",
			input:     "Header1,Header2,Header3\nValue1,Value2,Value3",
			output:    "Header1,Value1\nHeader2,Value2\nHeader3,Value3",
		},
		{
			delimiter: "\t",
			input:     "Header1\tHeader2\nValue1\tValue2",
			output:    "Header1\tValue1\nHeader2\tValue2",
		},
	}

	for _, c := range cases {
		var output bytes.Buffer
		err := Transpose(bytes.NewBufferString(c.input), &output, Options{Delimiter: c.delimiter})
		assert.Nil(t, err)
		assert.Equal(t, c.output, string(output.Bytes()))
	}
}
