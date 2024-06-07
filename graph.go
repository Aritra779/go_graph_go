package graph

import (
	"errors"
	"sync"
)

// Creates a new Graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		mutex: sync.RWMutex{},
	}
}

// Adds A Node with No Neighbors
func (graph *Graph) AddNode(nodeId string, nodeData any) error {
	graph.mutex.Lock()
	defer graph.mutex.Unlock()
	if _, exists := graph.Nodes[nodeId]; !exists {
		newNode := &Node{
			Id:            nodeId,
			Data:          nodeData,
			Neighbors:     make(map[string]*Node),
			NeighborMutex: sync.RWMutex{},
		}
		graph.Nodes[newNode.Id] = newNode
		return nil
	}
	return errors.New("node Already Exists in the graph")
}

// Removes a node from the graph along with all the connected edges
func (graph *Graph) RemoveNode(node *Node) error {
	graph.mutex.Lock()
	defer graph.mutex.Unlock()
	if _, exists := graph.Nodes[node.Id]; !exists {
		return errors.New("trying to delete a non existing Node")
	}
	for _, neighbor := range graph.Nodes[node.Id].Neighbors {
		neighbor.removeNeighbor(node)
	}
	delete(graph.Nodes, node.Id)
	return nil
}

// Adds an Edge between the two provided nodes
func (graph *Graph) AddEdge(node1Id, node2Id string) {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()
	node1, node1Exists := graph.Nodes[node1Id]
	node2, node2Exists := graph.Nodes[node2Id]

	if node1Exists && node2Exists {
		if !node1.isNeighbor(node2Id) {
			node1.addNeighbor(node2)
		}
		if !node2.isNeighbor(node1Id) {
			node2.addNeighbor(node1)
		}
	}
}

// Removes an Edge between two provided nodes
func (graph *Graph) RemoveEdge(node1Id, node2Id string) {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()
	node1, node1Exists := graph.Nodes[node1Id]
	node2, node2Exists := graph.Nodes[node2Id]

	if node1Exists && node2Exists {
		if node1.isNeighbor(node2Id) {
			node1.removeNeighbor(node2)
		}
		if node1.isNeighbor(node2Id) {
			node2.removeNeighbor(node1)
		}
	}
}

func (graph *Graph) AreNodesAdjacent(node1Id, node2Id string) bool {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()
	node1, node1Exists := graph.Nodes[node1Id]

	if !node1Exists {
		return false
	}

	_, node2Exists := graph.Nodes[node2Id]
	if !node2Exists {
		return false
	}

	return node1.isNeighbor(node2Id)
}

func (graph *Graph) UpdateNodeData(nodeId *string, newData any) {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()
	node1, exists := graph.Nodes[*nodeId]
	if !exists {
		return
	}
	node1.updateData(newData)
}
