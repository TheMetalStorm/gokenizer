package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/TheMetalStorm/gokenizer"
)

func main() {

	var inputFile string

	flag.StringVar(&inputFile, "file", "", "Input file path of .go file")
	flag.Parse()

	file := gokenizer.CheckAndGetValidFile(inputFile)

	tokens := gokenizer.Tokenize(string(file))
	for _, token := range tokens {
		trimmed := strings.TrimSpace(token)
		if trimmed == "" {
			continue
		}
		fmt.Printf("token: %v\n", trimmed) //this is also a comment
	}

}
