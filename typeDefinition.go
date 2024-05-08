package graph

type Node struct {
	Id        string
	Data      any
	Neighbors map[string]*Node
}

type Graph struct {
	Nodes map[string]*Node
}
