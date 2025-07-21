package engine

import (
	"fmt"

	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/stack"
	"github.com/riken127/symbolia/types"
)

// Executor manages the execution of functions and opcodes within a given execution context.
type Executor struct {
	ctx         *types.ExecutionContext                     // ctx is the current execution context that manages the stack, symbol table, and program counter for opcode execution.
	functions   map[int]types.AbstractFunction              // functions maps function IDs to their corresponding abstract function definitions.
	opcodeTable map[opcode.OpCode]opcode.InstructionHandler // opcodeTable maps OpCode values to their corresponding InstructionHandler implementations for execution.
}

// NewExecutor initializes and returns an Executor with the provided functions and invocation context.
func NewExecutor(funcs []types.AbstractFunction, inv types.Invocation) *Executor {
	fnMap := make(map[int]types.AbstractFunction)
	for _, f := range funcs {
		fnMap[f.ID] = f
	}

	fn, ok := fnMap[inv.FunctionID]
	if !ok {
		panic(fmt.Sprintf("function with ID %d not found", inv.FunctionID))
	}

	symbolTable := make(map[string]any)
	for i, name := range fn.Symbols {
		if i < len(inv.Args) {
			symbolTable[name] = inv.Args[i]
		}
	}

	ctx := &types.ExecutionContext{
		Stack:   stack.NewStack(),
		Symbols: symbolTable,
		PC:      0,
	}

	return &Executor{
		ctx:         ctx,
		functions:   fnMap,
		opcodeTable: opcode.OpCodeRegistry,
	}
}

// ExecuteInvocation runs the specified function by its ID, processing its stack and symbols in the execution context.
// It resolves arguments, executes instructions using handlers from the opcode table, and updates the program counter.
// Returns an error if the function ID is invalid, a required opcode is missing, or execution encounters an issue.
func (e *Executor) ExecuteInvocation(inv types.Invocation) error {
	fn, ok := e.functions[inv.FunctionID]
	if !ok {
		return fmt.Errorf("function %d not found", inv.FunctionID)
	}

	e.ctx.PC = 0

	for i, sym := range fn.Symbols {
		if i < len(inv.Args) {
			e.ctx.Symbols[sym] = inv.Args[i]
		}
	}

	for e.ctx.PC < len(fn.Stack) {
		instr := fn.Stack[e.ctx.PC]

		handler, ok := e.opcodeTable[opcode.OpCode(instr.OpCode)]
		if !ok {
			return fmt.Errorf("opcode %d not implemented", instr.OpCode)
		}

		resolvedArgs := e.resolveArgs(instr.Args, e.ctx.Symbols)
		resolvedInstr := types.Instruction{
			OpCode: instr.OpCode,
			Args:   resolvedArgs,
			PipeTo: instr.PipeTo,
		}

		nextPC, err := handler.Exec(e.ctx, resolvedInstr)
		if err != nil {
			return fmt.Errorf("error in %s: %w", handler.Name(), err)
		}

		e.ctx.PC = nextPC
	}

	return nil
}

// resolveArgs resolves a slice of arguments by replacing string keys with their corresponding values from the symbol table.
func (e *Executor) resolveArgs(args []any, symbolTable map[string]any) []any {
	resolved := make([]any, len(args))
	for i, arg := range args {
		if symName, ok := arg.(string); ok {
			if val, found := symbolTable[symName]; found {
				resolved[i] = val
				continue
			}
		}
		resolved[i] = arg
	}
	return resolved
}

// Stack retrieves the stack from the current execution context and returns it.
func (e *Executor) Stack() *stack.Stack {
	return e.ctx.Stack
}
