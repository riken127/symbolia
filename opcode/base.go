package opcode

import (
	"github.com/riken127/symbolia/types"
)

// OpCode represents the type for operation codes used in interpreting or executing instructions.
type OpCode int

// InstructionHandler is an interface for handling instructions in a virtual machine, associating execution logic and metadata.
type InstructionHandler interface {
	OpCode() OpCode
	Name() types.SymboliaString
	Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error)
}
