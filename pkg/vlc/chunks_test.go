package vlc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitBinaryChunks(t *testing.T) {
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
			name: "split binary from binary empty  binary expression with size 8",
			args: args{
				bin:  "",
				size: 8,
			},
			want: binaryChunks{},
		},
		{
			name: "split binary from binary expression \"0010\" with size 8",
			args: args{
				bin:  "0010",
				size: 8,
			},
			want: binaryChunks{"00100000"},
		},
		{
			name: "split binary from binary expression \"1111\" with size 8",
			args: args{
				bin:  "1111",
				size: 8,
			},
			want: binaryChunks{"11110000"},
		},
		{
			name: fmt.Sprintf("split binary for string %q with size %d", "Ted", 8),
			args: args{
				bin:  "001000100110100101",
				size: 8,
			},
			want: binaryChunks{"00100010", "01101001", "01000000"},
		},
		{
			name: fmt.Sprintf("split binary for string %q with size %d", "Ted", 4),
			args: args{
				bin:  "001000100110100101",
				size: 4,
			},
			want: binaryChunks{"0010", "0010", "0110", "1001", "0100"},
		},
		{
			name: fmt.Sprintf("split binary for string %q with size %d", "Ted", 6),
			args: args{
				bin:  "001000100110100101",
				size: 6,
			},
			want: binaryChunks{"001000", "100110", "100101"},
		},
		{
			name: fmt.Sprintf("split binary for string %q with size %d", "My name is Ted", 8),
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
			name: fmt.Sprintf("split binary for string %q with size %d", "Some pretty SUBsequence", 8),
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
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(
				t,
				test.want,
				splitBinaryChunks(test.args.bin, test.args.size),
				"splitBinaryChunks(%v, %v)", test.args.bin, test.args.size,
			)
		})
	}
}

func TestBinaryChunksString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		bcs  binaryChunks
		want string
	}{
		{
			name: "stringify empty binary chunks",
			bcs:  binaryChunks{},
			want: "",
		},
		{
			name: "stringify one binary chunks",
			bcs:  binaryChunks{"00000001"},
			want: "00000001",
		},
		{
			name: fmt.Sprintf("stringify binary chunks set for string %q", "Ted"),
			bcs:  binaryChunks{"00100010", "01101001", "01000000"},
			want: "001000100110100101000000",
		},
		{
			name: fmt.Sprintf("stringify binary chunks set for string %q", "My name is Ted"),
			bcs: binaryChunks{
				"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101",
				"00101000",
			},
			want: "001000000011000000111100000110000111011101001010111001000100110100101000",
		},
		{
			name: fmt.Sprintf("stringify binary chunks set for string %q", "Some pretty SUBsequence"),
			bcs: binaryChunks{
				"00100001", "01100010", "00011101", "11000010", "10100010", "11001100", "10000001", "11001000",
				"01010010", "00000110", "01000000", "00100101", "10100000", "00000010", "00111011", "00000001",
				"01101000",
			},
			want: "0010000101100010000111011100001010100010110011001000000111001000010100100000011001000000001001011010000000000010001110110000000101101000",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, test.bcs.String(), "binaryChunks(%v).String()", test.bcs)
		})
	}
}

func Test_binaryChunks_Bytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		bcs  binaryChunks
		want []byte
	}{
		{
			str:  "",
			bcs:  binaryChunks{},
			want: []byte{},
		},
		{
			str:  "Ted",
			bcs:  binaryChunks{"00100010", "01101001", "01000000"},
			want: []byte{0b00100010, 0b01101001, 0b01000000},
		},
		{
			str: "My name is Ted",
			bcs: binaryChunks{
				"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101",
				"00101000",
			},
			want: []byte{
				0b00100000, 0b00110000, 0b00111100, 0b00011000, 0b01110111, 0b01001010, 0b11100100, 0b01001101,
				0b00101000,
			},
		},
		{
			str: "Some pretty SUBsequence",
			bcs: binaryChunks{
				"00100001", "01100010", "00011101", "11000010", "10100010", "11001100", "10000001", "11001000",
				"01010010", "00000110", "01000000", "00100101", "10100000", "00000010", "00111011", "00000001",
				"01101000",
			},
			want: []byte{
				0b00100001, 0b01100010, 0b00011101, 0b11000010, 0b10100010, 0b11001100, 0b10000001, 0b11001000,
				0b01010010, 0b00000110, 0b01000000, 0b00100101, 0b10100000, 0b00000010, 0b00111011, 0b00000001,
				0b01101000,
			},
		},
	}
	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("convert chunks set for string %q to bytes", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, test.bcs.Bytes(), "binaryChunks(%v).Bytes()", test.bcs)
		})
	}
}

func TestBinaryChunkByte(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		bc   binaryChunk
		want byte
	}{
		{bc: binaryChunk("00100001"), want: byte(0b00100001)},
		{bc: binaryChunk("01100010"), want: byte(0b01100010)},
		{bc: binaryChunk("00011101"), want: byte(0b00011101)},
		{bc: binaryChunk("11000010"), want: byte(0b11000010)},
		{bc: binaryChunk("10100010"), want: byte(0b10100010)},
		{bc: binaryChunk("11001100"), want: byte(0b11001100)},
		{bc: binaryChunk("10000001"), want: byte(0b10000001)},
		{bc: binaryChunk("11001000"), want: byte(0b11001000)},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("convert chunk %q to bytes", test.bc)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, test.bc.Byte(), "binaryChunk(%v).Byte()", test.bc)
		})
	}
}

func TestSplitHexChunks(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		want hexChunks
	}{
		{
			name: "split hex chunks from string when empty chunk",
			str:  "",
			want: hexChunks{},
		},
		{
			name: "split hex chunks from string when one chunk",
			str:  "22",
			want: hexChunks{"22"},
		},
		{
			name: fmt.Sprintf("split hex chunks from string for expression %q", "Ted"),
			str:  "22 69 40",
			want: hexChunks{"22", "69", "40"},
		},
		{
			name: fmt.Sprintf("split hex chunks from string for expression %q", "My name is Ted"),
			str:  "20 30 3C 18 77 4A E4 4D 28",
			want: hexChunks{"20", "30", "3C", "18", "77", "4A", "E4", "4D", "28"},
		},
		{
			name: fmt.Sprintf("split hex chunks from string for expression %q", "Some pretty SUBsequence"),
			str:  "21 62 1D C2 A2 CC 81 C8 52 06 40 25 A0 02 3B 01 68",
			want: hexChunks{
				"21", "62", "1D", "C2", "A2", "CC", "81", "C8", "52", "06", "40", "25", "A0", "02", "3B", "01", "68",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, splitHexChunks(test.str), "splitHexChunks(%v)", test.str)
		})
	}
}

func TestHexChunksToBinary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		hcs  hexChunks
		want binaryChunks
	}{
		{str: "Ted", hcs: hexChunks{"22", "69", "40"}, want: binaryChunks{"00100010", "01101001", "01000000"}},
		{
			str: "My name is Ted",
			hcs: hexChunks{"20", "30", "3C", "18", "77", "4A", "E4", "4D", "28"},
			want: binaryChunks{
				"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101",
				"00101000",
			},
		},
		{
			str: "Some pretty SUBsequence",
			hcs: hexChunks{
				"21", "62", "1D", "C2", "A2", "CC", "81", "C8", "52", "06", "40", "25", "A0", "02", "3B", "01", "68",
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
		test.name = fmt.Sprintf("convert chunks set for string %q to binary", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, test.hcs.toBinary(), "hexChunks(%v).toBinary()", test.hcs)
		})
	}
}

func TestHexChunkToBinary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		hc   hexChunk
		want binaryChunk
	}{
		{hc: hexChunk("7F"), want: binaryChunk("01111111")},
		{hc: hexChunk("80"), want: binaryChunk("10000000")},
		{hc: hexChunk("21"), want: binaryChunk("00100001")},
		{hc: hexChunk("62"), want: binaryChunk("01100010")},
		{hc: hexChunk("1D"), want: binaryChunk("00011101")},
		{hc: hexChunk("C2"), want: binaryChunk("11000010")},
		{hc: hexChunk("A2"), want: binaryChunk("10100010")},
		{hc: hexChunk("CC"), want: binaryChunk("11001100")},
		{hc: hexChunk("81"), want: binaryChunk("10000001")},
		{hc: hexChunk("C8"), want: binaryChunk("11001000")},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("convert chunk %q to binary", test.hc)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, test.hc.toBinary(), "hexChunk(%v).toBinary()", test.hc)
		})
	}
}
