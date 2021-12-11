package hamming

type Block struct {
	// length of block 
	length uint
	index  uint
	bit    uint
}

// Encode encodes data 
// Data must be valid string of 0's and 1's
// Example "00011101"
func Encode(data string) (string){
	
	return ""
}

func Decode(data string) (string, error)

