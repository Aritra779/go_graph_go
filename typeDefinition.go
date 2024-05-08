package graph

type Node struct {
	id        string
	data      struct{}
	Neighbors map[string]*Node
}

type Graph struct {
	Nodes map[string]*Node
}
