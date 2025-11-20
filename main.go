/**
 * @author Daren Ileleji
 * @versopn 1.0.0
 * @date 2025-11-20
 * @fileoverview This program will display and calculate area of a circle with a radius of 5.6 cm
 */

package main

import "fmt"

func main () {
	// INPUT (Declaring Constants)
	const RADIUS float64 = 5.6
	const PI float64 = 3.14

	// PROCESSING (Calculating Area)
	answer := RADIUS * RADIUS * PI


	// OUTPUT (Showcasing Results)
	fmt.Println("The area of a circle with a radius of ", RADIUS, "cm is", answer, "cm2")
	fmt.Println("\nDone.")


}

