package huffman

import (
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/icza/bitio"
	"github.com/jedib0t/go-pretty/table"
)

// ABC -> 3 Byte -> 24 bits
// A:00 B:01 C:10 -> 000110 -> 8 bits

// An Encoder writes Huffman encoded something to an output stream.
type Encoder struct {
	w           io.Writer
	err         error
	data        string
	huffmanTree *huffmanNode
	queue       priorityQueue
	codeTable   map[rune]code
}

type code struct {
	size int
	code string // slice of 0 and 1
}

// NewEncoder returns a new Encoder which performs Huffman encoding.
// An encoded data is written to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w:           w,
		err:         nil,
		data:        "",
		huffmanTree: nil,
		queue:       make([]*huffmanNode, 0),
		codeTable:   make(map[rune]code),
	}
}

// Encode encodes string using Huffman coding algorithm.
func (e *Encoder) Encode(data string) error {
	e.data = data

	// get the frequency of each character
	freq := make(map[rune]uint)
	for _, char := range e.data {
		freq[char] += 1
	}

	e.buildHuffmanTree(freq)

	prefix := bytes.NewBufferString("")

	if err := e.traversHuffmanTreeAndWritePrefix(e.huffmanTree, prefix); err != nil {
		e.err = err
	}

	if e.err != nil {
		return e.err
	}

	e.printCodeTable()

	fmt.Println(e.codeTable)

	buf := bytes.NewBufferString("")
	e.writeEncodedBits(buf)

	e.w.Write(buf.Bytes())

	return nil
}

// buildhuffmanTree builds binary tree contains of huffmanNode
func (e *Encoder) buildHuffmanTree(charFrequency map[rune]uint) {
	// fill queue
	for char, freq := range charFrequency {
		treeNode := huffmanNode{
			char:  char,
			freq:  freq,
			left:  nil,
			right: nil,
		}
		e.queue = append(e.queue, &treeNode)
	}

	// Init heap
	heap.Init(&e.queue)

	for len(e.queue) > 1 {
		// tale away two elemnts from the queue
		x := heap.Pop(&e.queue).(*huffmanNode)
		y := heap.Pop(&e.queue).(*huffmanNode)

		// new element of the queue equals sum of two taken elemnts
		newNode := huffmanNode{
			char:  x.char + y.char,
			freq:  x.freq + y.freq,
			left:  x,
			right: y,
		}
		// add new elementto the queue
		heap.Push(&e.queue, &newNode)
	}

	e.huffmanTree = heap.Pop(&e.queue).(*huffmanNode)
}

// addBit adds new bit to the end of the prefix
func addBit(prefix *[]int, bit int) error {
	if bit != 1 && bit != 0 {
		return fmt.Errorf("bit can be 0 or 1")
	}

	*prefix = append(*prefix, bit)

	return nil
}

// removeBit removes last bit from the prefix
func removeBit(prefix *[]int) error {
	p := *prefix

	if len(p) == 0 {
		return fmt.Errorf("prefix is empty")
	}

	p = p[:len(p)-1]

	*prefix = p

	return nil
}

func (e *Encoder) traversHuffmanTreeAndWritePrefix(tree *huffmanNode, prefix *bytes.Buffer) error {
	if tree == nil {
		return fmt.Errorf("tree is empty")
	}

	if tree.left != nil || tree.right != nil {
		if tree.left != nil {
			// Add new bit to the prefix
			if _, err := prefix.WriteString("0"); err != nil {
				log.Fatalf("prefix.WriteString: %v", err)
			}

			e.traversHuffmanTreeAndWritePrefix(tree.left, prefix)

			// Remove last byte from the buffer
			prefix.Truncate(prefix.Len() - 1)
		}

		if tree.right != nil {
			// Add new bit to the prefix
			if _, err := prefix.WriteString("1"); err != nil {
				log.Fatalf("prefix.WriteString: %v", err)
			}

			e.traversHuffmanTreeAndWritePrefix(tree.right, prefix)

			// Remove last byte from the buffer
			prefix.Truncate(prefix.Len() - 1)
		}
	} else {
		// Add prefix to Huffman codes

		stringCode := prefix.String()

		e.codeTable[tree.char] = code{
			size: len(stringCode),
			code: stringCode,
		}
	}

	return nil
}

// writeEncodedBits writes encoded data to buf.
func (e *Encoder) writeEncodedBits(buf *bytes.Buffer) {
	w := bitio.NewWriter(buf)

	for _, char := range e.data {
		for _, bit := range e.codeTable[char].code {
			if string(bit) == "1" {
				w.WriteBool(true)
			} else {
				w.WriteBool(false)
			}
		}
	}

	w.Close()
}

func (e *Encoder) printCodeTable() string {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	body := make([]table.Row, 0)

	// for i, val := range []string{

	// }

	// mergeBits := func(bits []int) string {
	// 	var s string = ""
	// 	for _, bit := range bits {
	// 		if bit == 1 {
	// 			s += "1"
	// 		}
	// 		if bit == 0 {
	// 			s += "0"
	// 		}
	// 	}
	// 	return s
	// }

	for char, code := range e.codeTable {
		fmt.Println(code)
		row := table.Row{
			string(char),
			fmt.Sprintf("%b", char),
			code.code,
		}
		body = append(body, row)
	}

	t.AppendHeader(table.Row{"#", "ASCII", "Huffman"})
	t.AppendRows(body)

	return t.Render()
}
