package validator

import (
	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func (v *Validator) validateStatement(node ast.Node) {
	switch n := node.(type) {
	case *ast.AssignmentStatementNode:
		v.validateAssignment(n)
		break
	case *ast.ForStatementNode:
		v.validateForStatement(n)
		break
	case *ast.IfStatementNode:
		v.validateIfStatement(n)
		break
	case *ast.ReturnStatementNode:
		v.validateReturnStatement(n)
		break
	case *ast.SwitchStatementNode:
		v.validateSwitchStatement(n)
		break
	case *ast.ForEachStatementNode:
		v.validateForEachStatement(n)
		break
	case *ast.ImportStatementNode:
		v.validateImportStatement(n)
		return
	case *ast.PackageStatementNode:
		v.validatePackageStatement(n)
		return
	}
	v.finishedImports = true
}

func (v *Validator) validateAssignment(node *ast.AssignmentStatementNode) {

	for _, l := range node.Left {
		if l == nil {
			v.addError(node.Start(), errUnknown)
			return
		} else {
			switch l.Type() {
			case ast.CallExpression, ast.Literal, ast.MapLiteral,
				ast.ArrayLiteral, ast.SliceExpression, ast.FuncLiteral:
				v.addError(l.Start(), errInvalidExpressionLeft)
			}
		}
	}

	leftTuple := v.ExpressionTuple(node.Left)
	rightTuple := v.ExpressionTuple(node.Right)
	if len(leftTuple.Types) > len(rightTuple.Types) && len(rightTuple.Types) == 1 {
		right := rightTuple.Types[0]

		for _, left := range leftTuple.Types {
			if !v.vm.Assignable(v, left, right, node.Right[0]) {
				v.addError(node.Left[0].Start(), errInvalidAssignment, typing.WriteType(left), typing.WriteType(right))
			}
		}

		for i, left := range node.Left {
			if leftTuple.Types[i] == typing.Unknown() {
				if id, ok := left.(*ast.IdentifierNode); ok {
					ty := rightTuple.Types[0]
					id.Resolved = ty
					id.Resolved.SetModifiers(nil)
					ignored := "_"
					if id.Name != ignored {
						v.declareVar(id.Start(), id.Name, id.Resolved)
					}

				}
			}
		}

	} else {
		if len(leftTuple.Types) == len(rightTuple.Types) {
			// count helps to handle: a, b, c, d = producesTwo(), 6, 7
			// first two rely on the same expression
			count := 0
			remaining := 0
			for i, left := range leftTuple.Types {
				right := rightTuple.Types[i]
				if !v.vm.Assignable(v, left, right, node.Right[count]) {
					v.addError(node.Start(), errInvalidAssignment, typing.WriteType(leftTuple), typing.WriteType(rightTuple))
					break
				}
				if remaining == 0 {
					if node.Right[count] == nil {
						count++
					} else {
						switch a := node.Right[count].ResolvedType().(type) {
						case *typing.Tuple:
							remaining = len(a.Types) - 1
							break
						default:
							count++
						}
					}

				} else {
					remaining--
				}

			}
		} else {
			v.addError(node.Start(), errInvalidAssignment, typing.WriteType(leftTuple), typing.WriteType(rightTuple))
		}

		// length of left tuple should always equal length of left
		// this is because tuples are not first class types
		// cannot assign to tuple expressions
		if len(node.Left) == len(rightTuple.Types) {
			for i, left := range node.Left {
				if leftTuple.Types[i] == typing.Unknown() {
					if id, ok := left.(*ast.IdentifierNode); ok {
						id.Resolved = rightTuple.Types[i]
						if id.Name != "_" {
							//fmt.Printf("Declaring %s as %s\n", id.Name, typing.WriteType(rightTuple.Types[i]))
							v.declareVar(id.Start(), id.Name, id.Resolved)
						}
					}
				}
			}
		}
	}
}

func (v *Validator) validateIfStatement(node *ast.IfStatementNode) {

	v.openScope(nil, nil)

	if node.Init != nil {
		v.validateAssignment(node.Init.(*ast.AssignmentStatementNode))
	}

	for _, cond := range node.Conditions {
		// condition must be of type bool
		v.requireType(cond.Condition.Start(), typing.Boolean(), v.resolveExpression(cond.Condition))
		v.validateScope(node, cond.Body)
	}

	if node.Else != nil {
		v.validateScope(node, node.Else)
	}

	v.closeScope()
}

func (v *Validator) validateSwitchStatement(node *ast.SwitchStatementNode) {

	// no switch expression --> booleans
	switchType := typing.Boolean()

	if node.Target != nil {
		switchType = v.resolveExpression(node.Target)
	}

	// target must be matched by all cases
	for _, node := range node.Cases.Sequence {
		if node.Type() == ast.CaseStatement {
			v.validateCaseStatement(switchType, node.(*ast.CaseStatementNode))
		}
	}

}

func (v *Validator) validateCaseStatement(switchType typing.Type, clause *ast.CaseStatementNode) {
	for _, expr := range clause.Expressions {
		t := v.resolveExpression(expr)
		if !v.vm.Assignable(v, switchType, t, expr) {
			v.addError(clause.Start(), errInvalidSwitchTarget, typing.WriteType(switchType), typing.WriteType(t))
		}

	}
	v.validateScope(clause, clause.Block)
}

func (v *Validator) validateReturnStatement(node *ast.ReturnStatementNode) {
	for c := v.scope; c != nil; c = c.parent {
		if c.context != nil {
			switch a := c.context.(type) {
			case *ast.FuncDeclarationNode:
				results := a.Resolved.(*typing.Func).Results
				returned := v.ExpressionTuple(node.Results)
				if (results == nil || len(results.Types) == 0) && len(returned.Types) > 0 {
					v.addError(node.Start(), errInvalidReturnFromVoid, typing.WriteType(returned), a.Signature.Identifier)
					return
				}
				if !typing.AssignableTo(results, returned, false) {
					v.addError(node.Start(), errInvalidReturn, typing.WriteType(returned), a.Signature.Identifier, typing.WriteType(results))
				}
				return
			case *ast.FuncLiteralNode:
				results := a.Resolved.(*typing.Func).Results
				returned := v.ExpressionTuple(node.Results)
				if (results == nil || len(results.Types) == 0) && len(returned.Types) > 0 {
					v.addError(node.Start(), errInvalidReturnFromVoid, typing.WriteType(returned), "literal")
					return
				}
				if !typing.AssignableTo(results, returned, false) {
					v.addError(node.Start(), errInvalidReturn, typing.WriteType(returned), "literal", typing.WriteType(results))
				}
				return
			}
		}
	}
	v.addError(node.Start(), errInvalidReturnStatementOutsideFunc)
}

func (v *Validator) validateForEachStatement(node *ast.ForEachStatementNode) {
	// get type of

	v.openScope(nil, nil)

	gen := v.resolveExpression(node.Producer)
	var req int
	switch a := gen.(type) {
	case *typing.Map:
		// maps must handle k, v in MAP
		req = 2
		if len(node.Variables) != req {
			v.addError(node.Begin, errInvalidForEachVariables, len(node.Variables), req)
		} else {
			v.declareVar(node.Start(), node.Variables[0], a.Key)
			v.declareVar(node.Start(), node.Variables[1], a.Value)
		}
		break
	case *typing.Array:
		// arrays must handle i, v in ARRAY
		req = 2
		if len(node.Variables) != req {
			v.addError(node.Start(), errInvalidForEachVariables, len(node.Variables), req)
		} else {
			v.declareVar(node.Start(), node.Variables[0], v.LargestNumericType(false))
			v.declareVar(node.Start(), node.Variables[1], a.Value)
		}
		break
	default:
		v.addError(node.Start(), errInvalidForEachType, typing.WriteType(gen))
	}

	v.validateScope(node, node.Block)

	v.closeScope()

}

func (v *Validator) validateForStatement(node *ast.ForStatementNode) {

	v.openScope(nil, nil)

	if node.Init != nil {
		v.validateAssignment(node.Init)
	}

	// cond statement must be a boolean
	v.requireType(node.Cond.Start(), typing.Boolean(), v.resolveExpression(node.Cond))

	// post statement must be valid
	if node.Post != nil {
		v.validateStatement(node.Post)
	}

	v.validateScope(node, node.Block)

	v.closeScope()
}

func (v *Validator) createPackageType(path string) *typing.Package {
	scope, errs := ValidatePackage(v.vm, path)
	if errs != nil {
		v.errs = append(v.errs, errs...)
	}
	pkg := new(typing.Package)
	pkg.Variables = scope.variables
	pkg.Types = scope.types
	return pkg
}

func trimPath(n string) string {
	lastSlash := 0
	for i := 0; i < len(n); i++ {
		if n[i] == '/' {
			lastSlash = i
		}
	}
	return n[lastSlash:]
}

func (v *Validator) validateImportStatement(node *ast.ImportStatementNode) {
	if v.finishedImports {
		v.addError(node.Start(), errFinishedImports)
	}
	if node.Alias != "" {
		v.declareType(node.Start(), node.Alias, v.createPackageType(node.Path))
	} else {
		v.declareType(node.Start(), trimPath(node.Path), v.createPackageType(node.Path))
	}
}

func (v *Validator) validatePackageStatement(node *ast.PackageStatementNode) {
	if node.Name == "" {
		v.addError(node.Start(), errInvalidPackageName, node.Name)
		return
	}
	if v.packageName == "" {
		v.packageName = node.Name
	} else {
		if v.packageName != node.Name {
			v.addError(node.Start(), errDuplicatePackageName, node.Name, v.packageName)
		}
	}
}
