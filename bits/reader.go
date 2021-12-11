package bits

import (
	"bytes"
	"fmt"
	"io"
)

type BitReader struct {
	reader io.ByteReader
	byte   byte
	offset byte
}

func New(r io.ByteReader) *BitReader {
	return &BitReader{r, 0, 0}
}

func (r *BitReader) ReadBit() (bool, error) {
	if r.offset == 8 {
		r.offset = 0
	}
	if r.offset == 0 {
		var err error
		if r.byte, err = r.reader.ReadByte(); err != nil {
			return false, err
		}
	}
	bit := (r.byte & (0x80 >> r.offset)) != 0
	r.offset++
	return bit, nil
}

func (r *BitReader) ReadUint(nbits int) (uint64, error) {
	var result uint64
	for i := nbits - 1; i >= 0; i-- {
		bit, err := r.ReadBit()
		if err != nil {
			return 0, err
		}
		if bit {
			result |= 1 << uint(i)
		}
	}
	return result, nil
}

func main() {
	data := []byte{0xAF, 0x89}

	fmt.Println("----demo----")
	r := New(bytes.NewBuffer(data))
	version, _ := r.ReadUint(4)
	headerLength, _ := r.ReadUint(4)
	flags, _ := r.ReadUint(3)
	fmt.Printf("version: %d\nheader length: %d\nflags: %b\n", version, headerLength, flags)

	fmt.Println("----demo----")
	r = New(bytes.NewBuffer(data))
	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			bit, _ := r.ReadBit()
			fmt.Printf("%b[%d] = %t\n", data[i], j, bit)
		}
	}
}
