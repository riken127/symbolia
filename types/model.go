package types

import "github.com/riken127/symbolia/stack"

// Instruction represents a single operation to be executed, containing an OpCode, arguments, and an optional PipeTo value.
type Instruction struct {
	OpCode int   `json:"op_code"`
	Args   []any `json:"args"`
	PipeTo any   `json:"pipe_to,omitempty"`
}

// AbstractFunction represents a function definition with an identifier, symbol definitions, and a sequence of instructions.
type AbstractFunction struct {
	ID      int           `json:"id"`
	Symbols []string      `json:"symbols"`
	Stack   []Instruction `json:"stack"`
}

// Invocation represents a function invocation with a target function ID, arguments, and optional piping to another function.
type Invocation struct {
	FunctionID int   `json:"function_id"`
	Args       []any `json:"args"`
	PipeTo     any   `json:"pipe_to,omitempty"`
}

// ExecutionContext represents the state of execution, containing a stack, symbol table, and program counter.
type ExecutionContext struct {
	Stack   *stack.Stack
	Symbols map[string]any
	PC      int
}
