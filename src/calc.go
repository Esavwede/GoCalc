/*
@Authors: [ Ogagaoghene Esavwede ]
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

		// Determine UserOperation
		switch userOperation {

		case "A", "a":

			// Addition
			fmt.Println(" Addition Selected ")
			addition()

		case "M", "m":

			// Multiplication
			fmt.Println(" Multiplication Selected ")
			multiplication()

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

func addition() {

	// Message
	fmt.Println(" Only 2 numbers can be added currently ")

	// Addition Operation
	fmt.Println(" Enter the two integers you want to add ")

	// variables
	var integer1 int
	var integer2 int

	// Input Checks

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

func multiplication() {

	fmt.Println(" Enter two numbers you want to multiply ")

	// Variables
	var num1 int
	var num2 int
	var product int

	fmt.Println(" num1 ")
	fmt.Scanln(&num1)

	fmt.Println(" num2 ")
	fmt.Scanln(&num2)

	product = num1 * num2

	fmt.Printf(" The product of %v and %v is %v ", num1, num2, product)

}
