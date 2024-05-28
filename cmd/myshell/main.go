package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = input[:len(input) - 1]
		inputs := strings.Fields(input)


		
		switch {
			case input == "exit 0":
				os.Exit(0)
			case inputs[0] == "echo":
				fmt.Fprint(os.Stdout, input[5:] + "\n")
			default:
				fmt.Fprint(os.Stdout, input + ": command not found\n")
		}
	}
}
