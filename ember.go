package main

import (
	"fmt"
	"github.com/vojtechrichter/ember/lex"
	"log"
)

type Token struct {
	name string
}

func TokenizeTemplate(lex *lex.Lexer) []Token {
	tokens := make([]Token, 0)

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
					tokens = append(tokens, Token{
						name: tokenName,
					})
				}

				lex.Advance()

				switch lex.CurrentChar {
				case '/':
					{
						tokenName := "CLOSE_TAG{" + string(lastOpenTag) + "}"
						tokens = append(tokens, Token{
							name: tokenName,
						})
						openTag = false
					}
				}
			}
			//default:
			//	{
			//		tokens = append(tokens, Token{
			//			name: "HTML_CONTENT",
			//		})
			//	}
		}
	}

	return tokens
}

func main() {
	lexer, err := lex.LexerInit("template.em")
	if err != nil {
		log.Fatal(err)
	}

	tokens := TokenizeTemplate(lexer)

	for _, v := range tokens {
		fmt.Println(v.name)
	}
}
