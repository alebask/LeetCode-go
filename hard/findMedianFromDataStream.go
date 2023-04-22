package hard

type MedianFinder struct {
	minHeap *heap
	maxHeap *heap
}

func Constructor() MedianFinder {
	minHeap := newHeap(func(a, b int) bool { return a > b })
	maxHeap := newHeap(func(a, b int) bool { return a < b })
	return MedianFinder{minHeap: minHeap, maxHeap: maxHeap}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.minHeap.size() > 0 && num < mf.minHeap.peek() {
		mf.maxHeap.push(num)
		if mf.maxHeap.size()-mf.minHeap.size() > 1 {
			mf.minHeap.push(mf.maxHeap.pop())
		}
	} else {
		mf.minHeap.push(num)
		if mf.minHeap.size()-mf.maxHeap.size() > 1 {
			mf.maxHeap.push(mf.minHeap.pop())
		}
	}

}

func (mf *MedianFinder) FindMedian() float64 {

	if mf.minHeap.size() > mf.maxHeap.size() {
		return float64(mf.minHeap.peek())
	} else if mf.minHeap.size() < mf.maxHeap.size() {
		return float64(mf.maxHeap.peek())
	} else {
		return (float64(mf.maxHeap.peek()) + float64(mf.minHeap.peek())) / 2
	}
}

type heap struct {
	items []int
	less  func(a, b int) bool
}

func newHeap(less func(a, b int) bool) *heap {
	return &heap{
		items: make([]int, 0),
		less:  less,
	}
}

func (h *heap) push(v int) {
	h.items = append(h.items, v)
	h.heapifyUp((len(h.items) - 1))
}

func (h *heap) heapifyUp(i int) {
	for h.less(h.items[parentIndex(i)], h.items[i]) {
		h.swap(i, parentIndex(i))
		i = parentIndex(i)
	}
}

func (h *heap) pop() int {
	val := h.items[0]
	h.swap(0, h.size()-1)
	h.items = h.items[:h.size()-1]
	h.heapifyDown(0)
	return val
}

func (h *heap) heapifyDown(i int) {
	l, r := leftChildIndex(i), rightChildIndex(i)
	largest := i
	if l < len(h.items) && h.less(h.items[i], h.items[l]) {
		largest = l
	}
	if r < len(h.items) && h.less(h.items[largest], h.items[r]) {
		largest = r
	}
	if largest != i {
		h.swap(i, largest)
		h.heapifyDown(largest)
	}
}

func (h *heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *heap) peek() int {
	return h.items[0]
}

func (h *heap) size() int {
	return len(h.items)
}
func leftChildIndex(i int) int {
	return 2*i + 1
}
func rightChildIndex(i int) int {
	return 2*i + 2
}
func parentIndex(i int) int {
	return (i - 1) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
