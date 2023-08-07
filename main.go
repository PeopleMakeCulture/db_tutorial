package main

import (
	"bufio"
	"fmt"
	"os"
)

/** we need to: 
* print prompt 
* get input 
* process input
* repeat
*/
func main(){

	// print user instructions
	fmt.Println("Enter your command. Enter q to quit.")

	// create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// keep looping until user exits
	for {

		// prompt user for input
		fmt.Print("db > ")

		// read the input
		scanner.Scan()

		// retrieve the input value
		var input string = scanner.Text()

		// exit condition
		if input == "q" {
			fmt.Println("goodbye!")
			break 
		}

		// we will learn to parse commands in part 2!
		fmt.Println("Unrecognized command: ", input)

	}

}