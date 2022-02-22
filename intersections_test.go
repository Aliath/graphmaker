package graphmaker

import (
	"math"
	"testing"
)

const epsilon float64 = 1e-5

func areNodesEqual(a Node, b Node) bool {
	distance := math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))

	return distance < epsilon
}

func TestIntersectionOfTwoNodes(t *testing.T) {
	p1 := Node{x: 0, y: 0}
	var r1 float64 = 3
	p2 := Node{x: 6, y: 0}
	var r2 float64 = 3

	expected := Node{x: 3, y: 0}
	result := *(GetNodesOfTwoIntersections(p1, r1, p2, r2))[0]

	if !areNodesEqual(expected, result) {
		t.Fatalf(`expected (%f, %f) but got (%f, %f)`, expected.x, expected.y, result.x, result.y)
	}
}

func TestIntersectionOfThreeNodes(t *testing.T) {
	p1 := Node{x: 0, y: 0}
	var r1 float64 = 2
	p2 := Node{x: 2, y: 0}
	var r2 float64 = 2 * math.Sqrt(2)
	p3 := Node{x: 2, y: 2}
	var r3 float64 = 2

	expected := Node{x: 0, y: 2}
	result := *(GetNodeOfThreeIntersections(p1, r1, p2, r2, p3, r3))

	if !areNodesEqual(expected, result) {
		t.Fatalf(`expected (%f, %f) but got (%f, %f)`, expected.x, expected.y, result.x, result.y)
	}
}
