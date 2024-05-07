package graph

import (
	"errors"
)

type Graph struct {
	Nodes map[string]*Node
}

// Creates a new Graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

// Adds A Node with No Neighbors
func (graph *Graph) AddNode(nodeId string) error {
	if _, exists := graph.Nodes[nodeId]; !exists {
		newNode := &Node{
			id:        nodeId,
			Neighbors: make(map[string]*Node),
		}
		graph.Nodes[newNode.id] = newNode
		return nil
	}
	return errors.New("node Already Exists in the graph")
}

// Removes a node from the graph along with all the connected edges
func (graph *Graph) RemoveNode(node *Node) error {
	if _, exists := graph.Nodes[node.id]; !exists {
		return errors.New("trying to delete a non existing Node")
	}
	for _, neighbor := range graph.Nodes[node.id].Neighbors {
		neighbor.removeNeighbor(node)
	}
	delete(graph.Nodes, node.id)
	return nil
}

// Adds an Edge between the two provided nodes
func (graph *Graph) AddEdge(node1Id, node2Id string) {
	node1 := graph.Nodes[node1Id]
	node2 := graph.Nodes[node2Id]

	node1.addNeighbor(node2)
	node2.addNeighbor(node1)
}

// Removes an Edge between two provided nodes
func (graph *Graph) RemoveEdge(node1Id, node2Id string) {
	node1 := graph.Nodes[node1Id]
	node2 := graph.Nodes[node2Id]

	node1.removeNeighbor(node2)
	node2.removeNeighbor(node1)
}

func (graph *Graph) AreNodesAdjacent(node1Id, node2Id string) bool {
	return graph.Nodes[node1Id].isNeighbor(node2Id)
}
