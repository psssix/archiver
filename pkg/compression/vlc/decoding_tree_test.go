package vlc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDecodingTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		et   encodingTable
		want *decodingTree
	}{
		{
			name: "test with 3 symbols",
			et: encodingTable{
				'a': "11",
				'b': "1001",
				'c': "0101",
			},
			want: &decodingTree{
				zero: &decodingTree{
					one: &decodingTree{
						zero: &decodingTree{
							one: &decodingTree{char: 'c'},
						},
					},
				},
				one: &decodingTree{
					zero: &decodingTree{
						zero: &decodingTree{
							one: &decodingTree{char: 'b'},
						},
					},
					one: &decodingTree{char: 'a'},
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, newDecodingTree(test.et), "newDecodingTree(%v)", test.et)
		})
	}
}

func TestDecodingTreeDecodeBinary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, bString, want string
	}{
		{bString: "", want: ""},
		{bString: "001000100110100101", want: "!ted"},
		{bString: "001000000011000000111100000110000111011101001010111001000100110100101", want: "!my name is !ted"},
		{
			bString: "0010000101100010000111011100001010100010110011001000000111001000010100100000011001000000001001011010000000000010001110110000000101101",
			want:    "!some pretty !s!u!bsequence",
		},
	}

	dt := newDecodingTree(newEncodingTable())

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("decoding %q", test.want)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, dt.decodeBinary(test.bString), "decodingTree(...).decodeBinary(%v)", test.bString)
		})
	}
}
