package gokenizer

import (
	"fmt"
	"os"
	"path/filepath"
	"unicode/utf8"
)

func Tokenize(content string) []string {
	var tokens []string
	var currentToken string
	inQuotes := false
	quoteChar := ' '
	escapeNext := false

	for i := 0; i < len(content); {
		char, size := utf8.DecodeRuneInString(content[i:])
		if char == '\n' {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			i += size
			continue
		}
		if escapeNext {
			currentToken += string(char)
			escapeNext = false
			i += size
			continue
		}

		if inQuotes {
			if char == '\\' {
				escapeNext = true
				currentToken += string(char)
			} else {
				currentToken += string(char)
				if char == quoteChar {
					inQuotes = false
					if currentToken != "" {
						tokens = append(tokens, currentToken)
						currentToken = ""
					}
				}
			}
		} else {
			switch char {
			case '/':
				if i+1 < len(content) && content[i+1] == '/' {
					i += 2
					for content[i] != '\n' {
						i++
					}
					continue
				} else if i+1 < len(content) && content[i+1] == '*' {
					i += 2
					for i < len(content) && !(content[i] == '*' && i+1 < len(content) && content[i+1] == '/') {
						i++
					}
					i += 2
					continue
				} else if i+1 < len(content) && content[i+1] == '=' {
					tokens = append(tokens, "/=")
					i += 2

				}

			case ' ', '(', ')', '{', '}', '[', ']', ',', '\t':
				if currentToken != "" {
					tokens = append(tokens, currentToken)
					currentToken = ""
				}
				if char != ' ' && char != '\t' {
					tokens = append(tokens, string(char))
				}
			case '"', '\'':

				if currentToken != "" {
					tokens = append(tokens, currentToken)
					currentToken = ""
				}
				inQuotes = true
				quoteChar = char
				currentToken += string(char)
			case '!':
				if i+1 < len(content) && content[i+1] == '=' {
					tokens = append(tokens, "!=")
					i++ // Move past the '=' character
				} else {
					tokens = append(tokens, "!")
				}
			case '=':
				if currentToken == "!" {
					tokens = append(tokens, "!=")
					currentToken = ""
				} else {
					currentToken += string(char)
				}
			default:
				currentToken += string(char)
			}
		}
		i += size

	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	//TODO: dont iterate over array again, instead dont append them in the first place

	var filteredTokens []string
	for _, token := range tokens {
		if token != "" && token != " " && token != "\n" && token != "\t" && token != "\r" {
			filteredTokens = append(filteredTokens, token)
		}
	}

	return filteredTokens
}

func CheckAndGetValidFile(inputFile string) []byte {
	if inputFile == "" {
		fmt.Println("Please provide an input file using the -file flag.")
		os.Exit(1)
	}

	file, err := os.ReadFile(inputFile)

	if err != nil {
		fmt.Printf("File '%s' does not exist\n", inputFile)
		os.Exit(1)
	}

	if !checkFileType(inputFile, ".go") {
		fmt.Printf("File '%s' is not a .go file\n", inputFile)
		os.Exit(1)
	}

	return file

}

func checkFileType(path string, filetype string) bool {
	readFileType := filepath.Ext(path)
	return readFileType == filetype
}
