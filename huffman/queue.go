package huffman

type priorityQueue []*huffmanNode

type huffmanTree *huffmanNode

type huffmanNode struct {
	char rune
	freq uint

	left, right *huffmanNode
}

// Len returns length of a queue
func (q priorityQueue) Len() int {
	return len(q)
}

// Less returns true if i < j.
// Firstly checks frequencies, if they equal checks characters
func (pq priorityQueue) Less(i int, j int) bool {
	if pq[i].freq < pq[j].freq {
		return true
	}

	if pq[i].freq == pq[j].freq {
		return pq[i].char < pq[j].char
	}

	return false
}

// Swap changes positions of the elements
func (q priorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// Push inserts a new element in a queue
func (q *priorityQueue) Push(x interface{}) {
	*q = append(*q, x.(*huffmanNode))
}

// Pop deletes first element from the queue and returns its value
func (q *priorityQueue) Pop() interface{} {
	old := *q
	len := old.Len()
	toRemove := old[len-1]
	// avoid memory leek
	old[len-1] = nil

	*q = old[0 : len-1]
	return toRemove
}
