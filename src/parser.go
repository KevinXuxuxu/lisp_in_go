package lispingo

import (
	"strconv"
)

func (lexer *Lexer) parseIgnored() {
	if lexer.LookAhead() == TokenIgnored {
		lexer.NextTokenIs(TokenIgnored)
	}
}

// ParseFunctionCall is method to parse function call from source code
func (lexer *Lexer) ParseFunctionCall() *FunctionCall {
	lineNum := lexer.NextTokenIs(TokenLeftParen).lineNum
	lexer.parseIgnored()
	function := lexer.parseElement()
	arguments := lexer.parseElements()
	lexer.NextTokenIs(TokenRightParen)
	lexer.parseIgnored()
	return &FunctionCall{function, arguments, lineNum}
}

func (lexer *Lexer) parseElement() Element {
	switch lexer.LookAhead() {
	case TokenInt:
		return lexer.parseInt()
	case TokenDoubleQuote:
		return lexer.parseString()
	case TokenSingleQuote:
		return lexer.parseList()
	case TokenLeftParen:
		return lexer.ParseFunctionCall()
	default:
		return lexer.parseSymbol()
	}
}

func (lexer *Lexer) parseElements() []Element {
	elements := []Element{}
	for lexer.LookAhead() != TokenRightParen {
		elements = append(elements, lexer.parseElement())
	}
	return elements
}

func (lexer *Lexer) parseInt() *Int {
	intToken := lexer.NextTokenIs(TokenInt)
	lexer.parseIgnored()
	value, _ := strconv.Atoi(intToken.token)
	return &Int{value, intToken.lineNum}
}

func (lexer *Lexer) parseString() *String {
	if lexer.LookAhead() == TokenDoubleQuotePair {
		empty := lexer.NextTokenIs(TokenDoubleQuotePair)
		return &String{"", empty.lineNum}
	}
	lineNum := lexer.NextTokenIs(TokenDoubleQuote).lineNum
	stringContent := lexer.scanString()
	lexer.NextTokenIs(TokenDoubleQuote)
	lexer.parseIgnored()
	return &String{stringContent, lineNum}
}

func (lexer *Lexer) parseSymbol() *Symbol {
	symbolToken := lexer.NextTokenIs(TokenSymbol)
	lexer.parseIgnored()
	return &Symbol{symbolToken.token, symbolToken.lineNum}
}

func (lexer *Lexer) parseList() *List {
	lineNum := lexer.NextTokenIs(TokenSingleQuote).lineNum
	lexer.NextTokenIs(TokenLeftParen)
	lexer.parseIgnored()
	elements := lexer.parseElements()
	lexer.NextTokenIs(TokenRightParen)
	lexer.parseIgnored()
	return &List{elements, lineNum}
}
