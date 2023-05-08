package graph

type Vertex struct {
	Label      int
	WasVisited bool
}

type Graph struct {
	vertices int
	Edge int
	Adj map[int][]int

}