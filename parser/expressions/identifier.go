package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseIdentifier(p Parser) ast.Statement {
	token := p.Consume()

	if p.Current().Kind == tokens.LeftParenthesis {
		return ParseFunctionCall(p, token)
	}

	if p.Current().Kind == tokens.Identifier {
		return ParseAssignment(p, token)
	}

	return &ast.IdentifierLiteralStatement{
		Value: token.Value,
		Loc:   p.Current().Location,
	}
}
