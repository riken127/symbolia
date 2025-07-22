package misc

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// TypeIdentifier represents the opcode for the `type` operation
const TypeIdentifier opcode.OpCode = 31

// OpType represents the operation to get the type of the top element on the stack
type OpType struct{}

// OpCode returns the unique opcode for the `type` operation
func (o OpType) OpCode() opcode.OpCode {
	return TypeIdentifier
}

// Name returns the human-readable name of the operation
func (o OpType) Name() types.SymboliaString {
	return "type"
}

// Exec retrieves the type of the top element of the stack (e.g., "number", "string") and pushes the result as a string
func (o OpType) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	// Peek at the top element of the stack
	val := ctx.Stack.Pop()

	// Handle empty stack (stack underflow)
	if val == nil {
		return ctx.PC, fmt.Errorf("type: stack underflow, no element to inspect")
	}

	// Determine the type of the value
	var typeName string
	switch val.(type) {
	case types.SymboliaNumber:
		typeName = "s_number"
	case types.SymboliaString:
		typeName = "s_string"
	case types.SymboliaBool:
		typeName = "s_bool"
	case types.SymboliaInteger:
		typeName = "s_integer"
	default:
		typeName = "s_unknown"
	}

	// Push the type name as a string onto the stack
	ctx.Stack.Push(typeName)

	return ctx.PC + 1, nil
}
