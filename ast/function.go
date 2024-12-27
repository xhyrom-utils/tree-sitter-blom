package ast

import (
	"blom/compiler"
	"blom/tokens"
)

type FunctionArgument struct {
	Name string
	Type compiler.Type
}

type FunctionDeclaration struct {
	Name        string
	Arguments   []FunctionArgument
	Annotations []Annotation
	ReturnType  compiler.Type
	Body        *BlockStatement
	Loc         tokens.Location
}

func (f FunctionDeclaration) Kind() NodeKind {
	return FunctionDeclarationNode
}

func (f FunctionDeclaration) Location() tokens.Location {
	return f.Loc
}

func (f *FunctionDeclaration) SetLocation(row uint64, column uint64) {
	f.Loc.Row = row
	f.Loc.Column = column
}

type FunctionCall struct {
	Name       string
	Parameters []Expression
	Loc        tokens.Location
}

func (f FunctionCall) Kind() NodeKind {
	return FunctionCallNode
}

func (f FunctionCall) Location() tokens.Location {
	return f.Loc
}

func (f *FunctionCall) SetLocation(row uint64, column uint64) {
	f.Loc.Row = row
	f.Loc.Column = column
}

type CompileTimeFunctionCall struct {
	Name       string
	Parameters []Expression
	Loc        tokens.Location
}

func (c CompileTimeFunctionCall) Kind() NodeKind {
	return CompileTimeFunctionCallNode
}

func (c CompileTimeFunctionCall) Location() tokens.Location {
	return c.Loc
}

func (c *CompileTimeFunctionCall) SetLocation(row uint64, column uint64) {
	c.Loc.Row = row
	c.Loc.Column = column
}
