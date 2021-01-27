package lispingo

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	// TokenEOF is token for end of file
	TokenEOF = iota
	// TokenIgnored is token for ignored characters
	TokenIgnored
	// TokenLeftParen is token for (
	TokenLeftParen
	// TokenRightParen is token for )
	TokenRightParen
	// TokenSingleQuote is token for '
	TokenSingleQuote
	// TokenDoubleQuote is token for "
	TokenDoubleQuote
	// TokenDoubleQuotePair is token for ""
	TokenDoubleQuotePair
	// TokenSymbol is token for symbols
	TokenSymbol
	// TokenInt is token for integers
	TokenInt
)

var tokenNameMap = map[int]string{
	TokenEOF:				"EOF",
	TokenIgnored:			"Ignored",
	TokenLeftParen:			"(",
	TokenRightParen:		")",
	TokenSingleQuote:		"'",
	TokenDoubleQuote:		"\"",
	TokenDoubleQuotePair:	"\"\"",
	TokenSymbol:			"Symbol",
	TokenInt:				"Int",
}

// TokenInfo is type to wrap next token info
type TokenInfo struct {
	token		string
	tokenType	int
	lineNum		int
}

var regexSymbol = regexp.MustCompile(`^[^0-9()'"\t\n\v\f\r ][^()'"\t\n\v\f\r ]*`)
var regexIgnored = regexp.MustCompile(`^[\t\n\v\f\r ]+`)
var regexInt = regexp.MustCompile(`^[0-9]+`)
var regexString = regexp.MustCompile(`^[^"]+`)

// Lexer is the lexer type
type Lexer struct {
	sourceCode	string
	lineNum		int
	nextToken	*TokenInfo
}

// NewLexer is the constructor for lexer
func NewLexer(sourceCode string) *Lexer {
	return &Lexer{sourceCode, 0, nil}
}

func (lexer *Lexer) skipSourceCode(n int) {
	lexer.sourceCode = lexer.sourceCode[n:]
}

func (lexer *Lexer) nextSourceCodeIs(prefix string) bool {
	return strings.HasPrefix(lexer.sourceCode, prefix)
}

func (lexer *Lexer) scanPattern(pattern *regexp.Regexp) string {
	if matched := pattern.FindString(lexer.sourceCode); matched != "" {
		return matched
	}
	panic("unreachable!")
}

func (lexer *Lexer) scanSymbol() string {
	return lexer.scanPattern(regexSymbol)
}

func (lexer *Lexer) scanIgnored() string {
	return lexer.scanPattern(regexIgnored)
}

func (lexer *Lexer) scanInt() string {
	return lexer.scanPattern(regexInt)
}

func (lexer *Lexer) scanString() string {
	stringContent := lexer.scanPattern(regexString)
	lexer.skipSourceCode(len(stringContent))
	lexer.processNewLine(stringContent)
	return stringContent
}

func (lexer *Lexer) processNewLine(s string) {
	i := 0
	for i < len(s) {
		if len(s[i:]) > 1 && (s[i:][:2] == "\r\n" || s[i:][:2] == "\n\r") {
			lexer.lineNum++
			i += 2
		} else {
			if s[i] == '\r' || s[i] == '\n' {
				lexer.lineNum++
			}
			i++
		}
	}
}

// GetNextToken is method to get next token while parsing
func (lexer *Lexer) GetNextToken() TokenInfo {
	if lexer.nextToken != nil {
		nextToken := *(lexer.nextToken)
		lexer.nextToken = nil
		return nextToken
	}
	return lexer.MatchToken()
}

// MatchToken is method to match next token while parsing
func (lexer *Lexer) MatchToken() TokenInfo {
	if len(lexer.sourceCode) == 0 {
		return TokenInfo{"EOF", TokenEOF, lexer.lineNum}
	}
	switch lexer.sourceCode[0] {
	case '(':
		lexer.skipSourceCode(1)
		return TokenInfo{"(", TokenLeftParen, lexer.lineNum}
	case ')':
		lexer.skipSourceCode(1)
		return TokenInfo{")", TokenRightParen, lexer.lineNum}
	case '\'':
		lexer.skipSourceCode(1)
		return TokenInfo{"'", TokenSingleQuote, lexer.lineNum}
	case '"':
		if lexer.nextSourceCodeIs("\"\"") {
			lexer.skipSourceCode(2)
			return TokenInfo{"\"\"", TokenDoubleQuotePair, lexer.lineNum}
		}
		lexer.skipSourceCode(1)
		return TokenInfo{"\"", TokenDoubleQuote, lexer.lineNum}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': // Int
		matched := lexer.scanInt()
		lexer.skipSourceCode(len(matched))
		return TokenInfo{matched, TokenInt, lexer.lineNum}
	case '\t', '\n', '\v', '\f', '\r', ' ': // Ignored
		matched := lexer.scanIgnored()
		lineNum := lexer.lineNum
		lexer.skipSourceCode(len(matched))
		lexer.processNewLine(matched)
		return TokenInfo{matched, TokenIgnored, lineNum}
	default: // Symbol
		matched := lexer.scanSymbol()
		lexer.skipSourceCode(len(matched))
		return TokenInfo{matched, TokenSymbol, lexer.lineNum}
	}
}

// NextTokenIs is method to get next token given an inferred token type
func (lexer *Lexer) NextTokenIs(tokenType int) TokenInfo {
	fmt.Println(tokenNameMap[tokenType])
	nextToken := lexer.GetNextToken()
	if tokenType != nextToken.tokenType {
		err := fmt.Sprintf(
			"NextTokenIs(): syntax error near '%s': expecting '%s', got '%s'",
			nextToken.token,
			tokenNameMap[tokenType],
			nextToken.token,
		)
		panic(err)
	}
	return nextToken
}

// LookAhead is method to check type of next token withour consuming it
func (lexer *Lexer) LookAhead() int {
	if lexer.nextToken == nil {
		nextToken := lexer.GetNextToken()
		lexer.nextToken = &nextToken
	}
	return lexer.nextToken.tokenType
}
