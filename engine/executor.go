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

		preserveReferences := handler.Name() == "store"

		resolvedArgs := e.resolveArgs(instr.Args, e.ctx.Symbols, preserveReferences)
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

	if inv.PipeTo != nil {
		pipeToID, ok := inv.PipeTo.(int)
		if !ok {
			return fmt.Errorf("pipeTo argument is not int, got: %T", inv.PipeTo)
		}

		if e.ctx.Stack.Len() == 0 {
			return fmt.Errorf("pipeTo failed: stack is empty, no value to pass")
		}

		finalValue := e.ctx.Stack.Pop()

		pipeToInv := types.Invocation{
			FunctionID: pipeToID,
			Args:       []any{finalValue},
		}

		return e.ExecuteInvocation(pipeToInv)
	}

	return nil
}

// resolveArgs resolves the arguments based on the symbol table.
// Ensures that symbol names remain intact for specific instructions like `store`.
func (e *Executor) resolveArgs(args []any, symbolTable map[string]any, preserveReferences bool) []any {
	resolved := make([]any, len(args))

	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			if preserveReferences {
				// For instructions like `store`, keep the symbol name as-is.
				resolved[i] = v
			} else {
				// Resolve symbol value from the table if found.
				if val, found := symbolTable[v]; found {
					resolved[i] = val
				} else {
					resolved[i] = fmt.Sprintf("UndefinedSymbol: %s", v)
				}
			}
		default:
			// Retain other types (e.g., int, float64) as-is.
			resolved[i] = v
		}
	}

	return resolved
}

// Stack retrieves the stack from the current execution context and returns it.
func (e *Executor) Stack() *stack.Stack {
	return e.ctx.Stack
}
