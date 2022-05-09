package vlc

import (
	"fmt"
	"strconv"
	"strings"
)

const chunkSize = 8

type (
	binaryChunk  string
	binaryChunks []binaryChunk
)

// fromBinaryString splits binary string by chunks with given size.
//
// i.g.: '100101011001010110010101' -> '10010101 10010101 10010101'
func fromBinaryString(bString string, size int) binaryChunks {
	strLen := len(bString)
	count := strLen / size
	if strLen%size != 0 {
		count++
	}

	chunks := make(binaryChunks, 0, count)
	var buf strings.Builder

	for i, bit := range bString {
		buf.WriteRune(bit)

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

func fromBytes(bytes []byte) binaryChunks {
	chunks := make(binaryChunks, 0, len(bytes))

	for _, b := range bytes {
		chunk := binaryChunk(fmt.Sprintf("%08b", b))
		chunks = append(chunks, chunk)
	}

	return chunks
}

func (bcs binaryChunks) String() string {
	if len(bcs) == 0 {
		return ""
	}

	var buf strings.Builder

	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}

	return buf.String()
}

// Bytes join string representation of binary chunks info one lines and returns as byte slice
func (bcs binaryChunks) Bytes() []byte {
	bytes := make([]byte, 0, len(bcs))

	for _, bc := range bcs {
		bytes = append(bytes, bc.Byte())
	}

	return bytes
}

func (bc binaryChunk) Byte() byte {
	const binaryNumberBase = 2
	num, err := strconv.ParseUint(string(bc), binaryNumberBase, chunkSize)
	if err != nil {
		panic(fmt.Sprintf("can't parse binary chunk to number: %s", err.Error()))
	}

	return byte(num)
}
