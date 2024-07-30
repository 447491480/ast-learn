package lib

import (
	"strings"
	"unicode"
)

type SimpleLexer struct {
	tokenText strings.Builder
	tokens    []Token
	token     Token
}

func (lexer *SimpleLexer) IsAlpha(ch rune) bool {
	return unicode.IsLetter(ch)
}

func (lexer *SimpleLexer) IsDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (lexer *SimpleLexer) isBlank(ch rune) bool {
	return unicode.IsSpace(ch)
}

func (lexer *SimpleLexer) InitToken(ch rune) DfaState {
	if lexer.tokenText.Len() > 0 {
		lexer.token.Text = lexer.tokenText.String()
		lexer.tokens = append(lexer.tokens, lexer.token)
		lexer.tokenText.Reset()
		lexer.token = Token{}
	}

	var newState DfaState
	switch {
	case lexer.IsAlpha(ch):
		if ch == 'i' {
			newState = IdInt1
		} else {
			newState = Id
		}
		lexer.token.Type = Identifier
		lexer.tokenText.WriteRune(ch)
	case lexer.IsDigit(ch):
		newState = IntLiteralState
		lexer.token.Type = IntLiteral
		lexer.tokenText.WriteRune(ch)
	case ch == '>':
		newState = GTState
		lexer.token.Type = GT
		lexer.tokenText.WriteRune(ch)
	case ch == '+':
		newState = PlusState
		lexer.token.Type = Plus
		lexer.tokenText.WriteRune(ch)
	case ch == '-':
		newState = MinusState
		lexer.token.Type = Minus
		lexer.tokenText.WriteRune(ch)
	case ch == '*':
		newState = StarState
		lexer.token.Type = Star
		lexer.tokenText.WriteRune(ch)
	case ch == '/':
		newState = SlashState
		lexer.token.Type = Slash
		lexer.tokenText.WriteRune(ch)
	case ch == ';':
		newState = SemiColonState
		lexer.token.Type = SemiColon
		lexer.tokenText.WriteRune(ch)
	case ch == '(':
		newState = LeftParenState
		lexer.token.Type = LeftParen
		lexer.tokenText.WriteRune(ch)
	case ch == ')':
		newState = RightParenState
		lexer.token.Type = RightParen
		lexer.tokenText.WriteRune(ch)
	case ch == '=':
		newState = AssignmentState
		lexer.token.Type = Assignment
		lexer.tokenText.WriteRune(ch)
	default:
		newState = Initial
	}
	return newState
}

func (lexer *SimpleLexer) Tokenize(code string) []Token {
	lexer.tokens = []Token{}
	lexer.tokenText.Reset()
	lexer.token = Token{}
	state := Initial

	for _, ch := range code {
		switch state {
		case Initial:
			state = lexer.InitToken(ch)
		case Id:
			if lexer.IsAlpha(ch) || lexer.IsDigit(ch) {
				lexer.tokenText.WriteRune(ch)
			} else {
				state = lexer.InitToken(ch)
			}
		case GTState:
			if ch == '=' {
				lexer.token.Type = GE
				state = GEState
				lexer.tokenText.WriteRune(ch)
			} else {
				state = lexer.InitToken(ch)
			}
		case GEState, AssignmentState, PlusState, MinusState, StarState, SlashState, SemiColonState, LeftParenState, RightParenState:
			state = lexer.InitToken(ch)
		case IntLiteralState:
			if lexer.IsDigit(ch) {
				lexer.tokenText.WriteRune(ch)
			} else {
				state = lexer.InitToken(ch)
			}
		case IdInt1:
			if ch == 'n' {
				state = IdInt2
				lexer.tokenText.WriteRune(ch)
			} else if lexer.IsAlpha(ch) || lexer.IsDigit(ch) {
				state = Id
				lexer.tokenText.WriteRune(ch)
			} else {
				state = lexer.InitToken(ch)
			}
		case IdInt2:
			if ch == 't' {
				state = IdInt3
				lexer.tokenText.WriteRune(ch)
			} else if lexer.IsAlpha(ch) || lexer.IsDigit(ch) {
				state = Id
				lexer.tokenText.WriteRune(ch)
			} else {
				state = lexer.InitToken(ch)
			}
		case IdInt3:
			if lexer.isBlank(ch) {
				lexer.token.Type = Int
				state = lexer.InitToken(ch)
			} else {
				state = Id
				lexer.tokenText.WriteRune(ch)
			}
		}
	}

	if lexer.tokenText.Len() > 0 {
		lexer.InitToken(' ')
	}

	return lexer.tokens
}