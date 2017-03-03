package playground

import (
	"errors"
)

// procedure UniformCostSearch(Graph, start, goal)
//   node ← start
//   cost ← 0
//   frontier ← priority queue containing node only
//   explored ← empty set
//   do
//     if frontier is empty
//       return failure
//     node ← frontier.pop()
//     if node is goal
//       return solution
//     explored.add(node)
//     for each of node's neighbors n
//       if n is not in explored
//         if n is not in frontier
//           frontier.add(n)
//         else if n is in frontier with higher cost
//           replace existing node with n

type Frontier []*Vertex

func UniformCostSearch(g *Graph, start *Vertex, goal *Vertex) (map[string]*Vertex, error) {

	frontier := Frontier{}
	frontier = append(frontier, start)
	explored := make(map[string]*Vertex)

	for {
		length := len(frontier)
		if length == 0 {
			return nil, errGoalNoFound
		}

		node := frontier[length]
		if node == goal {
			break
		}

		explored[node.id] = node
		for _, e := range node.Edges() {
			n := g.Find(e.id)
			if _, ok := explored[n.id]; !ok {
				if !frontier.Contains(n.id) {
					frontier = append(frontier, n)
				}
			}
		}

	}

	return explored, nil
}

var errGoalNoFound = errors.New("Goal not found")

func (arr Frontier) Contains(ID string) bool {
	for _, v := range arr {
		if v.id == ID {
			return true
		}
	}
	return false
}

func AustraliaGraph() *Graph {
	drw1 := &Edge{id: "CNS", weight: 30}
	drw2 := &Edge{id: "ASP", weight: 15}
	drw3 := &Edge{id: "PER", weight: 48}
	drw := &Vertex{id: "1", value: "DRW", edges: []*Edge{drw1, drw2, drw3}}

	cns1 := &Edge{id: "DRW", weight: 30}
	cns2 := &Edge{id: "ASP", weight: 24}
	cns3 := &Edge{id: "BNE", weight: 22}
	cns := &Vertex{id: "2", value: "CNS", edges: []*Edge{cns1, cns2, cns3}}

	asp1 := &Edge{id: "DRW", weight: 15}
	asp2 := &Edge{id: "CNS", weight: 24}
	asp3 := &Edge{id: "BNE", weight: 31}
	asp4 := &Edge{id: "CBR", weight: 15}
	asp5 := &Edge{id: "ADP", weight: 15}
	asp := &Vertex{id: "3", value: "ASP", edges: []*Edge{asp1, asp2, asp3, asp4, asp5}}

	bne1 := &Edge{id: "CNS", weight: 22}
	bne2 := &Edge{id: "ASP", weight: 31}
	bne3 := &Edge{id: "SYD", weight: 9}
	bne := &Vertex{id: "4", value: "BNE", edges: []*Edge{bne1, bne2, bne3}}

	syd1 := &Edge{id: "BNE", weight: 9}
	syd2 := &Edge{id: "MEL", weight: 12}
	syd3 := &Edge{id: "CBR", weight: 4}
	syd := &Vertex{id: "5", value: "SYD", edges: []*Edge{syd1, syd2, syd3}}

	cbr1 := &Edge{id: "MEL", weight: 6}
	cbr2 := &Edge{id: "SYD", weight: 4}
	cbr3 := &Edge{id: "ASP", weight: 15}
	cbr := &Vertex{id: "6", value: "CBR", edges: []*Edge{cbr1, cbr2, cbr3}}

	mel1 := &Edge{id: "SYD", weight: 12}
	mel2 := &Edge{id: "CBR", weight: 6}
	mel3 := &Edge{id: "ASL", weight: 8}
	mel := &Vertex{id: "7", value: "MEL", edges: []*Edge{mel1, mel2, mel3}}

	adl1 := &Edge{id: "MEL", weight: 8}
	adl2 := &Edge{id: "ASP", weight: 15}
	adl3 := &Edge{id: "PER", weight: 32}
	adl := &Vertex{id: "8", value: "ADL", edges: []*Edge{adl1, adl2, adl3}}

	per1 := &Edge{id: "ADL", weight: 32}
	per2 := &Edge{id: "DRW", weight: 48}
	per := &Vertex{id: "9", value: "PER", edges: []*Edge{per1, per2}}
	vertices := make(map[string]*Vertex)
	vertices[drw.id] = drw
	vertices[cns.id] = cns
	vertices[asp.id] = asp
	vertices[bne.id] = bne
	vertices[syd.id] = syd
	vertices[cbr.id] = cbr
	vertices[mel.id] = mel
	vertices[adl.id] = adl
	vertices[per.id] = per
	g := &Graph{Vertices: vertices}
	return g
}

type Vertex struct {
	id    string
	value string
	edges []*Edge
}

func (v *Vertex) Edges() []*Edge {
	return v.edges
}

type Edge struct {
	id     string
	weight int
}

type Graph struct {
	Vertices map[string]*Vertex
}

func (g *Graph) Find(ID string) *Vertex {
	return g.Vertices[ID]
}
