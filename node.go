package graph

type Node struct {
	id        string
	Neighbors map[string]*Node
}

// Adds a neighbor to the node
func (node *Node) addNeighbor(neighborNode *Node) {
	node.Neighbors[neighborNode.id] = neighborNode
}

// removes a neighbor from a node
func (node *Node) removeNeighbor(neighborNode *Node) {
	delete(node.Neighbors, neighborNode.id)
}

func (node *Node) isNeighbor(neighborNodeId string) bool {
	_, ok := node.Neighbors[neighborNodeId]
	return ok
}
