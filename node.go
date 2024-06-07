package graph

// Adds a neighbor to the node
func (node *Node) addNeighbor(neighborNode *Node) {
	node.NeighborMutex.Lock()
	defer node.NeighborMutex.Unlock()
	node.Neighbors[neighborNode.Id] = neighborNode
}

// removes a neighbor from a node
func (node *Node) removeNeighbor(neighborNode *Node) {
	node.NeighborMutex.Lock()
	defer node.NeighborMutex.Unlock()
	delete(node.Neighbors, neighborNode.Id)
}

func (node *Node) GetNeighbor(neighborNodeId *string) (*Node, bool) {
	node.NeighborMutex.RLock()
	defer node.NeighborMutex.RUnlock()
	neighborNode, ok := node.Neighbors[*neighborNodeId]
	return neighborNode, ok
}

func (node *Node) isNeighbor(neighborNodeId string) bool {
	_, ok := node.GetNeighbor(&neighborNodeId)
	return ok
}

func (node *Node) updateData(newData any) {
	node.Data = newData
}
