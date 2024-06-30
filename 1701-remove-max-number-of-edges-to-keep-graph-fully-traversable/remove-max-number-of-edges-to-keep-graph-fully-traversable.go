func maxNumEdgesToRemove(n int, edges [][]int) int {
	result := len(edges)
	alice := NewUnionFind(n)
	bob := NewUnionFind(n)

	for _, edge := range edges {
		x, y := edge[1]-1, edge[2]-1
		if edge[0] == 3 && (!alice.InSameSet(x, y) || !bob.InSameSet(x, y)) {
			alice.Union(x, y)
			bob.Union(x, y)
			result--
		}
	}

	for _, edge := range edges {
		x, y := edge[1]-1, edge[2]-1
		if edge[0] == 1 {
			if alice.Union(x, y) {
				result--
			}
		}
		if edge[0] == 2 {
			if bob.Union(x, y) {
				result--
			}
		}
	}
	
	if alice.setCount > 1 || bob.setCount > 1 {
		return -1
	}

	return result
}

type UnionFind struct {
	parent   []int
	rank     []int
	setCount int
}

func NewUnionFind(n int) *UnionFind {
	p := make([]int, n)
	r := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = i
		r[i] = 1
	}
	return &UnionFind{
		parent:   p,
		rank:     r,
		setCount: n,
	}
}

func (uf *UnionFind) Find(x int) int {
	for x != uf.parent[x] {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UnionFind) Union(x, y int) bool {
	px := uf.Find(x)
	py := uf.Find(y)

	if px == py {
		return false
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.rank[px] += uf.rank[py]
	uf.parent[py] = px
	uf.setCount--
	return true
}

func (uf *UnionFind) InSameSet(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}