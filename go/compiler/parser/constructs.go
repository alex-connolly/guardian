package parser

type construct struct {
	name       string
	identifier func(*parser) bool
	processor  func(*parser)
}

func getConstructs() []construct {
	return []construct{
		construct{"class declaration", isClassDeclaration, processClassDeclaration},
	}
}
