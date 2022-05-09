package vlc

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	chunkSize          = 8
	hexChunksSeparator = " "
)

type (
	binaryChunk  string
	binaryChunks []binaryChunk
)

// splitBinaryChunks splits binary string by chunks with given size.
//
// i.g.: '100101011001010110010101' -> '10010101 10010101 10010101'
func splitBinaryChunks(bString string, size int) binaryChunks {
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

type (
	hexChunk  string
	hexChunks []hexChunk
)

// splitHexChunks splits string to hex chunks.
//
// i.g.: '20 30 3C 18 77 4A E4 4D 28' -> '20', '30', '3C', '18', '77', '4A', 'E4', '4D', '28'
func splitHexChunks(str string) hexChunks {
	if str == "" {
		return hexChunks{}
	}
	chunks := strings.Split(str, hexChunksSeparator)

	hcs := make(hexChunks, 0, len(chunks))
	for _, c := range chunks {
		hcs = append(hcs, hexChunk(c))
	}

	return hcs
}

func (hcs hexChunks) toBinary() binaryChunks {
	chunks := make(binaryChunks, 0, len(hcs))

	for _, hc := range hcs {
		chunks = append(chunks, hc.toBinary())
	}

	return chunks
}

func (hc hexChunk) toBinary() binaryChunk {
	const hexNumberBase = 16
	num, err := strconv.ParseUint(string(hc), hexNumberBase, chunkSize)
	if err != nil {
		panic(fmt.Sprintf("can't parse hex chunk to number: %s", err.Error()))
	}

	chunk := fmt.Sprintf("%08b", num)

	return binaryChunk(chunk)
}
