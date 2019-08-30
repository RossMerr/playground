package ucs

import "testing"
import "strconv"

func Test_UniformCostSearch(t *testing.T) {
	g := AustraliaGraph()
	start := g.Find("SYD")
	goal := g.Find("PER")
	result, err := UniformCostSearch(g, start, goal)

	if err != nil {
		t.Fatalf("Expected err to be nil but was %s", err)
	}

	count := len(result)
	if count != 5 {
		t.Fatalf("Expected count to be %s but was %s", "5", strconv.Itoa(count))
	}

	if result[0].id != "SYD" {
		t.Fatalf("Expected count to be %s but was %s", "SYD", result[0].id)
	}

	if result[1].id != "CBR" {
		t.Fatalf("Expected count to be %s but was %s", "CBR", result[1].id)
	}

	if result[2].id != "MEL" {
		t.Fatalf("Expected count to be %s but was %s", "MEL", result[2].id)
	}

	if result[3].id != "ADL" {
		t.Fatalf("Expected count to be %s but was %s", "ADL", result[3].id)
	}

	if result[4].id != "PER" {
		t.Fatalf("Expected count to be %s but was %s", "PER", result[4].id)
	}
}

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
