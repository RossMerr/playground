package playground

import "testing"

func Test_UniformCostSearch(t *testing.T) {
	g := AustraliaGraph()
	start := g.Find("SYD")
	goal := g.Find("PER")
	result, err := UniformCostSearch(g, start, goal)

	if err != nil {
		t.Fatalf("Expected err to be nil but was %s", err)
	}

	count := len(result)
	if count == 5 {
		t.Fatalf("Expected count to be %s but was %s", string(5), string(count))
	}

}
