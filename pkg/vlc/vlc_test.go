package vlc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		want string
	}{
		{str: "Ted", want: "22 69 40"},
		{str: "My name is Ted", want: "20 30 3C 18 77 4A E4 4D 28"},
		{str: "Some pretty SUBsequence", want: "21 62 1D C2 A2 CC 81 C8 52 06 40 25 A0 02 3B 01 68"},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("pack %q", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, Encode(test.str), "Encode(%v)", test.str)
		})
	}
}

func TestEscapeUpper(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		want string
	}{
		{str: "Ted", want: "!ted"},
		{str: "My name is Ted", want: "!my name is !ted"},
		{str: "Some pretty SUBsequence", want: "!some pretty !s!u!bsequence"},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("escape upper case chars in %q", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, escapeUpper(test.str), "escapeUpper(%v)", test.str)
		})
	}
}

func TestUnescapeUpper(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		want string
	}{
		{str: "!ted", want: "Ted"},
		{str: "!my name is !ted", want: "My name is Ted"},
		{str: "!some pretty !s!u!bsequence", want: "Some pretty SUBsequence"},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("unescape upper case chars for %q", test.want)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, unescapeUpper(test.str), "unescapeUpper(%v)", test.str)
		})
	}
}

func TestEncodeBinary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		want string
	}{
		{str: "!ted", want: "001000100110100101"},
		{str: "!my name is !ted", want: "001000000011000000111100000110000111011101001010111001000100110100101"},
		{
			str:  "!some pretty !s!u!bsequence",
			want: "0010000101100010000111011100001010100010110011001000000111001000010100100000011001000000001001011010000000000010001110110000000101101",
		},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("encoding %q", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, encodeBinary(test.str), "encodeBinary(%v)", test.str)
		})
	}
}
