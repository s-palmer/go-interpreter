// define tokens that the lexer will output
// define Token data struct
// need type attribute to distinguish b/w diff types of tokens
// need field to hold literal value of token so we can reuse it

package token

// defined this as a string so we can use many diff values as TokenTypes and means we can distinguish between different ypes of tokens
// might not be most performant but for this example, it will suffice
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// unknown token
	ILLEGAL = "ILLEGAL"
	// end of file
	EOF = "EOF"

	// identifiers and literals
	IDENT = "IDENT" // add, foobar, x, y etc
	INT   = "INT"   // 123123

	// operators
	ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"
    LT = "<"
    GT = ">"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE	 = "TRUE"
	FALSE	 = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN	 = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}