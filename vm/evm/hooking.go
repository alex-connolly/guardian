package evm

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/vmgen"
)

type hook struct {
	name     string
	position int
	bytecode vmgen.Bytecode
}

func (e *GuardianEVM) createFunctionBody(node *ast.FuncDeclarationNode) (body vmgen.Bytecode) {
	// all function bodies look the same
	// traverse the scope
	body.Concat(e.traverseScope(node.Body))

	// return address should be on top of the stack
	body.Add("JUMP")
	return body
}

func (e *GuardianEVM) createExternalFunctionComponents(node *ast.FuncDeclarationNode) (params, body vmgen.Bytecode) {
	// move params from calldata to memory

	return params, e.createFunctionBody(node)
}

func (e *GuardianEVM) createExternalParameters(node *ast.FuncDeclarationNode) (code vmgen.Bytecode) {
	offset := uint(0)
	for _, param := range node.Signature.Parameters {
		exp := param.(*ast.ExplicitVarDeclarationNode)
		for _ = range exp.Identifiers {
			offset += exp.Resolved.Size()

		}
	}
	return code
}

func (e *GuardianEVM) traverseExternalFunction(node *ast.FuncDeclarationNode) (code vmgen.Bytecode) {

	params := e.createExternalParameters(node)

	body := e.createFunctionBody(node)

	code.Concat(params)
	code.Concat(body)

	e.addExternalHook(node.Signature.Identifier, code)

	return code
}

/*
| return location |
| param 1 |
| param 2 |
*/
func (e *GuardianEVM) traverseInternalFunction(node *ast.FuncDeclarationNode) (code vmgen.Bytecode) {

	params := e.createInternalParameters(node)

	body := e.createFunctionBody(node)

	code.Concat(params)
	code.Concat(body)

	e.addInternalHook(node.Signature.Identifier, code)

	return code
}

func (e *GuardianEVM) createInternalParameters(node *ast.FuncDeclarationNode) (params vmgen.Bytecode) {
	// as internal functions can only be called from inside the contract
	// no need to have a hook
	// can just jump to the location
	params.Add("JUMPDEST")
	// load all parameters into memory blocks
	for _, param := range node.Signature.Parameters {
		exp := param.(*ast.ExplicitVarDeclarationNode)
		for _, i := range exp.Identifiers {
			e.allocateMemory(i, exp.Resolved.Size())
			// all parameters must be on the stack
			loc := e.lookupMemory(i)
			params.Concat(loc.retrieve())
		}
	}

	return params
}

func (e *GuardianEVM) traverseGlobalFunction(node *ast.FuncDeclarationNode) (code vmgen.Bytecode) {
	// hook here
	// get all parameters out of calldata and into memory
	// then jump over the parameter init in the internal declaration
	// leave the hook for later
	external := e.createExternalParameters(node)

	internal := e.createInternalParameters(node)

	body := e.createFunctionBody(node)

	code.Concat(pushMarker(external.Length()))

	code.Concat(external)

	code.Add("JUMPDEST")

	code.Concat(internal)

	code.Concat(body)

	e.addGlobalHook(node.Signature.Identifier, code)
	return code
}

func (e *GuardianEVM) addGlobalHook(id string, code vmgen.Bytecode) {
	if e.globalHooks == nil {
		e.globalHooks = make(map[string]hook)
	}
	e.globalHooks[id] = hook{
		name:     id,
		bytecode: code,
	}
}

func (e *GuardianEVM) addInternalHook(id string, code vmgen.Bytecode) {
	if e.internalHooks == nil {
		e.internalHooks = make(map[string]hook)
	}
	e.internalHooks[id] = hook{
		name:     id,
		bytecode: code,
	}
}

func (e *GuardianEVM) addExternalHook(id string, code vmgen.Bytecode) {
	if e.externalHooks == nil {
		e.externalHooks = make(map[string]hook)
	}
	e.externalHooks[id] = hook{
		name:     id,
		bytecode: code,
	}
}

func (e *GuardianEVM) addLifecycleHook(id string, code vmgen.Bytecode) {
	if e.lifecycleHooks == nil {
		e.lifecycleHooks = make(map[string]hook)
	}
	e.lifecycleHooks[id] = hook{
		name:     id,
		bytecode: code,
	}
}
