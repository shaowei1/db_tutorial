package main

import (
	"bufio"
	"fmt"
	"os"
)

type InputBuffer struct {
	buffer       *string
	bufferLength int
	inputLength  int
}

func newInputBuffer() *InputBuffer {
	return &InputBuffer{
		buffer:       nil,
		bufferLength: 0,
		inputLength:  0,
	}
}

func printPrompt() {
	fmt.Print("db > ")
}

func readInput(inputBuffer *InputBuffer) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		os.Exit(1)
	}
	line = line[:len(line)-1]
	inputBuffer.buffer = &line
	inputBuffer.bufferLength = len(line)
	inputBuffer.inputLength = len(line) - 1 // Ignore trailing newline
}

func main() {
	inputBuffer := newInputBuffer()
	for {
		printPrompt()
		readInput(inputBuffer)

		if *inputBuffer.buffer == ".exit" {
			os.Exit(0)
		} else {
			fmt.Printf("Unrecognized command '%s'.\n", *inputBuffer.buffer)
		}
	}
}
