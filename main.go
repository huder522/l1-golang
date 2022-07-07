package main

import "fmt"

/*
	FUNCTIONS
*/

// add function
func add() {

	var num1, num2 int

	fmt.Println( "This program will be to add two numbers:" )

	fmt.Println( "Input the number 1" )
	fmt.Scanln( &num1)
	fmt.Println( "Input the number 2" )
	fmt.Scanln( &num2 )

	fmt.Println( "The sum of the numbers is: ", num1 + num2 )

}

// division function
func division() {

	var num1, num2 float32

	fmt.Println( "This program will be to divides two numbers and returns the quotient and remainder:" )

	fmt.Println( "Input the number 1" )
	fmt.Scanln( &num1)
	fmt.Println( "Input the number 2" )
	fmt.Scanln( &num2 )

	var quotient float32 = num1 / num2
	var remainder  = int(num1) % int(num2)

	fmt.Println( "The quotient of these devision is: ", quotient )
	fmt.Println( "The remainder of these devision is: ", remainder )

}


func main() {
	
	add()

	division()

}