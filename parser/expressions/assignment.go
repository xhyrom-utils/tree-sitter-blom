package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses an assignment statement that can have form:
// <type> <identifier> = <expression>;
// <type> <identifier>;
// <identifier> = <expression>;
func ParseAssignment(p Parser, redeclaration bool) (*ast.DeclarationStatement, error) {

	if redeclaration {
		name := p.Consume()

		p.Consume()

		value, _ := p.ParseExpression()

		if p.Consume().Kind != tokens.Semicolon {
			dbg := debug.NewSourceLocationFromExpression(p.Source(), value)
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Add semicolon", ";"))
		}

		return &ast.DeclarationStatement{
			Name:          name.Value,
			Value:         value,
			Redeclaration: true,
		}, nil
	}

	valueType := p.Consume()

	name := p.Consume()
	var value ast.Expression = nil

	right := p.Consume()

	if right.Kind != tokens.Assign && right.Kind != tokens.Semicolon {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column+1)
		dbg.ThrowError("Expected assignment or semicolon", true, debug.NewHint("Add semicolon", ";"), debug.NewHint("Initialize variable", " = 0;"))
	}

	if right.Kind == tokens.Assign {
		exp, _ := p.ParseExpression()
		value = exp
	}

	if p.Consume().Kind != tokens.Semicolon {
		if value != nil {
			dbg := debug.NewSourceLocationFromExpression(p.Source(), value)
			dbg.ThrowError("Expected semicolon", true)
		} else {
			dbg := debug.NewSourceLocation(p.Source(), right.Location.Row, right.Location.Column+1)
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Add semicolon", ";"))
		}
	}

	return &ast.DeclarationStatement{
		Name:          name.Value,
		Value:         value,
		Redeclaration: false,
		Type:          int(valueType.Kind),
		Loc:           right.Location,
	}, nil
}
