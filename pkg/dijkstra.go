package pkg

import (
	"dijktras/config"
	"dijktras/internal"
)

type linkData struct {
	node string
	cost int
}

type LinkGraph struct {
	nodes map[string][]linkData
}

func (lg *LinkGraph) AddLinks(c config.Config) {
	lg.nodes[c.NodeStart] = append(lg.nodes[c.NodeStart], linkData{node: c.NodeEnd, cost: c.Cost})
	lg.nodes[c.NodeEnd] = append(lg.nodes[c.NodeEnd], linkData{node: c.NodeStart, cost: c.Cost})
}

func NewLinkGraph() *LinkGraph {
	return &LinkGraph{nodes: make(map[string][]linkData)}
}

func (lg *LinkGraph) GetLinks(node string) []linkData {
	return lg.nodes[node]
}

func (lg *LinkGraph) Dijkstra(start, end string) Result {
	heap := internal.NewHeap()
	heap.Insert(internal.Path{Cost: 0, Nodes: []string{start}})
	visited := make(map[string]bool)

	for heap.Length() > 0 {
		adjNode := heap.Extract()

		node := adjNode.Nodes[len(adjNode.Nodes)-1]

		if visited[node] {
			continue
		}

		if node == end {
			return Result{
				TotalCost: adjNode.Cost,
				Path:      adjNode.Nodes,
			}
		}

		for _, i := range lg.GetLinks(node) {
			if !visited[i.node] {
				heap.Insert(internal.Path{Cost: adjNode.Cost + i.cost, Nodes: append([]string{}, append(adjNode.Nodes, i.node)...)})
			}
		}
		visited[node] = true
	}
	return Result{}
}
