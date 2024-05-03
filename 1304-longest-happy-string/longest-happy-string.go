func longestDiverseString(a int, b int, c int) string {
	h := &pairHeap{}
	if a > 0 {
		heap.Push(h, pair{'a', a})
	}
	if b > 0 {
		heap.Push(h, pair{'b', b})
	}
	if c > 0 {
		heap.Push(h, pair{'c', c})
	}

	var result []byte
	for h.Len() > 0 {
		m1 := heap.Pop(h).(pair)
		if len(result) == 0 {
			n := min(m1.n, 2)
			for i := 0; i < n; i++ {
				result = append(result, m1.name)
			}
			m1.n -= n
		} else if h.Len() > 0 && len(result) > 0 && result[len(result)-1] == m1.name {
			m2 := heap.Pop(h).(pair)
			result = append(result, m2.name)
			m2.n -= 1
			if m2.n > 0 {
				heap.Push(h, m2)
			}
		} else if len(result) > 0 && result[len(result)-1] == m1.name && result[len(result)-2] == m1.name && h.Len() == 0 {
			break
		} else {
			n := min(m1.n, 2)
			for i := 0; i < n; i++ {
				result = append(result, m1.name)
			}
			m1.n -= n
		}
		if m1.n > 0 {
			heap.Push(h, m1)
		}
	}

	return string(result)
}

type pairHeap []pair

func (h pairHeap) Len() int { return len(h) }

func (h pairHeap) Less(i, j int) bool { return h[i].n > h[j].n }
func (h pairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pairHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}
func (h *pairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type pair struct {
	name byte
	n    int
}