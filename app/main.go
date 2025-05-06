package main

import (
	"flag"
	"fmt"

	"github.com/TheMetalStorm/gokenizer"
)

func main() {

	var inputFilePath string

	flag.StringVar(&inputFilePath, "file", "", "Input file path of .go file")
	flag.Parse()

	file := gokenizer.CheckAndGetValidFile(inputFilePath)

	tokens := gokenizer.Tokenize(string(file))
	for _, token := range tokens {
		fmt.Printf("token: %s\n", token)
	}

}
