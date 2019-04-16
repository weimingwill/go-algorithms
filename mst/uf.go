package mst

// UnionFind implements union find disjoint-set data structure.
type UnionFind struct {
	parent []int  // parent[i] = parent of i
	rank   []byte // rank[i] = rank of subtree rooted at i (never more than 31)
	count  int    // number of components
}

// NewUnionFind constructor.
func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		parent: make([]int, n),
		rank:   make([]byte, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.rank[i] = 0
	}

	return u
}

// Find finds the root of p.
func (u *UnionFind) Find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]] // path compression by halving
		p = u.parent[p]
	}
	return p
}

// Union connects p and q.
func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}

	// make root of smaller rank point to root of larger rank
	if u.rank[rootP] < u.rank[rootQ] {
		u.parent[rootP] = rootQ
	} else if u.rank[rootP] > u.rank[rootQ] {
		u.parent[rootQ] = rootP
	} else {
		u.parent[rootQ] = rootP
		u.rank[rootP]++
	}
	u.count--
}

// Count returns number of components.
func (u *UnionFind) Count() int {
	return u.count
}

// Connected checks whether two input are connected.
func (u *UnionFind) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
