package ast

import "blom/tokens"

type IdentifierLiteralStatement struct {
	Value string
	Loc   tokens.Location
}

func (l IdentifierLiteralStatement) Kind() NodeKind {
	return IdentifierLiteralNode
}

func (l IdentifierLiteralStatement) Location() tokens.Location {
	return l.Loc
}

func (l IdentifierLiteralStatement) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}

type CharLiteralStatement struct {
	Value rune
	Loc   tokens.Location
}

func (l CharLiteralStatement) Kind() NodeKind {
	return CharLiteralNode
}

func (l CharLiteralStatement) Location() tokens.Location {
	return l.Loc
}

func (l CharLiteralStatement) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}

type StringLiteralStatement struct {
	Value string
	Loc   tokens.Location
}

func (l StringLiteralStatement) Kind() NodeKind {
	return StringLiteralNode
}

func (l StringLiteralStatement) Location() tokens.Location {
	return l.Loc
}

func (l StringLiteralStatement) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}

type IntLiteralStatement struct {
	Value int32
	Loc   tokens.Location
}

func (l IntLiteralStatement) Kind() NodeKind {
	return IntLiteralNode
}

func (l IntLiteralStatement) Location() tokens.Location {
	return l.Loc
}

func (l IntLiteralStatement) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}

type FloatLiteralStatement struct {
	Value float32
	Loc   tokens.Location
}

func (l FloatLiteralStatement) Kind() NodeKind {
	return FloatLiteralNode
}

func (l FloatLiteralStatement) Location() tokens.Location {
	return l.Loc
}

func (l FloatLiteralStatement) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}
