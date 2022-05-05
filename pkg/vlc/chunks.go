package vlc

import (
	"fmt"
	"strconv"
	"strings"
)

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
	const binaryNumberBase = 2
	num, err := strconv.ParseUint(string(bc), binaryNumberBase, chunkSize)
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
