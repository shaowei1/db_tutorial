package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputBuffer struct {
	buffer       *string
	bufferLength int
	inputLength  int
}
type MetaCommandResult int

const (
	META_COMMAND_SUCCESS MetaCommandResult = iota
	META_COMMAND_UNRECOGNIZED_COMMAND
)

type PrepareResult int

const (
	PREPARE_SUCCESS PrepareResult = iota
	PREPARE_UNRECOGNIZED_STATEMENT
)

type StatementType int

const (
	STATEMENT_INSERT StatementType = iota
	STATEMENT_SELECT
)

type Statement struct {
	Type StatementType
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

func doMetaCommand(inputBuffer *InputBuffer) MetaCommandResult {
	if *inputBuffer.buffer == ".exit" {
		os.Exit(0)
		return META_COMMAND_SUCCESS
	} else {
		return META_COMMAND_UNRECOGNIZED_COMMAND
	}
}

func prepareStatement(inputBuffer *InputBuffer, statement *Statement) PrepareResult {
	if strings.HasPrefix(*inputBuffer.buffer, "insert") {
		statement.Type = STATEMENT_INSERT
		return PREPARE_SUCCESS
	}
	if *inputBuffer.buffer == "select" {
		statement.Type = STATEMENT_SELECT
		return PREPARE_SUCCESS
	}
	return PREPARE_UNRECOGNIZED_STATEMENT
}

func executeStatement(statement *Statement) {
	switch statement.Type {
	case STATEMENT_INSERT:
		fmt.Println("This is where we would do an insert.")
	case STATEMENT_SELECT:
		fmt.Println("This is where we would do a select.")
	}
}
func main() {
	inputBuffer := newInputBuffer()
	for {
		printPrompt()
		readInput(inputBuffer)

		if strings.HasPrefix(*inputBuffer.buffer, ".") {
			switch doMetaCommand(inputBuffer) {
			case META_COMMAND_SUCCESS:
				continue
			case META_COMMAND_UNRECOGNIZED_COMMAND:
				fmt.Printf("Unrecognized command '%s'\n", *inputBuffer.buffer)
				continue
			}
		}

		statement := &Statement{}
		switch prepareStatement(inputBuffer, statement) {
		case PREPARE_SUCCESS:
			break
		case PREPARE_UNRECOGNIZED_STATEMENT:
			fmt.Printf("Unrecognized keyword at start of '%s'\n", *inputBuffer.buffer)
			continue
		}
		executeStatement(statement)
		fmt.Println("Executed.")
	}
}
