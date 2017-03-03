package playground

import (
	"errors"
	"sort"
)

// UniformCostSearch (Graph, start, goal)
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
func UniformCostSearch(g *Graph, start *Vertex, goal *Vertex) ([]*Vertex, error) {
	frontier := frontier{&path{[]*Vertex{start}, 0}}
	explored := make(map[string]bool)

	for {
		if len(frontier) == 0 {
			return nil, errGoalNoFound
		}

		var p *path
		sort.Sort(frontier)
		p, frontier = frontier.pop()
		node := p.Vertices[len(p.Vertices)-1]
		explored[node.id] = true

		if node == goal {
			return p.Vertices, nil
		}

		for _, e := range node.Edges() {
			if _, ok := explored[e.id]; !ok {
				if n, ok := g.Find(e.id); ok {
					frontier = append(frontier, &path{append(p.Vertices, n), p.Cost + e.weight})
				} else {
					return nil, errors.New("Vertex not found " + e.id)
				}

			}
		}
	}
}

var errGoalNoFound = errors.New("Goal not found")

type path struct {
	Vertices []*Vertex
	Cost     int
}

type frontier []*path

func (f frontier) Len() int               { return len(f) }
func (f frontier) Swap(i, j int)          { f[i], f[j] = f[j], f[i] }
func (f frontier) Less(i, j int) bool     { return f[i].Cost < f[j].Cost }
func (f frontier) pop() (*path, frontier) { return f[0], f[1:] }

func AustraliaGraph() *Graph {
	drw1 := &Edge{id: "CNS", weight: 30}
	drw2 := &Edge{id: "ASP", weight: 15}
	drw3 := &Edge{id: "PER", weight: 48}
	drw := &Vertex{id: "DRW", edges: Edges{drw1, drw2, drw3}}

	cns1 := &Edge{id: "DRW", weight: 30}
	cns2 := &Edge{id: "ASP", weight: 24}
	cns3 := &Edge{id: "BNE", weight: 22}
	cns := &Vertex{id: "CNS", edges: Edges{cns1, cns2, cns3}}

	asp1 := &Edge{id: "DRW", weight: 15}
	asp2 := &Edge{id: "CNS", weight: 24}
	asp3 := &Edge{id: "BNE", weight: 31}
	asp4 := &Edge{id: "CBR", weight: 15}
	asp5 := &Edge{id: "ADL", weight: 15}
	asp := &Vertex{id: "ASP", edges: Edges{asp1, asp2, asp3, asp4, asp5}}

	bne1 := &Edge{id: "CNS", weight: 22}
	bne2 := &Edge{id: "ASP", weight: 31}
	bne3 := &Edge{id: "SYD", weight: 9}
	bne := &Vertex{id: "BNE", edges: Edges{bne1, bne2, bne3}}

	syd1 := &Edge{id: "BNE", weight: 9}
	syd2 := &Edge{id: "MEL", weight: 12}
	syd3 := &Edge{id: "CBR", weight: 4}
	syd := &Vertex{id: "SYD", edges: Edges{syd1, syd2, syd3}}

	cbr1 := &Edge{id: "MEL", weight: 6}
	cbr2 := &Edge{id: "SYD", weight: 4}
	cbr3 := &Edge{id: "ASP", weight: 15}
	cbr := &Vertex{id: "CBR", edges: Edges{cbr1, cbr2, cbr3}}

	mel1 := &Edge{id: "SYD", weight: 12}
	mel2 := &Edge{id: "CBR", weight: 6}
	mel3 := &Edge{id: "ADL", weight: 8}
	mel := &Vertex{id: "MEL", edges: Edges{mel1, mel2, mel3}}

	adl1 := &Edge{id: "MEL", weight: 8}
	adl2 := &Edge{id: "ASP", weight: 15}
	adl3 := &Edge{id: "PER", weight: 32}
	adl := &Vertex{id: "ADL", edges: Edges{adl1, adl2, adl3}}

	per1 := &Edge{id: "ADL", weight: 32}
	per2 := &Edge{id: "DRW", weight: 48}
	per := &Vertex{id: "PER", edges: Edges{per1, per2}}
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
	edges Edges
}

func (v *Vertex) Edges() Edges {
	return v.edges
}

type Edge struct {
	id     string
	weight int
}

type Graph struct {
	Vertices map[string]*Vertex
}

func (g *Graph) Find(ID string) (*Vertex, bool) {
	v, b := g.Vertices[ID]
	return v, b
}

type Edges []*Edge

func (a Edges) Len() int           { return len(a) }
func (a Edges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Edges) Less(i, j int) bool { return a[i].weight > a[j].weight }
