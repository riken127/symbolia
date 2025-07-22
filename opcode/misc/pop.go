package misc

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// PopIdentifier is the opcode representing the "pop" operation, used to remove the top value from the execution stack.
const PopIdentifier opcode.OpCode = 28

// OpPop represents an operation that removes the top value from the execution stack.
type OpPop struct{}

// OpCode returns the opcode associated with the "pop" operation.
func (o OpPop) OpCode() opcode.OpCode {
	return PopIdentifier
}

// Exec executes the "pop" operation by removing the top value from the stack. Returns the next program counter or an error.
func (o OpPop) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	val := ctx.Stack.Pop()

	if val == nil {
		return ctx.PC, fmt.Errorf("pop: stack underflow")
	}

	return ctx.PC + 1, nil
}

// Name returns the name of the operation, which is "pop".
func (o OpPop) Name() types.SymboliaString {
	return "pop"
}
