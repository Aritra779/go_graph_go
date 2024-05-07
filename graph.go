package go_graph_go

import "errors"

type Node struct {
	id        string
	Neighbors map[string]*Node
}

type Graph struct {
	Nodes map[string]*Node
}

// Creates a new Graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

// Adds a neighbor to the node
func (node *Node) addNeighbor(neighborNode *Node) {
	node.Neighbors[neighborNode.id] = neighborNode
}

// removes a neighbor from a node
func (node *Node) removeNeighbor(neighborNode *Node) {
	delete(node.Neighbors, neighborNode.id)
}

// Adds A Node with No Neighbors
func (graph *Graph) AddNode(node *Node) error {
	if _, exists := graph.Nodes[node.id]; !exists {
		newNode := &Node{
			id:        node.id,
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

	node1.addNeighbor(node1)
	node2.addNeighbor(node2)
}

// Removes an Edge between two provided nodes
func (graph *Graph) RemoveEdge(node1Id, node2Id string) {
	node1 := graph.Nodes[node1Id]
	node2 := graph.Nodes[node2Id]

	node1.removeNeighbor(node2)
	node2.removeNeighbor(node1)
}
