package statements

import (
	"axia/guardian/go/compiler/parser"
	"testing"
)

func TestSimpleSwitchEmpty(t *testing.T) {
	p := parser.ParseString("switch x {}")
}

func TestExclusiveSwitchEmpty(t *testing.T) {
	p := parser.ParseString("exclusive switch x {}")
}
