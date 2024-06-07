package graph

import "sync"

type Node struct {
	Id            string
	Data          any
	Neighbors     map[string]*Node
	NeighborMutex sync.RWMutex
}

type Graph struct {
	Nodes map[string]*Node
	mutex sync.RWMutex
}
