package day03

// Byte with an 'i' index - B i te
type Bite struct {
	i int
	b byte
}

type BiteHeap []Bite

func (h BiteHeap) Len() int           { return len(h) }
func (h BiteHeap) Less(i, j int) bool { return h[i].b > h[j].b } // Max heap
func (h BiteHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BiteHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Bite))
}

func (h *BiteHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
