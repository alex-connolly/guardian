package evm

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/typing"
	"github.com/end-r/vmgen"
)

type hook struct {
	name     string
	position int
	bytecode []byte
}

type callable struct {
	name     string
	position int
	bytecode []byte
}

func (e *GuardianEVM) addExternalHook(node *ast.FuncDeclarationNode) {
	// as external functions can only be called from outside the contract
	// all parameters will be in calldata
	var code vmgen.Bytecode
	f := node.Resolved.(typing.Func)
	offset := 0
	for _, p := range f.Params.Types {
		// get next 'size' bytes of calldata
		size := p.Size()

	}
}

func (e *GuardianEVM) createFunctionBody(node *ast.FuncDeclarationNode) (body vmgen.Bytecode){
    // all function bodies look the same
    // traverse the scope
	body.Concat(node.Body)

	// return address should be on top of the stack
	body.Add("JUMP")
}

func (e *GuardianEVM) createExternalFunctionComponents(node *ast.FuncDeclarationNode) (params, body vmgen.Bytecode){


}
/*
| return location |
| param 1 |
| param 2 |
*/
func (e *GuardianEVM) createInternalFunction(node *ast.FuncDeclarationNode) (code vmgen.Bytecode){
    p, b := e.createInternalFunctionComponents(node)
	code.Concat(p)
    code.Concat(b)
    return code
}

func (e *GuardianEVM) createInternalFunctionComponents(node *ast.FuncDeclarationNode) (params, body vmgen.Bytecode){
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

    return code, e.createFunctionBody(node)
}

func (e *GuardianEVM) createGlobalFunction(node *ast.FuncDeclarationNode) (code vmgen.Bytecode) {
    // hook here
    // get all parameters out of calldata and into memory
    // then jump over the parameter init in the internal declaration
    for _,

    // add normal function
    params, body := e.createInternalFunctionComponents(node)

    code.Concat(pushMarker(len(params)))

    code.Concat(params)

    code.Add("JUMPDEST")

    code.Concat(body)

    e.addGlobalHook(node.Signature.Identifier, code)
}

func (e *GuardianEVM) addGlobalHook(id string, code vmgen.Bytecode){

}
