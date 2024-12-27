package expressions

import "blom/ast"

func ParseUnary(p Parser) *ast.UnaryExpression {
	operator := p.Consume()
	operand := p.ParsePrimaryExpression()

	return &ast.UnaryExpression{
		Operator: operator.Kind,
		Operand:  operand,
		Loc:      operator.Location,
	}
}
