package dsheap

import "fmt"

/*
	Max Heap
	Complexity: O(log n)

	Features:
		Heapify: O(h) == O(log n)
			Insert
			Remove

	How it works:
		The max heap is a representation of an ordered 1d-array
		Can be forged into whatever is needed
		Element's key should be UNIQUE
*/

const dsName = "[MAXHEAP]"

type MaxHeap struct {
	data []int
}

//Insert Adds the obtained key to the heap. Returns an error if the Heapifying process has failed
func (h *MaxHeap) Insert(key int) error {
	const fnName = "[INSERT]"
	/* P1: Append key */
	h.data = append(h.data, key)
	/* P2: Heapify */
	if err := h._HeapifyDataUp(len(h.data) - 1); err != nil {
		return fmt.Errorf("%+v%+v > %w", dsName, fnName, err)
	}

	return nil
}

//Top Retrieves and returns the top node of the heap (max key value). Returns [int(0), error("-1")] if the Heap length is 0
func (h *MaxHeap) Top() (int, error) {
	const fnName = "[TOP]"

	if len(h.data) == 0 {
		return 0, fmt.Errorf("-1")
	}

	lastIndex := len(h.data) - 1
	top := h.data[0]
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]

	if err := h._HeapifyDataDown(0); err != nil {
		return 0, fmt.Errorf("%+v%+v > %w", dsName, fnName, err)

	}

	return top, nil
}

func (h *MaxHeap) _HeapifyDataUp(index int) error {
	/* Custom Helpers */
	getParent := func(i int) int {
		return (i - 1) / 2
	}

	/* Process */
	for h.data[getParent(index)] < h.data[index] {
		h._SwapIndexes(getParent(index), index)
		index = getParent(index)
	}

	return nil
}

func (h *MaxHeap) _HeapifyDataDown(index int) error {
	/* Custom Helpers */

	left := func(i int) int {
		return (2 * i) + 1
	}
	right := func(i int) int {
		return (2 * i) + 2
	}

	/* Process */

	lastIndex := len(h.data) - 1

	leftIndex, rightIndex := left(index), right(index)

	childToCompareIndex := 0

	for leftIndex <= lastIndex {
		if leftIndex == lastIndex {
			//when the leftIndex is the only child
			childToCompareIndex = leftIndex
		} else if h.data[leftIndex] > h.data[rightIndex] {
			//when the leftIndex is larger
			childToCompareIndex = leftIndex
		} else {
			//when the rightIndex is larger
			childToCompareIndex = rightIndex
		}

		if h.data[index] < h.data[childToCompareIndex] {
			h._SwapIndexes(index, childToCompareIndex)
			index = childToCompareIndex
			leftIndex, rightIndex = left(index), right(index)
		} else {
			return nil
		}

	}

	return nil
}

func (h *MaxHeap) _SwapIndexes(iA, iB int) {
	h.data[iA], h.data[iB] = h.data[iB], h.data[iA]
}
