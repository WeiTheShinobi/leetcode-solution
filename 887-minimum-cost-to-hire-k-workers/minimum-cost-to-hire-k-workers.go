func mincostToHireWorkers(quality, wage []int, k int) float64 {
    type pair struct{ q, w int }
    pairs := make([]pair, len(quality))
    for i, q := range quality {
        pairs[i] = pair{q, wage[i]}
    }
    slices.SortFunc(pairs, func(a, b pair) int { return a.w*b.q - b.w*a.q }) // 按照 r 值排序

    h := hp{make([]int, k)}
    sumQ := 0
    for i, p := range pairs[:k] {
        h.IntSlice[i] = p.q
        sumQ += p.q
    }
    heap.Init(&h)

    ans := float64(sumQ*pairs[k-1].w) / float64(pairs[k-1].q) // 选 r 值最小的 k 名工人

    for _, p := range pairs[k:] { // 后面的工人 r 值更大
        if p.q < h.IntSlice[0] { // 但是 sumQ 可以变小，从而可能得到更优的答案
            sumQ -= h.IntSlice[0] - p.q
            h.IntSlice[0] = p.q
            heap.Fix(&h, 0) // 更新堆顶
            ans = min(ans, float64(sumQ*p.w)/float64(p.q))
        }
    }
    return ans
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
