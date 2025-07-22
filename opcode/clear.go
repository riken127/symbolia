package opcode

import (
	"fmt"
	"github.com/riken127/symbolia/types"
)

const ClearIdentifier OpCode = 15

type OpClear struct{}

func (o OpClear) OpCode() OpCode {
	return ClearIdentifier
}

func (o OpClear) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	if len(instr.Args) == 0 || instr.Args[0] == "stack" {
		ctx.Stack.Clear()
	} else if instr.Args[0] == "symbols" {
		ctx.Symbols = make(map[string]any)
	} else if instr.Args[0] == "all" {
		ctx.Stack.Clear()
		ctx.Symbols = make(map[string]any)
	} else {
		return ctx.PC, fmt.Errorf("clear: invalid argument: %v", instr.Args[0])
	}

	return ctx.PC + 1, nil
}

func (o OpClear) Name() string {
	return "clear"
}
