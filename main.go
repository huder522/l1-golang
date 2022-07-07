package main

import "fmt"

// functions
func add() {

	var num1, num2 int

	fmt.Println( "This program will be to add two numbers:" )

	fmt.Println( "Input the number 1" )
	fmt.Scanln( &num1)
	fmt.Println( "Input the number 2" )
	fmt.Scanln( &num2 )

	fmt.Println( "The sum of the numbers is: ", num1 + num2 )

}



func main() {
	
	add()

}