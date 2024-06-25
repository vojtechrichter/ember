package main

import (
	"fmt"
	"os"
)

func ReadTemplate(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func TokenizeTemplate(template []byte) []string {
	for i := 0; i < len(template); i++ {
		switch template[i] {
		case '$':
			{
				fmt.Println("ID")
				// split
				i += 1
				id := make([]byte, 10)

				// split
				switch template[i] {
				case '/':
					{
						fmt.Println("CLOSE_ID")
					}
				}

				// split
				for ; i < len(template) && template[i] != '('; i++ {
					id = append(id, template[i])
				}

				switch string(id) {
				case "block":
					{
						fmt.Println("BLOCK")
					}
				}
			}
		}
	}

	return nil
}

func main() {
	file, err := ReadTemplate("template.em")
	if err != nil {
		panic(err)
	}

	TokenizeTemplate(file)
}
