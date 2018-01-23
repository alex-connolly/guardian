package validator

import (
	"github.com/end-r/guardian/util"

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
	// TODO: fundamental: define operator or not
	// valid assignments must have
	// 1. valid left hand expression (cannot be a call, literal, slice)
	// 2. types of left assignable to right
	// 3. right all valid expressions

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
			if !typing.AssignableTo(right, left, true) {
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
						v.declareContextualVar(id.Start(), id.Name, id.Resolved)
					}

				}
			}
		}

	} else {
		if !leftTuple.Compare(rightTuple) {
			v.addError(node.Left[0].Start(), errInvalidAssignment, typing.WriteType(leftTuple), typing.WriteType(rightTuple))
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
							v.declareContextualVar(id.Start(), id.Name, rightTuple.Types[i])
						}
					}
				}
			}
		}
	}
}

func (v *Validator) validateAssignmentWithoutDeclaring(node *ast.AssignmentStatementNode) (types typing.TypeMap, locations map[string]util.Location) {
	types = make(typing.TypeMap)

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
			if !typing.AssignableTo(right, left, true) {
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
						types[id.Name] = id.Resolved
						locations[id.Name] = left.Start()
					}
				}
			}
		}
	} else {
		if !leftTuple.Compare(rightTuple) {
			v.addError(node.Left[0].Start(), errInvalidAssignment, typing.WriteType(leftTuple), typing.WriteType(rightTuple))
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
							v.declareContextualVar(id.Start(), id.Name, rightTuple.Types[i])
						}
					}
				}
			}
		}
	}
	return types, locations
}

func (v *Validator) validateIfStatement(node *ast.IfStatementNode) {

	var types typing.TypeMap
	var locs map[string]util.Location
	if node.Init != nil {
		types, locs = v.validateAssignmentWithoutDeclaring(node.Init.(*ast.AssignmentStatementNode))
	}

	for _, cond := range node.Conditions {
		// condition must be of type bool
		v.requireType(cond.Condition.Start(), typing.Boolean(), v.resolveExpression(cond.Condition))
		v.validateScope(node, cond.Body, types, locs)
	}

	if node.Else != nil {
		v.validateScope(node, node.Else, types, locs)
	}
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
		v.requireType(expr.Start(), switchType, v.resolveExpression(expr))
	}
	v.validateScope(clause, clause.Block, nil, nil)
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
	vars := make(typing.TypeMap)
	locs := make(map[string]util.Location)
	gen := v.resolveExpression(node.Producer)
	var req int
	switch a := gen.(type) {
	case *typing.Map:
		// maps must handle k, v in MAP
		req = 2
		if len(node.Variables) != req {
			v.addError(node.Begin, errInvalidForEachVariables, len(node.Variables), req)
		} else {
			vars[node.Variables[0]] = a.Key
			locs[node.Variables[0]] = node.Start()
			vars[node.Variables[1]] = a.Value
			locs[node.Variables[1]] = node.Start()
		}
		break
	case *typing.Array:
		// arrays must handle i, v in ARRAY
		req = 2
		if len(node.Variables) != req {
			v.addError(node.Start(), errInvalidForEachVariables, len(node.Variables), req)
		} else {
			vars[node.Variables[0]] = v.LargestNumericType(false)
			locs[node.Variables[0]] = node.Start()
			vars[node.Variables[1]] = a.Value
			locs[node.Variables[1]] = node.Start()
		}
		break
	default:
		v.addError(node.Start(), errInvalidForEachType, typing.WriteType(gen))
	}

	v.validateScope(node, node.Block, vars, locs)

}

func (v *Validator) validateForStatement(node *ast.ForStatementNode) {

	var vars typing.TypeMap
	var locs map[string]util.Location

	if node.Init != nil {
		vars, locs = v.validateAssignmentWithoutDeclaring(node.Init)
	}

	// cond statement must be a boolean
	v.requireType(node.Cond.Start(), typing.Boolean(), v.resolveExpression(node.Cond))

	// post statement must be valid
	if node.Post != nil {
		v.validateStatement(node.Post)
	}

	v.validateScope(node, node.Block, vars, locs)
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
		v.declareContextualType(node.Start(), node.Alias, v.createPackageType(node.Path))
	} else {
		v.declareContextualType(node.Start(), trimPath(node.Path), v.createPackageType(node.Path))
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
		v.addError(node.Start(), errDuplicatePackageName, node.Name, v.packageName)
	}
}
