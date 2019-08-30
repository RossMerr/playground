package ucs

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
				n := g.Find(e.id)
				frontier = append(frontier, &path{append(p.Vertices, n), p.Cost + e.weight})
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

func (g *Graph) Find(ID string) *Vertex {
	return g.Vertices[ID]
}

type Edges []*Edge

func (a Edges) Len() int           { return len(a) }
func (a Edges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Edges) Less(i, j int) bool { return a[i].weight > a[j].weight }
