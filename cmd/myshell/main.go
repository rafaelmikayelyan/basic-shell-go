package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	builtins := [3]string{"exit", "echo", "type"}

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
			case inputs[0] == "type":
				if len(inputs) == 2 {
					isFound := false
					for _, builtin := range builtins {
						if builtin == inputs[1] {
							isFound = true
							break
						}
					}
					if isFound {
							fmt.Fprint(os.Stdout, inputs[1] + " is a shell builtin\n")
					} else {
							fmt.Fprint(os.Stdout, inputs[1] + " not found\n")
					}
				} else {
					fmt.Fprint(os.Stdout, "usage: type [command]\n")
				}
			default:
				fmt.Fprint(os.Stdout, input + ": command not found\n")
		}
	}
}
