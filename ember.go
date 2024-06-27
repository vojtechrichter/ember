package main

import (
	"fmt"
	"os"
)

func ReadTemplate(name string) ([]byte, error) {
	return os.ReadFile(name)
}

const (
	TOKEN_TYPE_SINGLE = iota
	TOKEN_TYPE_DOUBLE = iota
)

type Token struct {
	name string
}

func TokenizeTemplate(template []byte) []Token {
	tokens := make([]Token, 0)

	openTag := false
	lastOpTag := make([]byte, 1<<4)
	for i := 0; i < len(template); i++ {
		switch template[i] {
		case '$':
			{
				if !openTag {
					i += 1
					if len(lastOpTag) > 0 || cap(lastOpTag) > 0 {
						lastOpTag = nil
					}
					for ; i < len(template) && template[i] != '('; i++ {
						if template[i+1] == '(' {
							openTag = true
						}
						lastOpTag = append(lastOpTag, template[i])
					}
					tokenName := "OPEN_TAG{" + string(lastOpTag) + "}"
					tokens = append(tokens, Token{
						name: tokenName,
					})
				}

				i += 1

				switch template[i] {
				case '/':
					{
						tokenName := "CLOSE_TAG{" + string(lastOpTag) + "}"
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
	file, err := ReadTemplate("template.em")
	if err != nil {
		panic(err)
	}

	tokens := TokenizeTemplate(file)
	for _, v := range tokens {
		fmt.Println(v.name)
	}
}
