package graphmaker

import (
	"fmt"
	"math"
	"testing"
)

func ValidateGraph(edges []Edge, result Polygon) error {
	measurementMap := BuildEdgeMap(edges)

	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			source := result[i]
			target := result[j]

			distance, err := GetDistanceBetween(source.identifier, target.identifier, measurementMap)
			if err != nil {
				return err
			}

			polygonDistance := math.Sqrt(math.Pow(source.x-target.x, 2) + math.Pow(source.y-target.y, 2))
			delta := math.Abs(polygonDistance - distance)

			if delta > epsilon {
				return fmt.Errorf("delta for d(%s, %s) is equal %f but maximum is %f", source.identifier, target.identifier, delta, epsilon)
			}
		}
	}

	return nil
}

func TestBuildGraph(t *testing.T) {
	edges := []Edge{
		{"A", "B", 5},
		{"A", "C", 5 * math.Sqrt(2)},
		{"A", "D", 5},

		{"B", "C", 5},
		{"B", "D", 5 * math.Sqrt(2)},

		{"C", "D", 5},
	}
	nodeIdentifiers := []string{"A", "B", "C", "D"}

	result, err := BuildGraph(edges, nodeIdentifiers)
	fmt.Println(result)

	if err != nil {
		t.Errorf("%s", err)
	}

	err = ValidateGraph(edges, result)

	if err != nil {
		t.Errorf("%s", err)
	}
}
