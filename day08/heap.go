package day08

type PairHeap []Pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x any) {
	*h = append(*h, x.(Pair))
	// Trying to keep the size of the heap down by only storing the top X values
	// This doesn't really work
	// for len(*h) > HeapLimit+1 {
	// 	*h = (*h)[0 : len(*h)-1]
	// }
}

func (h *PairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Part 2 --

type BoxPairHeap []BoxPair

func (h BoxPairHeap) Len() int           { return len(h) }
func (h BoxPairHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h BoxPairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BoxPairHeap) Push(x any) {
	*h = append(*h, x.(BoxPair))
}

func (h *BoxPairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type CircuitHeap []*Circuit

func (h CircuitHeap) Len() int           { return len(h) }
func (h CircuitHeap) Less(i, j int) bool { return len(h[i].boxes) < len(h[j].boxes) }
func (h CircuitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *CircuitHeap) Push(x any) {
	n := len(*h)
	circ := x.(*Circuit)
	circ.index = n
	*h = append(*h, circ)
}

func (h *CircuitHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h CircuitHeap) Peek() any {
	return h[0]
}
