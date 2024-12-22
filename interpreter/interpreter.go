package interpreter

import (
	"blom/ast"
	"blom/env"
	"blom/interpreter/expressions"
	"fmt"
)

type Interpreter struct{}

func New() *Interpreter {
	return &Interpreter{}
}

func (intrepreter *Interpreter) Interpret(program *ast.Program) env.Object {
	program.Body = append(program.Body, &ast.FunctionCall{
		Name:       "main",
		Parameters: []ast.Statement{},
	})

	return intrepreter.InterpretBlock(program.Body, env.New())
}

func (interpreter *Interpreter) InterpretBlock(body []ast.Statement, environment *env.Environment) env.Object {
	envi := env.New(*environment)

	for _, stmt := range body {
		value := interpreter.InterpretStatement(stmt, envi)

		switch stmt.(type) {
		case *ast.FunctionCall, ast.FunctionCall:
			if stmt.(*ast.FunctionCall).Name == "main" {
				return value
			}
		case *ast.ReturnStatement, ast.ReturnStatement:
			return value
		}
	}

	return nil
}

func (intrepreter *Interpreter) InterpretStatement(stmt ast.Statement, environment *env.Environment) env.Object {
	switch stmt := stmt.(type) {
	case *ast.IntLiteralStatement:
		return &env.IntegerObject{Value: stmt.Value}
	case ast.IntLiteralStatement:
		return &env.IntegerObject{Value: stmt.Value}
	case *ast.FloatLiteralStatement:
		return &env.FloatObject{Value: stmt.Value}
	case ast.FloatLiteralStatement:
		return &env.FloatObject{Value: stmt.Value}
	case *ast.IdentifierLiteralStatement:
		return environment.Get(stmt.Value)
	case ast.IdentifierLiteralStatement:
		return environment.Get(stmt.Value)
	case *ast.FunctionDeclaration:
		expressions.InterpretFunctionDeclaration(intrepreter, environment, stmt)
	case ast.FunctionDeclaration:
		expressions.InterpretFunctionDeclaration(intrepreter, environment, &stmt)
	case *ast.BinaryExpression:
		return expressions.InterpretBinaryExpression(intrepreter, environment, stmt)
	case ast.BinaryExpression:
		return expressions.InterpretBinaryExpression(intrepreter, environment, &stmt)
	case *ast.UnaryExpression:
		return expressions.InterpretUnaryExpression(intrepreter, environment, stmt)
	case ast.UnaryExpression:
		return expressions.InterpretUnaryExpression(intrepreter, environment, &stmt)
	case *ast.DeclarationStatement:
		expressions.InterpretDeclarationStatement(intrepreter, environment, stmt)
	case ast.DeclarationStatement:
		expressions.InterpretDeclarationStatement(intrepreter, environment, &stmt)
	case *ast.ReturnStatement:
		return expressions.InterpretReturnStatement(intrepreter, environment, stmt)
	case ast.ReturnStatement:
		return expressions.InterpretReturnStatement(intrepreter, environment, &stmt)
	case *ast.FunctionCall:
		return expressions.InterpretFunctionCall(intrepreter, environment, stmt)
	case ast.FunctionCall:
		return expressions.InterpretFunctionCall(intrepreter, environment, &stmt)
	default:
		fmt.Printf("Unknown statement type: %T\n", stmt)
	}

	return nil
}

func (interpreter *Interpreter) InterpretExpression(expression ast.Expression, environment *env.Environment) env.Object {
	switch expression := expression.(type) {
	case *ast.IntLiteralStatement:
		return &env.IntegerObject{Value: expression.Value}
	case ast.IntLiteralStatement:
		return &env.IntegerObject{Value: expression.Value}
	case *ast.FloatLiteralStatement:
		return &env.FloatObject{Value: expression.Value}
	case ast.FloatLiteralStatement:
		return &env.FloatObject{Value: expression.Value}
	case *ast.IdentifierLiteralStatement:
		return environment.Get(expression.Value)
	case ast.IdentifierLiteralStatement:
		return environment.Get(expression.Value)
	case *ast.BinaryExpression:
		return expressions.InterpretBinaryExpression(interpreter, environment, expression)
	case ast.BinaryExpression:
		return expressions.InterpretBinaryExpression(interpreter, environment, &expression)
	case *ast.UnaryExpression:
		return expressions.InterpretUnaryExpression(interpreter, environment, expression)
	case ast.UnaryExpression:
		return expressions.InterpretUnaryExpression(interpreter, environment, &expression)
	case *ast.FunctionCall:
		return expressions.InterpretFunctionCall(interpreter, environment, expression)
	case ast.FunctionCall:
		return expressions.InterpretFunctionCall(interpreter, environment, &expression)
	default:
		fmt.Printf("Unknown expression type: %T\n", expression)
	}

	return nil
}
