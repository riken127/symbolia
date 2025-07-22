package opcode

import (
	"fmt"
	"github.com/riken127/symbolia/types"
)

const ModIdentifier = 24

type OpMod struct{}

func (o OpMod) OpCode() OpCode {
	return ModIdentifier
}

func (o OpMod) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	left := ctx.Stack.Pop()
	right := ctx.Stack.Pop()
	if left == nil || right == nil {
		return ctx.PC, fmt.Errorf("mod: stack underflow")
	}

	leftFloat, ok1 := left.(int)
	rightFloat, ok2 := right.(int)
	if !ok1 || !ok2 {
		return ctx.PC, fmt.Errorf("mod: operands must be numeric values (got %T, %T)", left, right)
	}

	ctx.Stack.Push(((rightFloat % leftFloat) + leftFloat) % leftFloat)

	return ctx.PC + 1, nil
}

func (o OpMod) Name() string {
	return "mod"
}
