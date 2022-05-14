package vlc

import (
	"strings"
	"unicode"
)

type Codec struct{}

func New() Codec {
	return Codec{}
}

func (_ Codec) Pack(str string) ([]byte, error) {
	bString, err := encodeBinary(escapeUpper(str))
	if err != nil {
		return nil, err
	}
	return fromBinaryString(bString).Bytes(), nil
}

func (_ Codec) Unpack(bytes []byte) (string, error) {
	bString := fromBytes(bytes).String()
	tree := newDecodingTree(newEncodingTable())
	return unescapeUpper(tree.decodeBinary(bString)), nil
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
