package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/env"
	"fmt"
	"strconv"
	"strings"
)

type Compiler struct {
	Source      string
	data        []string
	dataCounter int
	Environment *env.Environment[*Variable]
}

type Additional struct {
	Name string
	Id   int
	Type compiler.Type
	Raw  bool
}

func (a *Additional) String() string {
	return fmt.Sprintf("%s %s", a.Type, a.Name)
}

func New(file string) Compiler {
	return Compiler{
		Source:      file,
		Environment: env.New[*Variable](),
	}
}

func (c *Compiler) Compile(program *ast.Program) (string, error) {
	result := ""

	block, _ := c.CompileBlock(ast.BlockStatement{
		Body: program.Body,
		Loc:  program.Loc,
	}, false)

	for _, data := range c.data {
		result += data + "\n"
	}

	for _, block := range block {
		result += block
	}

	return result, nil
}

func (c *Compiler) CompileBlock(block ast.BlockStatement, labels bool) ([]string, *Additional) {
	result := make([]string, 0)

	id := c.Environment.TempCounter
	if labels {
		result = append(result, fmt.Sprintf("@block.start.%d", id))
	}

	for _, stmt := range block.Body {
		indentation := strings.Repeat("    ", 1)

		if stmt.Kind() == ast.FunctionDeclarationNode || labels {
			indentation = ""
		}

		compiled, _ := c.CompileStatement(stmt, nil)
		for _, compiled := range compiled {
			if strings.HasPrefix(compiled, "@") {
				result = append(result, compiled+"\n")
			} else {
				result = append(result, indentation+compiled+"\n")
			}
		}
	}

	if labels {
		result = append(result, fmt.Sprintf("@block.end.%d", id))

		c.Environment.TempCounter += 1
	}

	return result, &Additional{
		Id: id,
	}
}

func (c *Compiler) CompileStatement(stmt ast.Statement, expectedType *compiler.Type) ([]string, *Additional) {
	c.Environment.TempCounter += 1

	switch stmt := stmt.(type) {
	case *ast.IntLiteralStatement:
		val := strconv.FormatInt(int64(stmt.Value), 10)
		return []string{}, &Additional{
			Name: val,
			Type: compiler.Word,
			Raw:  true,
		}
	case *ast.FloatLiteralStatement:
		return c.CompileFloatLiteralStatement(stmt, expectedType)
	case *ast.StringLiteralStatement:
		id := c.dataCounter

		c.data = append(c.data, fmt.Sprintf("data $%s.%d = { b \"%s\", b 0 }", c.Environment.CurrentFunction.Name, id, stmt.Value))
		c.dataCounter += 1

		return []string{}, &Additional{
			Name: fmt.Sprintf("$%s.%d", c.Environment.CurrentFunction.Name, id),
			Type: compiler.String,
		} //fmt.Sprintf("l $%s.%d", c.Environment.CurrentFunction.Name, id)
	case *ast.DeclarationStatement:
		return c.CompileDeclarationStatement(stmt)
	case *ast.IdentifierLiteralStatement:
		variable := c.Environment.Get(stmt.Value)
		if variable == nil {
			dbg := debug.NewSourceLocation(c.Source, stmt.Loc.Row, stmt.Loc.Column)
			dbg.ThrowError(fmt.Sprintf("Variable '%s' is not defined", stmt.Value), true)
		}

		return []string{}, &Additional{
			Name: fmt.Sprintf("%%%s.%d", stmt.Value, variable.Id),
			Type: variable.Type,
		}
	case *ast.FunctionCall:
		return c.CompileFunctionCall(stmt, expectedType)
	case *ast.FunctionDeclaration:
		return c.CompileFunctionDeclaration(stmt), nil
	case *ast.ReturnStatement:
		return c.CompileReturnStatement(stmt, expectedType)
	case *ast.BlockStatement:
		return c.CompileBlock(*stmt, true)
	case *ast.BinaryExpression:
		return c.CompileBinaryExpression(stmt, expectedType)
	case *ast.UnaryExpression:
		return c.CompileUnaryExpression(stmt, expectedType)
	case *ast.IfStatement:
		return c.CompileIfStatement(stmt)
	case *ast.CompileTimeFunctionCall:
		return c.CompileCompileTimeFunctionCall(stmt)
	}

	fmt.Printf("Unknown statement: %T\n", stmt)

	return []string{}, nil
}
