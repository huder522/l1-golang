package main

import "fmt"

func main() {
	var number int64;
	fmt.Println( "Input a number:" )
	fmt.Scanln( &number )

	if number == 0 {
		fmt.Println( "The number is 0" )
	} else if number % 2 == 0 {
		fmt.Println("The number is even")
	} else if number % 2 != 0 {
		fmt.Println("The number is odd")
	} else {
		fmt.Println("Undefined")
	}

}