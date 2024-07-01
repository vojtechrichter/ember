package lex

import "os"

type Lexer struct {
	TemplateName    string
	TemplateContent []byte
	TemplateSize    int

	Idx         int
	CurrentChar byte
}

func LexerInit(templateName string) (*Lexer, error) {
	lexer := new(Lexer)
	lexer.TemplateName = templateName

	content, err := os.ReadFile(templateName)
	if err != nil {
		return lexer, err
	}
	lexer.TemplateContent = content
	lexer.TemplateSize = len(content)

	lexer.Idx = 0
	lexer.CurrentChar = lexer.TemplateContent[lexer.Idx]

	return lexer, nil
}

func (l *Lexer) Advance() {
	l.Idx += 1
	l.CurrentChar = l.TemplateContent[l.Idx]
}

func (l *Lexer) AdvanceBy(n int) {
	l.Idx += n
	l.CurrentChar = l.TemplateContent[l.Idx]
}

func (l *Lexer) PeekBy(pos int) byte {
	return l.TemplateContent[l.Idx+pos]
}
