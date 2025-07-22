package conditional

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// NotIdentifier is a constant representing the operation code for the logical "not" operation.
const NotIdentifier = 27

// OpNot represents a logical "not" operation in the instruction set, negating a boolean value on the execution stack.
type OpNot struct{}

// OpCode returns the operation code associated with the logical "not" operation.
func (o OpNot) OpCode() opcode.OpCode {
	return NotIdentifier
}

// Exec executes the logical "not" operation by negating the boolean value at the top of the stack in the given execution context.
// Returns the updated program counter and an error if the stack is empty or the top value is not boolean.
func (o OpNot) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	val := ctx.Stack.Pop()
	if val == nil {
		return ctx.PC, fmt.Errorf("not: stack underflow")
	}

	valBool, ok := val.(types.SymboliaBool)

	if !ok {
		return ctx.PC, fmt.Errorf("not: operand is not boolean")
	}

	ctx.Stack.Push(!valBool)

	return ctx.PC + 1, nil
}

// Name returns the name of the logical "not" operation as a string.
func (o OpNot) Name() types.SymboliaString {
	return "not"
}
