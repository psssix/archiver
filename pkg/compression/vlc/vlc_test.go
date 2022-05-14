package vlc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodecPack(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, str string
		want      []byte
	}{
		{str: "", want: []byte{}},
		{str: "Ted", want: []byte{0b00100010, 0b01101001, 0b01000000}},
		{
			str: "My name is Ted",
			want: []byte{
				0b00100000, 0b00110000, 0b00111100, 0b00011000, 0b01110111, 0b01001010, 0b11100100, 0b01001101,
				0b00101000,
			},
		},
		{
			str: "Some pretty SUBsequence",
			want: []byte{
				0b00100001, 0b01100010, 0b00011101, 0b11000010, 0b10100010, 0b11001100, 0b10000001, 0b11001000,
				0b01010010, 0b00000110, 0b01000000, 0b00100101, 0b10100000, 0b00000010, 0b00111011, 0b00000001,
				0b01101000,
			},
		},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("packing %q", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			bytes, _ := New().Pack(test.str)
			assert.Equalf(t, test.want, bytes, "Codec.Pack(%v)", test.str)
		})
	}
}

func TestCodecPackUnknownCharacterError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, str, error string
	}{
		{str: "®", error: "encoding to binary error, unknown character '®'"},
		{str: "√", error: "encoding to binary error, unknown character '√'"},
		{str: "∫", error: "encoding to binary error, unknown character '∫'"},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("packing %q with error", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			bytes, err := New().Pack(test.str)
			assert.Nilf(t, bytes, "Codec.Pack(%v) not empty result when error", test.str)
			assert.IsTypef(t, &EncodingError{}, err, "Codec.Pack(%v) unexpected error type", test.str)
			assert.Equalf(t, test.error, err.Error(), "Codec.Pack(%v) unexpected error message", test.str)
		})
	}
}

func TestCodecUnpacking(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		bytes []byte
		want  string
	}{
		{bytes: []byte{}, want: ""},
		{bytes: []byte{0b00100010, 0b01101001, 0b01000000}, want: "Ted"},
		{
			bytes: []byte{
				0b00100000, 0b00110000, 0b00111100, 0b00011000, 0b01110111, 0b01001010, 0b11100100, 0b01001101,
				0b00101000,
			},
			want: "My name is Ted",
		},
		{
			bytes: []byte{
				0b00100001, 0b01100010, 0b00011101, 0b11000010, 0b10100010, 0b11001100, 0b10000001, 0b11001000,
				0b01010010, 0b00000110, 0b01000000, 0b00100101, 0b10100000, 0b00000010, 0b00111011, 0b00000001,
				0b01101000,
			},
			want: "Some pretty SUBsequence",
		},
	}
	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("unpacking %q", test.want)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			str, _ := New().Unpack(test.bytes)
			assert.Equalf(t, test.want, str, "Codec.Unpack(%v)", test.bytes)
		})
	}
}

func TestEscapeUpper(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, str, want string
	}{
		{str: "", want: ""},
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
		name, str, want string
	}{
		{str: "", want: ""},
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
