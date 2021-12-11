package main

import (
	"fmt"
	"strconv"

	"github.com/2thousandmax/compression-algorithms/lzw"
)

const (
	lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	hello = "Hello, World!"
	test  = "thisisthe"
)

func main() {
	originalString := test
	encodedString := lzw.Encode(originalString)
	fmt.Println("Original string: ", originalString)
	fmt.Println("Encoded string:  ", string(encodedString))
	fmt.Println("Original string: ", []int32(originalString))
	fmt.Println("Encoded string:  ", encodedString)
	
	decodedString, _ := lzw.Decode(encodedString)
	fmt.Println("Decoded string:  ", decodedString)
	
	fmt.Println("Compression ratio: ", fmt.Sprintf("%.f%%", compressionRatio([]int32(originalString), encodedString)))
}

func compressionRatio(originalString, encodedString []int32) float64 {
	ol, el := 0, 0
	for _, i := range originalString {
		ol += len(strconv.FormatInt(int64(i), 2))
	}

	for _, i := range encodedString {
		el += len(strconv.FormatInt(int64(i), 2))
	}

	return float64(el) / float64(ol) * 100
}
