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

func UniformCostSearch(g Graph, start *Vertex, goal *Vertex) (map[string]*Vertex, error) {

	frontier := Frontier{}
	frontier = append(frontier, start)
	visited := make(map[string]*Vertex)

	for {
		length := len(frontier)
		if length == 0 {
			return nil, errGoalNoFound
		}

		node := frontier[length]
		if node == goal {
			break
		}

		visited[node.id] = node
		for _, e := range node.Edges() {
			n := g.Find(e.id)
			if _, ok := visited[n.id]; !ok {
				if !frontier.Contains(n.id) {
					frontier = append(frontier, n)
				}
			}
		}

	}

	return visited, nil
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

//
type Vertex struct {
	id string
}

func (v *Vertex) Edges() []*Edge {
	return nil
}

type Edge struct {
	id     string
	weight int
}

type Graph struct {
	Vertices []*Vertex
}

func (g *Graph) Find(ID string) *Vertex {
	return nil
}
