package internal

type MinHeap struct {
	array []Path
}

type Path struct {
	Cost  int
	Nodes []string
}

func NewHeap() MinHeap {
	return MinHeap{}
}

func (h *MinHeap) Insert(p Path) {
	h.array = append(h.array, Path{Cost: p.Cost, Nodes: p.Nodes})
	h.minHeapifyUp(len(h.array) - 1)
}

func (h *MinHeap) Extract() Path {

	extracted := h.array[0]
	if len(h.array) == 0 {
		return Path{}
	}
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]
	h.minHeapifyDown(0)
	return extracted
}

func (h *MinHeap) minHeapifyUp(index int) {
	for h.array[parent(index)].Cost > h.array[index].Cost {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) minHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childtoCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childtoCompare = l
		} else if h.array[l].Cost < h.array[r].Cost {
			childtoCompare = l
		} else {
			childtoCompare = r
		}

		if h.array[index].Cost > h.array[childtoCompare].Cost {
			h.swap(index, childtoCompare)
			index = childtoCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (h *MinHeap) Length() int {
	return len(h.array)
}
