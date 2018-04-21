package bytes

import (
	"bytes"
)

func Bytes(a, b string) (str string, n int, err error) {
	// Initalize Buffer
	buf := new(bytes.Buffer)

	// First Round
	amt1, err := buf.WriteString(a)
	if err != nil {
		return "", amt1, err
	}

	// Second Round
	amt2, err := buf.WriteString(b)
	if err != nil {
		return "", amt2, err
	}

	// Assemble Result
	str = buf.String()
	n = amt1 + amt2
	err = nil
	return str, n, err
}