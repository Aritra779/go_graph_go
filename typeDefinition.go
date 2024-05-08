package graph

type Node struct {
	id        string
	data      any
	Neighbors map[string]*Node
}

type Graph struct {
	Nodes map[string]*Node
}
