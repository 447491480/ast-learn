package lib

type SimpleTokenReader struct {
	Tokens []Token
	Pos    int
}

func (reader *SimpleTokenReader) Read() *Token {
	if reader.Pos < len(reader.Tokens) {
		token := &reader.Tokens[reader.Pos]
		reader.Pos++
		return token
	}
	return nil
}

func (reader *SimpleTokenReader) Peek() *Token {
	if reader.Pos < len(reader.Tokens) {
		return &reader.Tokens[reader.Pos]
	}
	return nil
}

func (reader *SimpleTokenReader) Unread() {
	if reader.Pos > 0 {
		reader.Pos--
	}
}

func (reader *SimpleTokenReader) GetPosition() int {
	return reader.Pos
}

func (reader *SimpleTokenReader) SetPosition(position int) {
	if position >= 0 && position < len(reader.Tokens) {
		reader.Pos = position
	}
}