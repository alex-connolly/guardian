package parser

type construct struct {
	name  string
	is    func(*Parser) bool
	parse func(*Parser)
}

func getPrimaryConstructs() []construct {

	standards := []construct{
		construct{"ignored", isIgnored, parseIgnored},
		construct{"modifiers", isModifier, parseModifiers},
		construct{"annotations", isAnnotation, parseAnnotation},
		construct{"group", isGroup, parseGroup},
		construct{"new line", isNewLine, parseNewLine},

		construct{"explict var declaration", isExplicitVarDeclaration, parseExplicitVarDeclaration},
		construct{"class declaration", isClassDeclaration, parseClassDeclaration},
		construct{"contract declaration", isContractDeclaration, parseContractDeclaration},
		construct{"interface declaration", isInterfaceDeclaration, parseInterfaceDeclaration},
		construct{"func declaration", isFuncDeclaration, parseFuncDeclaration},
		construct{"lifecycle declaration", isLifecycleDeclaration, parseLifecycleDeclaration},
		construct{"enum declaration", isEnumDeclaration, parseEnumDeclaration},
		construct{"type declaration", isTypeDeclaration, parseTypeDeclaration},
		construct{"event declaration", isEventDeclaration, parseEventDeclaration},

		construct{"if statement", isIfStatement, parseIfStatement},
		construct{"for each statement", isForEachStatement, parseForEachStatement},
		construct{"for statement", isForStatement, parseForStatement},
		construct{"return statment", isReturnStatement, parseReturnStatement},
		construct{"switch statement", isSwitchStatement, parseSwitchStatement},
		construct{"case statement", isCaseStatement, parseCaseStatement},
		construct{"flow statement", isFlowStatement, parseFlowStatement},
		construct{"import statement", isImportStatement, parseImportStatement},
		construct{"package statement", isPackageStatement, parsePackageStatement},
	}
	return standards
}
