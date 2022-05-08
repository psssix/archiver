package vlc

import (
	"fmt"
	"strings"
	"unicode"
)

func Encode(str string) string {
	str = escapeUpper(str)
	chunks := splitBinaryChunks(encodeBinary(str), chunkSize)

	return chunks.toHex().String()
}

func Decode(str string) string {
	bin := splitHexChunks(str).toBinary().String()
	decodingTree := newDecodingTree(newEncodingTable())
	
	return unescapeUpper(decodingTree.decodeBinary(bin))
}

// escapeUpper escape upper case chars:
// changes <upper case letter> to ! + <lower case letter>.
//
// i.g.: My name is Ted -> !my name is !ted
func escapeUpper(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

// unescapeUpper unescape upper case chars:
// changes <! + lower case letter> to <upper case letter> .
//
// It opposite to escapeUpper.
//
// i.g.: !my name is !ted -> My name is Ted
func unescapeUpper(str string) string {
	var (
		buf        strings.Builder
		capitalize = false
	)
	for _, ch := range str {
		if capitalize {
			buf.WriteRune(unicode.ToUpper(ch))
			capitalize = false

			continue
		}

		if ch == '!' {
			capitalize = true

			continue
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

// encodeBinary encode string into binary codes string without spaces.
func encodeBinary(str string) string {
	var buf strings.Builder

	for _, char := range str {
		buf.WriteString(toBinary(char))
	}

	return buf.String()
}

func toBinary(char rune) string {
	table := newEncodingTable()

	code, ok := table[char]
	if !ok {
		panic(fmt.Sprintf("unknown character %q", char))
	}

	return code
}

type encodingTable map[rune]string

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
