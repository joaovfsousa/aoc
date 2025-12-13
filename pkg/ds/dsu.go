package ds

type DisjointSet[T comparable] struct {
	parent map[T]T
	rank   map[T]int
	size   map[T]int // size of each root's set
}

func NewDisjointSet[T comparable](items ...T) *DisjointSet[T] {
	ds := &DisjointSet[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
		size:   make(map[T]int),
	}
	for _, item := range items {
		ds.MakeSet(item)
	}
	return ds
}

func (ds *DisjointSet[T]) MakeSet(x T) {
	if _, exists := ds.parent[x]; !exists {
		ds.parent[x] = x
		ds.rank[x] = 0
		ds.size[x] = 1
	}
}

func (ds *DisjointSet[T]) Find(x T) T {
	// lazy init if unknown
	if _, ok := ds.parent[x]; !ok {
		ds.MakeSet(x)
		return x
	}
	// path compression
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x])
	}
	return ds.parent[x]
}

func (ds *DisjointSet[T]) Union(a, b T) {
	rootA := ds.Find(a)
	rootB := ds.Find(b)

	if rootA == rootB {
		return
	}

	// union by rank, and update sizes
	if ds.rank[rootA] < ds.rank[rootB] {
		ds.parent[rootA] = rootB
		ds.size[rootB] += ds.size[rootA]
	} else if ds.rank[rootA] > ds.rank[rootB] {
		ds.parent[rootB] = rootA
		ds.size[rootA] += ds.size[rootB]
	} else {
		ds.parent[rootB] = rootA
		ds.rank[rootA]++
		ds.size[rootA] += ds.size[rootB]
	}
}

// Connected returns true if a and b are in the same set.
func (ds *DisjointSet[T]) Connected(a, b T) bool {
	return ds.Find(a) == ds.Find(b)
}

// GetSize returns the size of the set containing x.
func (ds *DisjointSet[T]) GetSize(x T) int {
	root := ds.Find(x)
	return ds.size[root]
}

// AllSetSizes returns a map[root]size for every current root.
func (ds *DisjointSet[T]) AllSetSizes() map[T]int {
	result := make(map[T]int)
	// First, compress all paths to ensure parent pointers are up to date
	for node := range ds.parent {
		ds.Find(node)
	}
	// Now collect sizes only for actual roots (where parent[x] == x)
	for node := range ds.parent {
		if ds.parent[node] == node {
			result[node] = ds.size[node]
		}
	}
	return result
}

func (ds *DisjointSet[T]) CountSets() int {
	// First, compress all paths to ensure parent pointers are up to date
	for node := range ds.parent {
		ds.Find(node)
	}
	total := 0
	// Now count only actual roots (where parent[x] == x)
	for node := range ds.parent {
		if ds.parent[node] == node {
			total++
		}
	}
	return total
}
