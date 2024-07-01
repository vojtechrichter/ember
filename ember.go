package main

import (
	"fmt"
	"github.com/vojtechrichter/ember/lex"
	"log"
)

type Token struct {
	name string
}

func TokenizeTemplate(lex *lex.Lexer) map[Token]byte {
	tokens := make(map[Token]byte)

	openTag := false
	lastOpenTag := make([]byte, 1<<4)

	for ; lex.Idx < lex.TemplateSize-1; lex.Advance() {
		switch lex.CurrentChar {
		case '$':
			{
				if !openTag {
					lex.Advance()
					if len(lastOpenTag) > 0 || cap(lastOpenTag) > 0 {
						lastOpenTag = nil
					}
					for ; lex.Idx < lex.TemplateSize && lex.CurrentChar != '('; lex.Advance() {
						if lex.PeekBy(1) == '(' {
							openTag = true
						}
						lastOpenTag = append(lastOpenTag, lex.CurrentChar)
					}
					tokenName := "OPEN_TAG{" + string(lastOpenTag) + "}"
					tokens[Token{name: tokenName}] = lex.CurrentChar
				}

				lex.Advance()

				switch lex.CurrentChar {
				case '/':
					{
						tokenName := "CLOSE_TAG{" + string(lastOpenTag) + "}"
						tokens[Token{name: tokenName}] = lex.CurrentChar
						openTag = false
					}
				}
			}
		default:
			{
				tokens[Token{name: "HTML_CONTENT"}] = lex.CurrentChar
			}
		}
	}

	return tokens
}

func main() {
	lexer, err := lex.LexerInit("template.eb")
	if err != nil {
		log.Fatal(err)
	}

	tokens := TokenizeTemplate(lexer)

	for k, v := range tokens {
		fmt.Printf("token: %s, value: %s\n", k.name, string(v))
	}
}
