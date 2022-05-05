package vlc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		want string
	}{
		{str: "My name is Ted", want: "!my name is !ted"},
		{str: "Some pretty SUBsequence", want: "!some pretty !s!u!bsequence"},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("preparing %q", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, prepareText(test.str), "prepareText(%v)", test.str)
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
		{
			str:  "!my name is !ted",
			want: "001000000011000000111100000110000111011101001010111001000100110100101",
		},
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

func TestSplitChunks(t *testing.T) {
	t.Parallel()

	type args struct {
		bin  string
		size int
	}
	tests := []struct {
		name string
		str  string
		args args
		want binaryChunks
	}{
		{
			str: "Ted",
			args: args{
				bin:  "001000100110100101",
				size: 8,
			},
			want: binaryChunks{"00100010", "01101001", "01000000"},
		},
		{
			str: "Ted",
			args: args{
				bin:  "001000100110100101",
				size: 4,
			},
			want: binaryChunks{"0010", "0010", "0110", "1001", "0100"},
		},
		{
			str: "Ted",
			args: args{
				bin:  "001000100110100101",
				size: 6,
			},
			want: binaryChunks{"001000", "100110", "100101"},
		},
		{
			str: "My name is Ted",
			args: args{
				bin:  "001000000011000000111100000110000111011101001010111001000100110100101",
				size: 8,
			},
			want: binaryChunks{
				"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101",
				"00101000",
			},
		},
		{
			str: "Some pretty SUBsequence",
			args: args{
				bin:  "0010000101100010000111011100001010100010110011001000000111001000010100100000011001000000001001011010000000000010001110110000000101101",
				size: 8,
			},
			want: binaryChunks{
				"00100001", "01100010", "00011101", "11000010", "10100010", "11001100", "10000001", "11001000",
				"01010010", "00000110", "01000000", "00100101", "10100000", "00000010", "00111011", "00000001",
				"01101000",
			},
		},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("split binary for string %q with size %d", test.str, test.args.size)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(
				t,
				test.want,
				splitChunks(test.args.bin, test.args.size),
				"splitChunks(%v, %v)", test.args.bin, test.args.size,
			)
		})
	}
}

func TestBinaryChunkToHex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		bc   binaryChunk
		want hexChunk
	}{
		{bc: binaryChunk("00100001"), want: "21"},
		{bc: binaryChunk("01100010"), want: "62"},
		{bc: binaryChunk("00011101"), want: "1D"},
		{bc: binaryChunk("11000010"), want: "C2"},
		{bc: binaryChunk("10100010"), want: "A2"},
		{bc: binaryChunk("11001100"), want: "CC"},
		{bc: binaryChunk("10000001"), want: "81"},
		{bc: binaryChunk("11001000"), want: "C8"},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("convert chunk %q to hex", test.bc)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, test.bc.toHex(), "toHex(%v)()", test.bc)
		})
	}
}

func TestBinaryChunksToHex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		bcs  binaryChunks
		want hexChunks
	}{
		{str: "Ted", bcs: binaryChunks{"00100010", "01101001", "01000000"}, want: hexChunks{"22", "69", "40"}},
		{
			str: "My name is Ted",
			bcs: binaryChunks{
				"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101",
				"00101000",
			},
			want: hexChunks{"20", "30", "3C", "18", "77", "4A", "E4", "4D", "28"},
		},
		{
			str: "Some pretty SUBsequence",
			bcs: binaryChunks{
				"00100001", "01100010", "00011101", "11000010", "10100010", "11001100", "10000001", "11001000",
				"01010010", "00000110", "01000000", "00100101", "10100000", "00000010", "00111011", "00000001",
				"01101000",
			},
			want: hexChunks{
				"21", "62", "1D", "C2", "A2", "CC", "81", "C8", "52", "06", "40", "25", "A0", "02", "3B", "01", "68",
			},
		},
	}
	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("convert chunks set for string %q to hex", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, test.bcs.toHex(), "toHex(%v)()", test.bcs)
		})
	}
}

func TestHexChunksToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		hcs  hexChunks
		want string
	}{
		{
			name: "stringify empty hex chunks",
			hcs:  hexChunks{},
			want: "",
		},
		{
			name: "stringify one hex chunks",
			hcs:  hexChunks{"22"},
			want: "22",
		},
		{
			name: fmt.Sprintf("stringify hex chunks set for string %q", "Ted"),
			hcs:  hexChunks{"22", "69", "40"},
			want: "22 69 40",
		},
		{
			name: fmt.Sprintf("stringify hex chunks set for string %q", "My name is Ted"),
			hcs:  hexChunks{"20", "30", "3C", "18", "77", "4A", "E4", "4D", "28"},
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
		{
			name: fmt.Sprintf("stringify hex chunks set for string %q", "Some pretty SUBsequence"),
			hcs: hexChunks{
				"21", "62", "1D", "C2", "A2", "CC", "81", "C8", "52", "06", "40", "25", "A0", "02", "3B", "01", "68",
			},
			want: "21 62 1D C2 A2 CC 81 C8 52 06 40 25 A0 02 3B 01 68",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, test.hcs.String(), "String(%v)", test.hcs)
		})
	}
}

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
