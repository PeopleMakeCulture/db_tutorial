package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/** we need to: 
* print prompt 
* get input 
* process input
* repeat
*/
func main(){

	// create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// print user instructions
	fmt.Println("Enter your command. Enter q to quit.")

	// keep looping until user exits
	for {

		// prompt user for input
		fmt.Print("db > ")

		// read the input
		scanner.Scan()

		// retrieve the input value
		input := scanner.Text()
		input = strings.ToLower(input)

		// exit conditions
		exitCommands := []string{"q", "quit", "exit"}

		found := false
		for _ , phrase := range exitCommands {
			if input == phrase {
				found = true
				break
			}
		}
		if found {
			fmt.Println("goodbye!")
			break 
		}

		// we will learn to parse commands in part 2!
		fmt.Println("Unrecognized command: ", input)

	}

}