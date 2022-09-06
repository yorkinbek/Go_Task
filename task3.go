package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64 = 10.04

	// fmt.Scanf("%f", &r)
	fmt.Println("R =", r)

	fmt.Printf("Area: %0.2f\n", area(r))
}

func area(r float64) (area float64) {
	//
	// WRITE YOUR CODE HERE
	//
	const Pi = math.Pi
	var areaOfOneSquere = 4 * (math.Pow(r, 2))
	var areaOfCircle = r * r * Pi
	area = areaOfOneSquere - areaOfCircle
	return area
}
