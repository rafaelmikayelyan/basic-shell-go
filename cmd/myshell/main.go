package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	builtins := []string{"exit", "echo", "type"}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = input[:len(input) - 1]
		inputs := strings.Fields(input)

		switch {
		case len(inputs) < 2:
			fmt.Fprint(os.Stdout, "usage: [command] [arg]\n")
		case input == "exit 0":
			os.Exit(0)
		case inputs[0] == "echo":
			fmt.Fprint(os.Stdout, input[5:] + "\n")
		case inputs[0] == "type":
			if !isFound(builtins, inputs) {
				fmt.Fprintf(os.Stdout, "%s not found\n", inputs[1])
			}
		default:
			fmt.Fprint(os.Stdout, input + ": command not found\n")
		}
	}
}

func isFound(builtins []string, inputs []string) bool {
	for _, builtin := range builtins {
		if builtin == inputs[1] {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", inputs[1])
			return true
		}
	}

	for _, path := range []string{"/bin", "/usr/bin", "/usr/local/bin"} {
		if isExternalFileFound(inputs[1], path) {
			return true
		}
	}
	return false
}

func isExternalFileFound(filename string, path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprint(os.Stdout, err)
	}

	for _, file := range files {
		if file.Name() == filename {
			fmt.Fprintf(os.Stdout, "%s is %s/%s\n", filename, path, filename)
			return true
		}
	}
	return false
}
