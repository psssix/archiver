package vlc

import (
	"fmt"
	"strings"
)

type (
	encodingTable map[rune]string
	EncodingError struct {
		msg string
		r   rune
	}
)

func NewEncodingError(msg string, r rune) *EncodingError {
	return &EncodingError{msg: msg, r: r}
}

func (e *EncodingError) Error() string {
	return fmt.Sprintf("encoding to binary error, %s %q", e.msg, e.r)
}

// encodeBinary encode string into binary codes string without spaces.
func encodeBinary(str string) (string, error) {
	var buf strings.Builder

	for _, r := range str {
		binary, err := toBinary(r)
		if err != nil {
			return "", err
		}
		buf.WriteString(binary)
	}

	return buf.String(), nil
}

func toBinary(r rune) (string, error) {
	table := newEncodingTable()

	code, ok := table[r]
	if !ok {
		return "", NewEncodingError("unknown character", r)
	}

	return code, nil
}

func newEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}
