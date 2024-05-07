package go_graph_go

type Node struct {
	id        string
	Neighbors []*Node
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (node *Node) AddNeighbor(neighborNode *Node) {}

func (node *Node) RemoveNeighbor(neighborNode *Node) {}

func (graph *Graph) AddNode(node *Node) {}

func (graph *Graph) RemoveNode(node *Node) {}

func (graph *Graph) AddEdge(node1Id, node2Id string) {}

func (graph *Graph) RemoveEdge(node1Id, node2Id string) {}
