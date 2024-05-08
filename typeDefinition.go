package graph

type Node struct {
	id        string
	Neighbors map[string]*Node
}

type Graph struct {
	Nodes map[string]*Node
}
