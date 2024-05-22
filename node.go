package graph

// Adds a neighbor to the node
func (node *Node) addNeighbor(neighborNode *Node) {
	node.Neighbors[neighborNode.Id] = neighborNode
}

// removes a neighbor from a node
func (node *Node) removeNeighbor(neighborNode *Node) {
	delete(node.Neighbors, neighborNode.Id)
}

func (node *Node) isNeighbor(neighborNodeId string) bool {
	_, ok := node.Neighbors[neighborNodeId]
	return ok
}

func (node *Node) updateData(newData any) {
	node.Data = newData
}
