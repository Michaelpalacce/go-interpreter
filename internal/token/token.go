package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// List of available tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT is an identifier like add, foobar, x, y, etc
	IDENT = "IDENT"

	// 1234567890
	INT    = "INT"
	TRUE   = "true"
	FALSE  = "false"
	IF     = "if"
	ELSE   = "else"
	RETURN = "return"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	EQ     = "=="
	NOT_EQ = "!="

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	// Conditionals
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
