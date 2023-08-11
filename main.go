package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func do_meta_command(input string) error {

	switch input {
	case ".exit":
		fmt.Println("goodbye!")
		os.Exit(0)
	default:
		return fmt.Errorf("unrecognized command: %s", input)
	}
	// we did stuff successfully
	return nil
}

func execute_command(input string) error {
	switch input {
	case "select":
		fmt.Println("SELECT")
	case "insert":
		fmt.Println("INSERT")
	default:
		return fmt.Errorf("unrecognized command: %s", input)
	}
	return nil
}

/** we need to:
* print prompt
* get input
* process input
* repeat
 */
func main() {

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

		// if input starts with .
		if input[0] == '.' {
			err := do_meta_command(input)

			if err != nil {
				fmt.Println(err)
				continue
			}

		}

		// do regular command
		err := execute_command(input)

		if err != nil {
			fmt.Println(err)
			continue
		}

	}

}
