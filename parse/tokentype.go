package parse

type TokenType string

const (
	NULLTYPE TokenType = ""
	SEPERATOR
	IDENTIFIER
	OPERATOR
	STRING
	NUMBER
)
