package opcode

import (
	"fmt"

	"github.com/riken127/symbolia/types"
)

// PrintIdentifier is a constant representing the opcode for the print operation in the virtual machine.
const PrintIdentifier = 6

// OpPrint represents the print operation in the virtual machine instruction set.
// It handles printing the top value from the execution context's stack.
type OpPrint struct{}

// OpCode returns the operation code associated with the print instruction. It identifies this operation as PrintIdentifier.
func (o OpPrint) OpCode() OpCode { return PrintIdentifier }

// Exec executes the print operation by popping the top value from the stack, printing it, and advancing the program counter.
func (o OpPrint) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	val := ctx.Stack.Pop()
	fmt.Print(val)
	return ctx.PC + 1, nil
}

// Name returns the name of the operation as a string identifier, which is "print".
func (o OpPrint) Name() string { return "print" }
