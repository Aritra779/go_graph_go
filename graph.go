package go_graph_go

type Node struct {
	id        string
	Neighbors map[string]*Node
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (node *Node) AddNeighbor(neighborNode *Node) {
	node.Neighbors[neighborNode.id] = neighborNode
}

func (node *Node) RemoveNeighbor(neighborNode *Node) {
	node.Neighbors[neighborNode.id] = nil
}

func (graph *Graph) AddNode(node *Node) {}

func (graph *Graph) RemoveNode(node *Node) {}

func (graph *Graph) AddEdge(node1Id, node2Id string) {}

func (graph *Graph) RemoveEdge(node1Id, node2Id string) {}