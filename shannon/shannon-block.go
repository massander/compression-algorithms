package shannon

import (
	"fmt"
)

type ShannonBlockEncoder struct {
	// Raw input data
	data []byte
	// Frequency table for every character in input data
	frTable map[string]int
	// Probability table for two character block based on freuency table
	probTable map[string]float64
	// Codes store encoded char blocks
	codes map[string]string
}

// Constructor
func NewShannonBlockEncoder() *ShannonBlockEncoder {
	return &ShannonBlockEncoder{}
}

func (e *ShannonBlockEncoder) calculateFrequency() {
	for _, v := range e.data {
		e.frTable[string(v)] += 1
	}

	charCount := len(e.frTable)

	for char, fr := range e.frTable {
		for char2, fr2 := range e.frTable {
			block := string(char) + string(char2)
			e.probTable[block] = float64((fr / charCount) * (fr2 / charCount))
		}
	}
}

func sortMap(data map[string]float64) ([]string, []float64) {
	keys := make([]string, 0)
	values := make([]float64, 0)

	flag := false

	for k, v := range data {
		if !flag {
			keys = append(keys, k)
			values = append(values, v)
			flag = true
		}

		for idx, val := range values {
			if val < v {
				continue
			}
			keys = append(keys[:idx], append([]string{k}, keys[idx:]...)...)
			values = append(values[:idx], append([]float64{v}, values[idx:]...)...)
			break
		}
	}

	return keys, values
}

func (e *ShannonBlockEncoder) Encode(data []byte) ([]byte, error) {
	if !(len(data) > 0) {
		return nil, fmt.Errorf("No data to encode")
	}

	e.data = data

	e.calculateFrequency()
	fmt.Println("Probability Table:", e.probTable)

	keys, _ := sortMap(e.probTable)

	klen := len(keys)
	for i := 0; i < klen; i++ {
		e.codes[keys[i]] += "1"
		for j := range keys[i:] {
			e.codes[keys[j]] += "0"
		}
	}

	res := make([]byte, 0)
	for i := 0; i < len(e.data); i += 2 {
		
	}

	return nil, nil
}

func (n *ShannonBlockEncoder) Entropy() int {
	return 0
}
