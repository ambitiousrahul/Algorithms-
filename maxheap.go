package polynomialhash

import "fmt"

//WordsFrequency struct returns the words and its count.
type WordsFrequency struct {
	Word      string
	Frequency int
}

//Maxheap returns maxheap.
type Maxheap struct {
	heapArray []WordsFrequency
	size      int
	maxsize   int
}

//NewMaxheap returns new maxheap object.
func NewMaxheap(maxsize int) *Maxheap {
	maxheap := &Maxheap{
		heapArray: []WordsFrequency{},
		size:      0,
		maxsize:   maxsize,
	}
	return maxheap
}

//NewMaxheapWithElements returns max heap object with elements
func NewMaxheapWithElements(maxsize int, elements []WordsFrequency) *Maxheap {
	maxheap := &Maxheap{
		heapArray: elements,
		size:      len(elements),
		maxsize:   maxsize,
	}
	return maxheap
}

func (m *Maxheap) leftchild(index int) int {
	return 2*index + 1
}
func (m *Maxheap) rightchild(index int) int {
	return 2*index + 2
}
func (m *Maxheap) parent(index int) int {
	return (index - 1) / 2 //index of the parent
}

func (m *Maxheap) swap(first, second int) {
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second] = temp
}

func (m *Maxheap) maxHeapify(index int) {
	leftChildIndex := m.leftchild(index)
	rightChildIndex := m.rightchild(index)
	largest := 0
	if leftChildIndex < len(m.heapArray) && m.heapArray[leftChildIndex].Frequency > m.heapArray[index].Frequency {
		largest = leftChildIndex
	} else {
		largest = index
	}
	if rightChildIndex < len(m.heapArray) && m.heapArray[rightChildIndex].Frequency > m.heapArray[largest].Frequency {
		largest = rightChildIndex
	}
	if index != largest {
		m.swap(largest, index)
		m.maxHeapify(largest)
	}
}

//Maximum returns the top element of the max heap.
func (m *Maxheap) Maximum() WordsFrequency {
	return m.heapArray[0]
}

//ExtractMaximum pops the maximum element of the max heap.
func (m *Maxheap) ExtractMaximum() WordsFrequency {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:m.size-1]
	m.size--
	m.maxHeapify(0)
	return top
}

//BuildMaxHeap builds the max heap from the array of elements.
func (m *Maxheap) BuildMaxHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.maxHeapify(index)
	}
}

//Insert inserts an element in the heap.
func (m *Maxheap) Insert(item WordsFrequency) error {
	if m.size >= m.maxsize {
		return fmt.Errorf("Heal is ful")
	}
	m.heapArray = append(m.heapArray, item)
	m.size++

	parentNodeIndex := m.parent(m.size - 1)
	i := parentNodeIndex

	//if the parent node is smaller than the index node.
	for i > 0 && m.heapArray[parentNodeIndex].Frequency < item.Frequency {
		m.swap(i, parentNodeIndex)
		i = parentNodeIndex
	}
	return nil
}
