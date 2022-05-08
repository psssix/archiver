package vlc

import (
	"strings"
	"unicode"
)

func Pack(str string) string {
	bin := encodeBinary(escapeUpper(str))
	return splitBinaryChunks(bin, chunkSize).toHex().String()
}

func Unpack(str string) string {
	bin := splitHexChunks(str).toBinary().String()
	tree := newDecodingTree(newEncodingTable())
	return unescapeUpper(tree.decodeBinary(bin))
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
