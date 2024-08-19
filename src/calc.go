/*
@Authors: Ogagaoghene Esavwede
@description: A basic calculator program written in the Go programming language. This calculator can only perform basic arithmetic operations.
@createdOn: 19/09/2024
*/

package main

import "fmt"

func main() {
	for {
		// Main Message
		fmt.Println(" Welcome to GoCalc. \n A basic calculator written in Go ")

		fmt.Println(" What operation do you want to perform ")
		fmt.Println(" Enter 'A' for addition \n 'M' for multiplication  \n 'S' for subtraction \n 'D' for division \n 'E' for exponentiation ")

		// Get User input
		var userOperation string
		fmt.Scanln(&userOperation)

		// Analyse Input
		switch userOperation {

		case "A", "a":

			// Addition
			fmt.Println(" Addition Selected ")
			additionFunction()

		case "M", "m":

			// Multiplication
			fmt.Println(" Multiplication Selected ")
			fmt.Println(" Multiplication Currently Unavailable ")

		case "S", "s":

			// Subtraction
			fmt.Println(" Subtraction Selected ")
			fmt.Println(" Subtraction Currently Unavailable ")

		case "D", "d":

			// Division
			fmt.Println(" Division Selected ")
			fmt.Println(" Division Currently Unavailable ")

		case "E", "e":

			// Exponentiation
			fmt.Println(" Exponentiation Selected ")
			fmt.Println(" Exponentiation Currently Unavailable ")

		default:

			// Unknown Operation
			fmt.Println(" Unknown Operation Selected ")
		}

		fmt.Println(" Press q to exit ")
		var exitState string

		fmt.Scanln(&exitState)

		if exitState == "q" || exitState == "Q" {
			// exit program
			break
		}
	}
}

func additionFunction() {

	// Message
	fmt.Println(" Only 2 numbers can be added currently ")

	// Addition Operation
	fmt.Println(" Enter the two integers you want to add ")

	// variables
	var integer1 int
	var integer2 int

	// Addition Input
	fmt.Println(" Integer 1: ")
	fmt.Scanln(&integer1)

	fmt.Println(" Integer 2: ")
	fmt.Scanln(&integer2)

	// Result
	var sum = integer1 + integer2

	// Result Output
	fmt.Printf(" The sum of %v and %v is %v \n ", integer1, integer2, sum)
}
