package vlc

import "strings"

type decodingTree struct {
	char rune
	zero *decodingTree
	one  *decodingTree
}

func newDecodingTree(et encodingTable) *decodingTree {
	tree := new(decodingTree)

	for char, code := range et {
		tree.add(code, char)
	}

	return tree
}

func (dt *decodingTree) add(code string, char rune) {
	current := dt

	for _, bit := range code {
		switch bit {
		case '0':
			if current.zero == nil {
				current.zero = &decodingTree{}
			}
			current = current.zero
		case '1':
			if current.one == nil {
				current.one = &decodingTree{}
			}
			current = current.one
		}
	}

	current.char = char
}

func (dt *decodingTree) decodeBinary(bString string) string {
	var buf strings.Builder

	current := dt

	for _, bit := range bString {
		if current.char != rune(0) {
			buf.WriteRune(current.char)
			current = dt
		}

		switch bit {
		case '0':
			current = current.zero
		case '1':
			current = current.one
		}
	}

	if current.char != rune(0) {
		buf.WriteRune(current.char)
		current = dt
	}

	return buf.String()
}
