package opcode

import (
	"fmt"

	"github.com/riken127/symbolia/types"
)

// StoreIdentifier is a constant representing the opcode for the store operation in the virtual machine.
const StoreIdentifier OpCode = 11

// OpStore represents an instruction for storing a value from the stack into the symbol table in a virtual machine.
type OpStore struct{}

// OpCode returns the operation code identifying the store operation in the virtual machine.
func (o OpStore) OpCode() OpCode { return StoreIdentifier }

// Exec executes the store operation, storing a value from the stack into the symbol table with the provided name.
// Returns the next program counter value or an error if the operation fails.
func (o OpStore) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	if len(instr.Args) != 1 {
		return ctx.PC, fmt.Errorf("store: expected 1 argument (symbol name)")
	}
	symName, ok := instr.Args[0].(string)
	if !ok {
		return ctx.PC, fmt.Errorf("store: argument must be string symbol name, got: %d", instr.Args[0])
	}

	val := ctx.Stack.Pop()
	if val == nil {
		return ctx.PC, fmt.Errorf("store: stack underflow")
	}

	ctx.Symbols[symName] = val

	return ctx.PC + 1, nil
}

// Name returns the name of the operation as a string.
func (o OpStore) Name() string {
	return "store"
}
