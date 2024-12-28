package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) CompileBinaryExpression(stmt *ast.BinaryExpression, expectedType *compiler.Type) ([]string, *Additional) {
	name := fmt.Sprintf("%%tmp.%d", c.Environment.TempCounter)

	left, leftVar := c.CompileStatement(stmt.Left, expectedType)
	right, rightVar := c.CompileStatement(stmt.Right, expectedType)

	if leftVar.Type != rightVar.Type {
		dbg := debug.NewSourceLocation(c.Source, stmt.OperatorLoc.Row, stmt.OperatorLoc.Column)
		dbg.ThrowError(fmt.Sprintf("Cannot perform binary operation on two different types \"%s\" and \"%s\"!", leftVar.Type.Inspect(), rightVar.Type.Inspect()), true)
	}

	result := make([]string, 0)

	for _, l := range left {
		result = append(result, l)
	}

	for _, r := range right {
		result = append(result, r)
	}

	exp := fmt.Sprintf("%s =%s ", name, c.StoreType(leftVar.Type))

	switch stmt.Operator {
	case tokens.Plus:
		exp += "add"
	case tokens.Minus:
		exp += "sub"
	case tokens.Asterisk:
		exp += "mul"
	case tokens.Slash:
		exp += "div"
	case tokens.PercentSign:
		exp += "rem"
	case tokens.Ampersand:
		exp += "and"
	case tokens.VerticalLine:
		exp += "or"
	case tokens.CircumflexAccent:
		exp += "xor"
	case tokens.DoubleLessThan:
		exp += "shl"
	case tokens.DoubleGreaterThan:
		exp += "shr"
	case tokens.Equals:
		exp += "ceq" + c.StoreType(leftVar.Type).String()
	case tokens.LessThan:
		exp += "cslt" + c.StoreType(leftVar.Type).String()
	case tokens.LessThanOrEqual:
		exp += "csle" + c.StoreType(leftVar.Type).String()
	case tokens.GreaterThan:
		exp += "csgt" + c.StoreType(leftVar.Type).String()
	case tokens.GreaterThanOrEqual:
		exp += "csge" + c.StoreType(leftVar.Type).String()
	}

	exp += " " + c.StoreVal(leftVar) + ", " + c.StoreVal(rightVar)

	result = append(result, exp)

	result = append(result, "# ^ binary expression\n")

	return result, &Additional{
		Name: name,
		Type: leftVar.Type,
	} //fmt.Sprintf("l %s", name)
}
