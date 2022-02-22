package graphmaker

import (
	"errors"
	"fmt"
)

type measurementMap map[string]map[string]Edge

type Node struct {
	x          float64
	y          float64
	identifier string
}
type Polygon []Node

type Edge struct {
	source string
	target string
	weight float64
}

func validateEdges(nodeIdentifiers []string, measurementMap measurementMap) (err error) {
	if len(nodeIdentifiers) < 3 {
		return errors.New("graph should have at least 3 nodes")
	}

	for i := 0; i < len(nodeIdentifiers); i++ {
		for j := i + 1; j < len(nodeIdentifiers); j++ {
			if _, err := GetDistanceBetween(nodeIdentifiers[i], nodeIdentifiers[j], measurementMap); err != nil {
				return err
			}
		}
	}

	return nil
}

func BuildEdgeMap(edges []Edge) measurementMap {
	edgesBySource := make(measurementMap)

	for _, measurement := range edges {
		if _, ok := edgesBySource[measurement.source]; !ok {
			edgesBySource[measurement.source] = make(map[string]Edge)
		}

		edgesBySource[measurement.source][measurement.target] = measurement
	}

	return edgesBySource
}

func GetDistanceBetween(source string, target string, measurementMap measurementMap) (result float64, err error) {
	if _, ok := measurementMap[source]; ok {
		if value, ok := measurementMap[source][target]; ok {
			return value.weight, nil
		}
	}

	if _, ok := measurementMap[target]; ok {
		if value, ok := measurementMap[target][source]; ok {
			return value.weight, nil
		}
	}

	return 0, fmt.Errorf("no readings between %s and %s", source, target)
}

func placeStartPoints(result *Polygon, measurementMap measurementMap, nodeIdentifiers []string) (err error) {
	(*result)[0] = Node{
		identifier: nodeIdentifiers[0],
		x:          0,
		y:          0,
	}

	distanceBetweenStartPoints, err := GetDistanceBetween(nodeIdentifiers[0], nodeIdentifiers[1], measurementMap)

	if err != nil {
		return err
	}

	// place Node on (x, 0)
	(*result)[1] = Node{
		identifier: nodeIdentifiers[1],
		x:          distanceBetweenStartPoints,
		y:          0,
	}

	return nil
}

func placeThirdPoint(result *Polygon, measurementMap measurementMap, nodeIdentifiers []string) (err error) {
	distanceOfFirstPoint, err := GetDistanceBetween(nodeIdentifiers[2], nodeIdentifiers[0], measurementMap)

	if err != nil {
		return err
	}

	distanceOfSecondPoint, err := GetDistanceBetween(nodeIdentifiers[1], nodeIdentifiers[0], measurementMap)

	if err != nil {
		return err
	}

	intersectionOfTwoPoints := GetNodesOfTwoIntersections((*result)[0], distanceOfFirstPoint, (*result)[1], distanceOfSecondPoint)[1]
	intersectionOfTwoPoints.identifier = nodeIdentifiers[2]

	(*result)[2] = *intersectionOfTwoPoints
	return nil
}

func placeRestPoints(result *Polygon, measurementMap measurementMap, nodeIdentifiers []string) (err error) {
	for i := 3; i < len(nodeIdentifiers); i++ {
		p1 := (*result)[i-3]
		p2 := (*result)[i-2]
		p3 := (*result)[i-1]

		r1, err := GetDistanceBetween(nodeIdentifiers[i], p1.identifier, measurementMap)
		if err != nil {
			return err
		}

		r2, err := GetDistanceBetween(nodeIdentifiers[i], p2.identifier, measurementMap)
		if err != nil {
			return err
		}

		r3, err := GetDistanceBetween(nodeIdentifiers[i], p3.identifier, measurementMap)
		if err != nil {
			return err
		}

		(*result)[i] = *GetNodeOfThreeIntersections(p1, r1, p2, r2, p3, r3)
		(*result)[i].identifier = nodeIdentifiers[i]
	}

	return nil
}

func BuildGraph(edges []Edge, nodeIdentifiers []string) (result Polygon, err error) {
	measurementMap := BuildEdgeMap(edges)

	if err = validateEdges(nodeIdentifiers, measurementMap); err != nil {
		return result, err
	}

	result = make(Polygon, len(nodeIdentifiers))

	// place two first Nodes on (0, 0) and (x, 0)
	if err = placeStartPoints(&result, measurementMap, nodeIdentifiers); err != nil {
		return result, err
	}

	// place third point on specific intersection of two circles
	if err = placeThirdPoint(&result, measurementMap, nodeIdentifiers); err != nil {
		return result, err
	}

	// place rest of points by intersectino of three circles
	if err = placeRestPoints(&result, measurementMap, nodeIdentifiers); err != nil {
		return result, err
	}

	return result, nil
}
