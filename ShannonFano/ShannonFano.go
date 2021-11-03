package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

type ShannonFanoCode struct {
	originalString          string
	originalStringLength    int
	compressedResult        map[rune]string
	compressedString        []string
	characterFrequency      map[rune]float64
	lengthBeforeCompression int
	lengthAfterCompression  int
}

var _ fmt.Stringer = (*ShannonFanoCode)(nil)

func newShannonFanoCode(str string) *ShannonFanoCode {
	return &ShannonFanoCode{
		originalString:          str,
		originalStringLength:    len(str),
		compressedResult:        make(map[rune]string),
		characterFrequency:      make(map[rune]float64),
		lengthBeforeCompression: 0,
		lengthAfterCompression:  0,
	}
}

func NewShannonFanoCode(str string) *ShannonFanoCode {
	var _ShannonFanoCode *ShannonFanoCode = newShannonFanoCode(str)

	_ShannonFanoCode.calculateFrequency()
	_ShannonFanoCode.compressString()
	_ShannonFanoCode.calculateLengthBeforeCompression()
	_ShannonFanoCode.calculateLengthAfterCompression()

	return _ShannonFanoCode
}

func NewShannonFanoCodeWihtProbability(str string, probability map[rune]float64) *ShannonFanoCode {
	if probability == nil {
		fmt.Println("Probability map is empty. Will try to create from string")
		return NewShannonFanoCode(str)
	}

	var _ShannonFanoCode *ShannonFanoCode = newShannonFanoCode(str)

	characterFrequency := make(map[rune]float64)

	var checkSum float64

	for _, char := range str {
		if _, ok := probability[char]; !ok {
			log.Fatalf("No such character %c in probability map\n", char)
		}

		checkSum += probability[char]

		characterFrequency[char] = probability[char] * float64(_ShannonFanoCode.originalStringLength)
	}

	if checkSum != float64(1) {
		log.Fatalln("Sum of probabilities does not mutch 1.0")
	}

	_ShannonFanoCode.compressString()
	_ShannonFanoCode.calculateLengthBeforeCompression()
	_ShannonFanoCode.calculateLengthAfterCompression()

	return _ShannonFanoCode
}

func (SHFCode *ShannonFanoCode) calculateFrequency() {
	_characterFrequency := SHFCode.characterFrequency

	for _, char := range SHFCode.originalString {
		if _, ok := _characterFrequency[char]; ok {
			_characterFrequency[char] += float64(1)
		} else {
			_characterFrequency[char] = float64(1)
		}
	}

	// Character probability
	// for _, char := range SHFCode.originalString {
	// 	_characterFrequency[char] /= float64(SHFCode.originalStringLength)
	// }

}

func (SHFCode *ShannonFanoCode) compressString() {
	_characterFrequency := SHFCode.characterFrequency

	length := SHFCode.originalStringLength

	charSlice := make([]rune, 0, length)
	freqSlice := make([]float64, 0, length)

	for char, freq := range _characterFrequency {

		if len(freqSlice) == 0 || freqSlice[len(freqSlice)-1] < freq {
			freqSlice = append(freqSlice, freq)
			charSlice = append(charSlice, char)
			continue
		}

		for i := 0; i < len(freqSlice); i++ {
			if freqSlice[i] >= freq {
				freqSlice = append(freqSlice[:i], append([]float64{freq}, freqSlice[i:]...)...)
				charSlice = append(charSlice[:i], append([]rune{char}, charSlice[i:]...)...)
				break
			}
		}

	}

	appendBit(SHFCode.compressedResult, charSlice, false)

	for _, char := range SHFCode.originalString {
		SHFCode.compressedString = append(SHFCode.compressedString, SHFCode.compressedResult[char])
	}
}

func appendBit(result map[rune]string, charSlice []rune, addBit bool) {
	var bit string = "0"

	if len(result) >= 0 {
		if addBit {
			bit = "1"
		}
	}

	for _, char := range charSlice {
		result[char] += bit
	}

	if len(charSlice) >= 2 {
		separator := len(charSlice) / 2
		if len(charSlice) != int(separator)*2 {
			separator++
		}

		appendBit(result, charSlice[:separator], false)

		appendBit(result, charSlice[separator:], true)
	}
}

func (SHFCode *ShannonFanoCode) calculateLengthBeforeCompression() {
	var result string
	for _, str := range SHFCode.originalString {
		result += fmt.Sprintf("%b", str)
	}

	SHFCode.lengthBeforeCompression = len(result)
}

func (SHFCode *ShannonFanoCode) calculateLengthAfterCompression() {
	SHFCode.lengthAfterCompression = len(strings.Join(SHFCode.compressedString, ""))
}

func (SHFCode *ShannonFanoCode) String() string {
	var buf bytes.Buffer

	buf.WriteString(
		fmt.Sprintf("Original string:  %s", SHFCode.originalString),
	)

	buf.WriteString(
		fmt.Sprintf("%s%s\n",
			strings.Repeat(" ", 18),
			func() string {
				var _buf bytes.Buffer

				for _, str := range SHFCode.originalString {
					_buf.WriteString(fmt.Sprintf("%b ", str))
				}
				buf.WriteString("\n")

				return _buf.String()
			}(),
		),
	)

	buf.WriteString(
		fmt.Sprintf(
			"Original string length: %v bits\n\n",
			SHFCode.lengthBeforeCompression,
		),
	)

	buf.WriteString(
		fmt.Sprintf(
			"Compressed strng: %s\n",
			strings.Join(SHFCode.compressedString, " "),
		),
	)

	// Print length of compressed string
	buf.WriteString(
		fmt.Sprintf(
			"Compressed string length: %v bits\n\n",
			SHFCode.lengthAfterCompression,
		),
	)

	buf.WriteString(
		fmt.Sprintf(
			"Compression ratio: %.4f (%.f%%)",
			float64(SHFCode.lengthBeforeCompression)/float64(SHFCode.lengthAfterCompression),
			float64(SHFCode.lengthAfterCompression)/float64(SHFCode.lengthBeforeCompression)*100,
		),
	)

	return buf.String()
}

func main() {
	testString := "shannon fano algorithm"
	
	fmt.Println(NewShannonFanoCode(testString))
}
