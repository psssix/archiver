package vlc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeBinary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, str, want string
	}{
		{str: "", want: ""},
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
			binary, err := encodeBinary(test.str)
			assert.Equalf(t, test.want, binary, "encodeBinary(%v)", test.str)
			assert.Nil(t, err)
		})
	}
}

func TestEncodingUnknownCharacterError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, str, error string
	}{
		{str: "π", error: "encoding to binary error, unknown character 'π'"},
		{str: "!ted Ω", error: "encoding to binary error, unknown character 'Ω'"},
		{str: "!my name is ∑", error: "encoding to binary error, unknown character '∑'"},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("expect error when encoding an unknown character %q", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			binary, err := encodeBinary(test.str)
			assert.Emptyf(t, binary, "encodeBinary(%v) not empty result when error", test.str)
			assert.IsTypef(t, &EncodingError{}, err, "encodeBinary(%v) unexpected error type", test.str)
			assert.Equalf(t, test.error, err.Error(), "encodeBinary(%v) unexpected error message", test.str)
		})
	}
}
