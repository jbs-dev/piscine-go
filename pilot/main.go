package main

import "fmt"

// Define a constant called AIRCRAFT1 and set it equal to 1.
const AIRCRAFT1 = 1

// Define a struct called Pilot with four fields: Name, Life, Age, and Aircraft.
type Pilot struct {
	Name     string
	Life     int
	Age      int
	Aircraft int
}

func main() {
	// Declare a variable called donnie of type Pilot.
	var donnie Pilot
	// Set the value of the Name field of the donnie struct to "Donnie".
	donnie.Name = "Donnie"
	// Set the value of the Life field of the donnie struct to 100.0.
	donnie.Life = 100.0
	// Set the value of the Age field of the donnie struct to 24.
	donnie.Age = 24
	// Set the value of the Aircraft field of the donnie struct to AIRCRAFT1.
	donnie.Aircraft = AIRCRAFT1

	// Print the value of the donnie variable to the console.
	fmt.Println(donnie)
}
