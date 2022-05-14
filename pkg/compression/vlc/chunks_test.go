package vlc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromBinaryString(t *testing.T) {
	t.Parallel()

	type args struct {
		bString string
		size    int
	}
	tests := []struct {
		name string
		str  string
		args args
		want binaryChunks
	}{
		{
			name: "split binary empty string to chunks with size 8",
			args: args{
				bString: "",
				size:    8,
			},
			want: binaryChunks{},
		},
		{
			name: "split binary string \"0010\" to chunks with size 8",
			args: args{
				bString: "0010",
				size:    8,
			},
			want: binaryChunks{"00100000"},
		},
		{
			name: "split binary string \"1111\" to chunks with size 8",
			args: args{
				bString: "1111",
				size:    8,
			},
			want: binaryChunks{"11110000"},
		},
		{
			name: fmt.Sprintf("split binary string %q to chunks with size %d", "Ted", 8),
			args: args{
				bString: "001000100110100101",
				size:    8,
			},
			want: binaryChunks{"00100010", "01101001", "01000000"},
		},
		{
			name: fmt.Sprintf("split binary string %q to chunks with size %d", "Ted", 4),
			args: args{
				bString: "001000100110100101",
				size:    4,
			},
			want: binaryChunks{"0010", "0010", "0110", "1001", "0100"},
		},
		{
			name: fmt.Sprintf("split binary string %q to chunks with size %d", "Ted", 6),
			args: args{
				bString: "001000100110100101",
				size:    6,
			},
			want: binaryChunks{"001000", "100110", "100101"},
		},
		{
			name: fmt.Sprintf("split binary string %q to chunks with size %d", "My name is Ted", 8),
			args: args{
				bString: "001000000011000000111100000110000111011101001010111001000100110100101",
				size:    8,
			},
			want: binaryChunks{
				"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101",
				"00101000",
			},
		},
		{
			name: fmt.Sprintf("split binary string %q to chunks with size %d", "Some pretty SUBsequence", 8),
			args: args{
				bString: "0010000101100010000111011100001010100010110011001000000111001000010100100000011001000000001001011010000000000010001110110000000101101",
				size:    8,
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
				fromBinaryString(test.args.bString, test.args.size),
				"fromBinaryString(%v, %v)", test.args.bString, test.args.size,
			)
		})
	}
}

func TestFromBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		str   string
		bytes []byte
		want  binaryChunks
	}{
		{str: "", bytes: []byte{}, want: binaryChunks{}},
		{
			str:   "Ted",
			bytes: []byte{0b00100010, 0b01101001, 0b01000000},
			want:  binaryChunks{"00100010", "01101001", "01000000"},
		},
		{
			str: "My name is Ted",
			bytes: []byte{
				0b00100000, 0b00110000, 0b00111100, 0b00011000, 0b01110111, 0b01001010, 0b11100100, 0b01001101,
				0b00101000,
			},
			want: binaryChunks{
				"00100000", "00110000", "00111100", "00011000", "01110111", "01001010", "11100100", "01001101",
				"00101000",
			},
		},
		{
			str: "Some pretty SUBsequence",
			bytes: []byte{
				0b00100001, 0b01100010, 0b00011101, 0b11000010, 0b10100010, 0b11001100, 0b10000001, 0b11001000,
				0b01010010, 0b00000110, 0b01000000, 0b00100101, 0b10100000, 0b00000010, 0b00111011, 0b00000001,
				0b01101000,
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
		test.name = fmt.Sprintf("convert bytes for string %q to chunks", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, test.want, fromBytes(test.bytes), "fromBytes(%v)", test.bytes)
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

func TestBinaryChunksBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		bcs  binaryChunks
		want []byte
	}{
		{str: "", bcs: binaryChunks{}, want: []byte{}},
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
		test.name = fmt.Sprintf("convert chunks for string %q to bytes", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			b, err := test.bcs.Bytes()
			assert.Equalf(t, test.want, b, "binaryChunks(%v).Bytes()", test.bcs)
			assert.Nil(t, err)
		})
	}
}
func TestBinaryChunksBytesError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		str   string
		bcs   binaryChunks
		error string
	}{
		{
			str:   "Ted",
			bcs:   binaryChunks{"00100010", "01101001", "100000011"},
			error: "can't parse binary chunk to number: strconv.ParseUint: parsing \"100000011\": value out of range for chunk \"100000011\"",
		},
		{
			str: "My name is Ted",
			bcs: binaryChunks{
				"00100000", "00110000", "00111100", "00011000", "01110111", "101010001", "01001010", "11100100",
				"01001101",
			},
			error: "can't parse binary chunk to number: strconv.ParseUint: parsing \"101010001\": value out of range for chunk \"101010001\"",
		},
	}

	for _, test := range tests {
		test := test
		test.name = fmt.Sprintf("convert chunks for string %q to bytes", test.str)
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			b, err := test.bcs.Bytes()
			assert.Emptyf(t, b, "binaryChunk(%v).Bytes() not empty result when error", test.bcs)
			assert.IsTypef(t, &ParseBinaryError{}, err, "binaryChunk(%v).Bytes() unexpected error type", test.bcs)
			assert.Equalf(t, test.error, err.Error(), "binaryChunk(%v).Bytes() unexpected error message", test.bcs)
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
			b, err := test.bc.Byte()
			assert.Equalf(t, test.want, b, "binaryChunk(%v).Byte()", test.bc)
			assert.Nil(t, err)
		})
	}
}

func TestBinaryChunkByteError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		bc    binaryChunk
		error string
	}{
		{
			name:  "lol",
			bc:    binaryChunk("101000011"),
			error: "can't parse binary chunk to number: strconv.ParseUint: parsing \"101000011\": value out of range for chunk \"101000011\""},
		{
			name:  "lol1",
			bc:    binaryChunk("111111110"),
			error: "can't parse binary chunk to number: strconv.ParseUint: parsing \"111111110\": value out of range for chunk \"111111110\"",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			b, err := test.bc.Byte()
			assert.Emptyf(t, b, "binaryChunk(%v).Byte() not empty result when error", test.bc)
			assert.IsTypef(t, &ParseBinaryError{}, err, "binaryChunk(%v).Byte() unexpected error type", test.bc)
			assert.Equalf(t, test.error, err.Error(), "binaryChunk(%v).Byte() unexpected error message", test.bc)
		})
	}
}
