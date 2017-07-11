package parser

type construct struct {
	name  string
	is    func(*parser) bool
	parse func(*parser)
}

func getPrimaryConstructs() []construct {
	return []construct{
		construct{"class declaration", isClassDeclaration, parseClassDeclaration},
		construct{"contract declaration", isContractDeclaration, parseContractDeclaration},
		construct{"interface declaration", isInterfaceDeclaration, parseInterfaceDeclaration},
		construct{"func declaration", isFuncDeclaration, parseFuncDeclaration},

		construct{"if statement", isIfStatement, parseIfStatement},
		construct{"for statement", isForStatement, parseForStatement},
		construct{"assignment statement", isAssignmentStatement, parseAssignmentStatement},
		construct{"return statment", isReturnStatement, parseReturnStatement},
		construct{"switch statement", isSwitchStatement, parseSwitchStatement},
		construct{"case statement", isCaseStatement, parseCaseStatement},
	}
}
