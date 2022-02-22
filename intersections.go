package graphmaker

import "math"

func GetNodesOfTwoIntersections(p1 node, r1 float64, p2 node, r2 float64) (result [2]*node) {
	distance := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
	length := (math.Pow(r1, 2) - math.Pow(r2, 2) + math.Pow(distance, 2)) / (2 * distance)
	height := math.Sqrt(math.Pow(r1, 2) - math.Pow(length, 2))

	for index := 0; index < 2; index++ {
		sign := math.Pow(-1, float64(index))

		result[index] = &node{
			x: (length/distance)*(p2.x-p1.x) + sign*(height/distance)*(p2.y-p1.y) + p1.x,
			y: (length/distance)*(p2.y-p1.y) + sign*(height/distance)*(p2.x-p1.x) + p1.y,
		}
	}

	return result
}

func GetPointsOfThreeIntersections(
	p1 node,
	r1 float64,
	p2 node,
	r2 float64,
	p3 node,
	r3 float64,
) *node {
	y := ((p2.x-p3.x)*((math.Pow(p2.x, 2)-math.Pow(p1.x, 2))+(math.Pow(p2.y, 2)-math.Pow(p1.y, 2))+(math.Pow(r1, 2)-math.Pow(r2, 2))) - (p1.x-p2.x)*((math.Pow(p3.x, 2)-math.Pow(p2.x, 2))+(math.Pow(p3.y, 2)-math.Pow(p2.y, 2))+(math.Pow(r2, 2)-math.Pow(r3, 2)))) / (2 * ((p1.y-p2.y)*(p2.x-p3.x) - (p2.y-p3.y)*(p1.x-p2.x)))
	x := ((p2.y-p3.y)*((math.Pow(p2.y, 2)-math.Pow(p1.y, 2))+(math.Pow(p2.x, 2)-math.Pow(p1.x, 2))+(math.Pow(r1, 2)-math.Pow(r2, 2))) - (p1.y-p2.y)*((math.Pow(p3.y, 2)-math.Pow(p2.y, 2))+(math.Pow(p3.x, 2)-math.Pow(p2.x, 2))+(math.Pow(r2, 2)-math.Pow(r3, 2)))) / (2 * ((p1.x-p2.x)*(p2.y-p3.y) - (p2.x-p3.x)*(p1.y-p2.y)))

	return &node{
		x: x,
		y: y,
	}
}
