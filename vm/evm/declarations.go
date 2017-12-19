package evm

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/ast"
)

func (e *GuardianEVM) traverseType(n *ast.TypeDeclarationNode) (code vmgen.Bytecode) {
	// do nothing
	return code
}

func (e *GuardianEVM) traverseClass(n *ast.ClassDeclarationNode) (code vmgen.Bytecode) {
	// create constructor hooks
	// create function hooks
	if n.Body.Declarations != nil {
		for _, d := range n.Body.Declarations.Map() {
			switch a := d.(type) {
			case *ast.ExplicitVarDeclarationNode:
				e.traverseExplicitVarDecl(a)
				break
			case *ast.LifecycleDeclarationNode:
				e.addLifecycleHook(n.Identifier, a)
				break
			case *ast.FuncDeclarationNode:
				e.addFunctionHook(n.Identifier, a)
				break
			default:
				e.traverse(a.(ast.Node))
			}
		}
	}
	return code
}

func (e *GuardianEVM) traverseInterface(n *ast.InterfaceDeclarationNode) (code vmgen.Bytecode) {
	// don't need to be interacted with
	// all interfaces are dealt with by the type system
	return code
}

func (e *GuardianEVM) traverseEnum(n *ast.EnumDeclarationNode) (code vmgen.Bytecode) {
	// don't create anything
	return code
}

func (e *GuardianEVM) traverseContract(n *ast.ContractDeclarationNode) (code vmgen.Bytecode) {
	// create hooks for functions
	// create hooks for constructors
	// create hooks for events
	// traverse everything else?
	for _, d := range n.Body.Declarations.Map() {
		switch a := d.(type) {
		case *ast.LifecycleDeclarationNode:
			e.addLifecycleHook(n.Identifier, a)
			break
		case *ast.FuncDeclarationNode:
			e.addFunctionHook(n.Identifier, a)
			break
		case *ast.EventDeclarationNode:
			e.addEventHook(n.Identifier, a)
			break
		default:
			e.traverse(a.(ast.Node))
		}
	}

	return code
}

func (e *GuardianEVM) addLifecycleHook(parent string, node *ast.LifecycleDeclarationNode) {
	//e.hooks[parent].hook[]
}

func (e *GuardianEVM) addFunctionHook(parent string, node *ast.FuncDeclarationNode) {
	/*e.hooks[parent][node.Identifier] = hook {
		name: node.Identifier,
	}*/

}

func (e *GuardianEVM) addEventHook(parent string, node *ast.EventDeclarationNode) {
	/*e.hooks[parent][node.Identifier] = hook{
		name: node.Identifier,
	}*/
}

func (e *GuardianEVM) addHook(name string) {
	h := hook{
		name: name,
	}
	if e.hooks == nil {
		e.hooks = make([]hook, 0)
	}
	e.hooks = append(e.hooks, h)
}

func (e *GuardianEVM) traverseEvent(n *ast.EventDeclarationNode) (code vmgen.Bytecode) {

	hook := string(EncodeName(n.Identifier))

	e.addHook(hook)

	code.Add("JUMPDEST")
	code.Add("LOG")
	return code
}

func (e *GuardianEVM) traverseParameters(params []*ast.ExplicitVarDeclarationNode) (code vmgen.Bytecode) {
	storage := false
	for _, p := range params {
		// function parameters are passed in on the stack and then assigned
		// default: memory, can be overriden to be storage
		// check if it's in storage
		for _, m := range p.Modifiers {
			if m == token.Storage {
				storage = true
			}
		}
		for _, i := range p.Identifiers {
			if storage {
				e.allocateStorage(i, p.Resolved.Size())
				//code.Push(e.lookupStorage(i))
				code.Add("SSTORE")
			} else {
				e.allocateMemory(i, p.Resolved.Size())
				//code.Push(uintAsBytes(location)...)
				code.Add("MSTORE")
			}
		}
		storage = false
	}
	return code
}

func (e *GuardianEVM) traverseFunc(n *ast.FuncDeclarationNode) (code vmgen.Bytecode) {

	// don't worry about hooking

	code.Add("JUMPDEST")

	var ps []*ast.ExplicitVarDeclarationNode
	for _, n := range n.Signature.Parameters {
		ps = append(ps, n.(*ast.ExplicitVarDeclarationNode))
	}

	params := e.traverseParameters(ps)

	code.Concat(params)

	for _, n := range n.Body.Sequence {
		code.Concat(e.traverse(n))
	}

	return code
}

func (e *GuardianEVM) traverseExplicitVarDecl(n *ast.ExplicitVarDeclarationNode) (code vmgen.Bytecode) {
	// variable declarations don't require storage (yet), just have to designate a slot
	storage := e.inStorage() || hasModifier(n.Modifiers, token.Storage)
	for _, id := range n.Identifiers {
		if storage {
			e.allocateStorage(id, n.Resolved.Size())
		} else {
			e.allocateMemory(id, n.Resolved.Size())
		}
	}
	return code
}
