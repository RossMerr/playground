package playground

import "testing"
import "strconv"

func Test_UniformCostSearch(t *testing.T) {
	g := AustraliaGraph()
	start, _ := g.Find("SYD")
	goal, _ := g.Find("PER")
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
