package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const chunkSize = 8

func Encode(str string) string {
	str = prepareText(str)
	chunks := splitChunks(encodeBinary(str), chunkSize)

	return chunks.toHex().String()
}

// prepareText prepares text to be fit for encode:
// changes upper case letter to: ! + lower case letter.
//
// i.g.: My name is Ted -> !my name is !ted
func prepareText(str string) string {
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

// encodeBinary encode string into binary codes string without spaces.
func encodeBinary(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(toBinary(ch))
	}

	return buf.String()
}

func toBinary(ch rune) string {
	table := newEncodingTable()

	bin, ok := table[ch]
	if !ok {
		panic(fmt.Sprintf("unknown character %q", ch))
	}

	return bin
}

type encodingTable map[rune]string

func newEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

type binaryChunk string

type binaryChunks []binaryChunk

// splitChunks splits binary string by chunks with given size.
//
// i.g.: '100101011001010110010101' -> '10010101 10010101 10010101'
func splitChunks(bin string, size int) binaryChunks {
	strLen := len(bin)
	count := strLen / size
	if strLen%size != 0 {
		count++
	}

	chunks := make(binaryChunks, 0, count)
	var buf strings.Builder

	for i, ch := range bin {
		buf.WriteRune(ch)

		if (i+1)%size == 0 {
			chunks = append(chunks, binaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		last := buf.String()
		last += strings.Repeat("0", size-len(last))
		chunks = append(chunks, binaryChunk(last))
	}

	return chunks
}

type hexChunk string

type hexChunks []hexChunk

func (bcs binaryChunks) toHex() hexChunks {
	chunks := make(hexChunks, 0, len(bcs))

	for _, bc := range bcs {
		chunks = append(chunks, bc.toHex())
	}

	return chunks
}

func (bc binaryChunk) toHex() hexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic(fmt.Sprintf("can't parse binary chunk: %s", err.Error()))
	}

	chunk := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(chunk) == 1 {
		chunk = "0" + chunk
	}

	return hexChunk(chunk)
}

func (hcs hexChunks) String() string {
	const separator = " "

	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))

	for _, hc := range hcs[1:] {
		buf.WriteString(separator)
		buf.WriteString(string(hc))
	}

	return buf.String()
}
