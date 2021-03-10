package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) int {
	fmt.Println("The amount of fuel left is:", fuel)
	return fuel
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int
	if planet == "Venus" {
		fuel = 300000
	} else if planet == "Mercury" {
		fuel = 500000
	} else if planet == "Mars" {
		fuel = 700000
	} else {
		fuel = 0
	}
	return fuel
}

// Create the function greetPlanet() here

func greetPlanet(planet string) {
	fmt.Printf("You have arrived at %v \n", planet)
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the avaialable fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {

	// Create `planetChoice` and `fuel`
	var fuel int = 1000000
	var planetChoice string = "Venus"
	// And then liftoff!
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
