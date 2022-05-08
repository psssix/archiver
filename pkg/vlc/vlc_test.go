package vlc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPack(t *testing.T) {
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
		test.name = fmt.Sprintf("packing %q", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, Pack(test.str), "Pack(%v)", test.str)
		})
	}
}

func TestUnpacking(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		want string
	}{
		{str: "22 69 40", want: "Ted"},
		{str: "20 30 3C 18 77 4A E4 4D 28", want: "My name is Ted"},
		{str: "21 62 1D C2 A2 CC 81 C8 52 06 40 25 A0 02 3B 01 68", want: "Some pretty SUBsequence"},
	}
	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("unpacking %q", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, Unpack(test.str), "Unpack(%v)", test.str)
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
