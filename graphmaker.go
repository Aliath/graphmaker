package graphmaker

import (
	"errors"
	"fmt"
)

type measurementMap map[string]map[string]Measurement

type node struct {
	x          float64
	y          float64
	identifier string
}
type Polygon []node

type Measurement struct {
	source   string
	target   string
	distance float64
}

func validateMeasurements(pointIdentifiers []string, measurementMap measurementMap) (err error) {
	if len(pointIdentifiers) < 3 {
		return errors.New("graph should have at least 3 nodes")
	}

	for i := 0; i < len(pointIdentifiers); i++ {
		for j := i + 1; j < len(pointIdentifiers); j++ {
			if _, err := getDistanceBetween(pointIdentifiers[i], pointIdentifiers[j], measurementMap); err != nil {
				return err
			}
		}
	}

	return nil
}

func buildmeasurementMap(measurements []Measurement) measurementMap {
	measurementsBySource := make(measurementMap)

	for _, measurement := range measurements {
		if _, ok := measurementsBySource[measurement.source]; !ok {
			measurementsBySource[measurement.source] = make(map[string]Measurement)
		}

		measurementsBySource[measurement.source][measurement.target] = measurement
	}

	return measurementsBySource
}

func getDistanceBetween(source string, target string, measurementMap measurementMap) (result float64, err error) {
	if _, ok := measurementMap[source]; ok {
		if value, ok := measurementMap[source][target]; ok {
			return value.distance, nil
		}
	}

	if _, ok := measurementMap[target]; ok {
		if value, ok := measurementMap[target][source]; ok {
			return value.distance, nil
		}
	}

	return 0, fmt.Errorf("no readings between %s and %s", source, target)
}

func placeStartPoints(result *Polygon, measurementMap measurementMap, pointIdentifiers []string) (err error) {
	(*result)[0] = node{
		identifier: pointIdentifiers[0],
		x:          0,
		y:          0,
	}

	distanceBetweenStartPoints, err := getDistanceBetween(pointIdentifiers[0], pointIdentifiers[1], measurementMap)

	if err != nil {
		return err
	}

	// place node on (x, 0)
	(*result)[1] = node{
		identifier: pointIdentifiers[1],
		x:          distanceBetweenStartPoints,
		y:          0,
	}

	return nil
}

func placeThirdPoint(result *Polygon, measurementMap measurementMap, pointIdentifiers []string) (err error) {
	distanceOfFirstPoint, err := getDistanceBetween(pointIdentifiers[2], pointIdentifiers[0], measurementMap)

	if err != nil {
		return err
	}

	distanceOfSecondPoint, err := getDistanceBetween(pointIdentifiers[1], pointIdentifiers[0], measurementMap)

	if err != nil {
		return err
	}

	intersectionOfTwoPoints := GetNodesOfTwoIntersections((*result)[0], distanceOfFirstPoint, (*result)[1], distanceOfSecondPoint)[0]
	intersectionOfTwoPoints.identifier = pointIdentifiers[2]

	(*result)[2] = *intersectionOfTwoPoints
	return nil
}

func placeRestPoints(result *Polygon, measurementMap measurementMap, pointIdentifiers []string) (err error) {
	for i := 3; i < len(pointIdentifiers); i++ {
		p1 := (*result)[i-3]
		p2 := (*result)[i-2]
		p3 := (*result)[i-1]

		r1, err := getDistanceBetween(pointIdentifiers[i], p1.identifier, measurementMap)
		if err != nil {
			return err
		}

		r2, err := getDistanceBetween(pointIdentifiers[i], p2.identifier, measurementMap)
		if err != nil {
			return err
		}

		r3, err := getDistanceBetween(pointIdentifiers[i], p3.identifier, measurementMap)
		if err != nil {
			return err
		}

		(*result)[i] = *GetPointsOfThreeIntersections(p1, r1, p2, r2, p3, r3)
		(*result)[i].identifier = pointIdentifiers[i]
	}

	return nil
}

func BuildGraph(measurements []Measurement, pointIdentifiers []string) (result Polygon, err error) {
	measurementMap := buildmeasurementMap(measurements)

	if err = validateMeasurements(pointIdentifiers, measurementMap); err != nil {
		return result, err
	}

	result = make(Polygon, len(pointIdentifiers))

	// place two first nodes on (0, 0) and (x, 0)
	if err = placeStartPoints(&result, measurementMap, pointIdentifiers); err != nil {
		return result, err
	}

	// place third point on specific intersection of two circles
	if err = placeThirdPoint(&result, measurementMap, pointIdentifiers); err != nil {
		return result, err
	}

	// place rest of points by intersectino of three circles
	if err = placeRestPoints(&result, measurementMap, pointIdentifiers); err != nil {
		return result, err
	}

	return result, nil
}
