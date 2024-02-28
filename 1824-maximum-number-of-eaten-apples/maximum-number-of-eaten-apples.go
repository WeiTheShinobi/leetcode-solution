func eatenApples(apples []int, days []int) int {
	result := 0
	currDay := 0
	h := &PairHeap{}

	n := len(apples)
	for h.Len() > 0 || currDay < n {
		if currDay < n {
			heap.Push(h, Pair{apples[currDay], days[currDay] + currDay})
		}
		currDay++
		if h.Len() > 0 {
			p := heap.Pop(h).(Pair)
			for (p.day < currDay || p.apple <= 0) && h.Len() > 0 {
				p = heap.Pop(h).(Pair)
			}
			if p.day < currDay || p.apple <= 0 {
				continue
			}
			if p.apple > 0 {
				p.apple--
				result++
				if p.apple > 0 {
					heap.Push(h, p)
				}
			}
		}
	}

	return result
}


type Pair struct {
	apple, day int
}

type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }

func (h PairHeap) Less(i, j int) bool { return h[i].day < h[j].day }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}