package tokens

import "slices"

type TokenKind int

const (
	Eof TokenKind = iota
	Illegal

	// Identifiers and literals
	Identifier
	CharLiteral
	StringLiteral
	IntLiteral
	FloatLiteral
	BooleanLiteral

	// Statements
	Assign
	If
	Else

	// Operators
	Equals
	Plus
	Minus
	Asterisk
	Slash
	PercentSign
	Ampersand
	VerticalLine
	CircumflexAccent
	Tilde
	And
	Or
	LessThan
	DoubleLessThan
	LessThanOrEqual
	GreaterThan
	DoubleGreaterThan
	GreaterThanOrEqual

	// Delimiters
	Dot
	Comma
	Colon
	Semicolon
	AtMark
	LeftParenthesis
	RightParenthesis
	LeftSquareBracket
	RightSquareBracket
	LeftCurlyBracket
	RightCurlyBracket

	// Keywords
	Fun
	Return
)

var tokens = []string{
	Eof:                "EOF",
	Illegal:            "Illegal",
	Identifier:         "Identifier",
	CharLiteral:        "CharLiteral",
	StringLiteral:      "StringLiteral",
	IntLiteral:         "IntLiteral",
	FloatLiteral:       "FloatLiteral",
	BooleanLiteral:     "BooleanLiteral",
	Assign:             "=",
	If:                 "if",
	Else:               "else",
	Equals:             "==",
	Plus:               "+",
	Minus:              "-",
	Asterisk:           "*",
	Slash:              "/",
	PercentSign:        "%",
	Ampersand:          "&",
	VerticalLine:       "|",
	CircumflexAccent:   "^",
	Tilde:              "~",
	And:                "and",
	Or:                 "or",
	LessThan:           "<",
	DoubleLessThan:     "<<",
	LessThanOrEqual:    "<=",
	GreaterThan:        ">",
	DoubleGreaterThan:  ">>",
	GreaterThanOrEqual: ">=",
	Dot:                ".",
	Comma:              ",",
	Colon:              ":",
	Semicolon:          ";",
	AtMark:             "@",
	LeftParenthesis:    "(",
	RightParenthesis:   ")",
	LeftSquareBracket:  "[",
	RightSquareBracket: "]",
	LeftCurlyBracket:   "{",
	RightCurlyBracket:  "}",
	Fun:                "fun",
	Return:             "return",
}

var reserved = []string{
	If:     "if",
	Else:   "else",
	Fun:    "fun",
	Return: "return",
}

func (t TokenKind) String() string {
	return tokens[t]
}

func FromIdentifier(identifier string) TokenKind {
	if identifier == "true" || identifier == "false" {
		return BooleanLiteral
	}

	index := slices.Index(reserved, identifier)
	if index == -1 {
		return Illegal
	}

	return TokenKind(index)
}

type Precedence int

const (
	LowestPrecedence Precedence = iota
	OrPrecedence
	AndPrecedence
	BitwiseOrPrecedence
	BitwiseXorPrecedence
	BitwiseAndPrecedence
	EqualityPrecedence
	RelationalPrecedence
	ShiftPrecedence
	AdditivePrecedence
	MultiplicativePrecedence
	HighestPrecedence
)

func (kind TokenKind) Precedence() Precedence {
	switch kind {
	case Or:
		return OrPrecedence
	case And:
		return AndPrecedence
	case VerticalLine:
		return BitwiseOrPrecedence
	case CircumflexAccent:
		return BitwiseXorPrecedence
	case Ampersand:
		return BitwiseAndPrecedence
	case Equals:
		return EqualityPrecedence
	case LessThan, LessThanOrEqual, GreaterThan, GreaterThanOrEqual:
		return RelationalPrecedence
	case DoubleLessThan, DoubleGreaterThan:
		return ShiftPrecedence
	case Plus, Minus:
		return AdditivePrecedence
	case Asterisk, Slash, PercentSign:
		return MultiplicativePrecedence
	default:
		return LowestPrecedence
	}
}

type Location struct {
	File string
	Row  int
	Col  int
}

type Token struct {
	Kind     TokenKind
	Location Location
	Value    string
}

func (l *Location) Copy() Location {
	return Location{
		File: l.File,
		Row:  l.Row,
		Col:  l.Col,
	}
}
