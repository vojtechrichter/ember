package lex

import "os"

type Lexer struct {
	templateName    string
	templateContent []byte
	templateSize    int

	idx         uintptr
	currentChar byte
}

func LexerInit(templateName string) (*Lexer, error) {
	lexer := new(Lexer)
	lexer.templateName = templateName

	content, err := os.ReadFile(templateName)
	if err != nil {
		return lexer, err
	}
	lexer.templateContent = content
	lexer.templateSize = len(content)

	lexer.idx = 0
	lexer.currentChar = lexer.templateContent[lexer.idx]

	return lexer, nil
}

func (l *Lexer) Advance() {
	l.idx += 1
	l.currentChar = l.templateContent[l.idx]
}

func (l *Lexer) AdvanceBy(n uintptr) {
	l.idx += n
	l.currentChar = l.templateContent[l.idx]
}

func (l *Lexer) Peek(pos uintptr) byte {
	return l.templateContent[pos]
}
