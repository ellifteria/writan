package main

import (
	"flag"
	"log"
	"os"

	"github.com/ellifteria/writan"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error

	inputFileNamePtr := flag.String("i", "input.wrtn", "path to input file")
	outputFileNamePtr := flag.String("o", "output.html", "path to output file")

	flag.Parse()

	plainWritan, err := os.ReadFile(*inputFileNamePtr)
	check(err)

	tokenizer := writan.MakeDefaultTokenizer()

	token := tokenizer.Tokenize(string(plainWritan))

	parser := writan.MakeParser()

	node := parser.Match(token)

	baseGenerator := writan.MakeGenerator()

	generatedHtml := baseGenerator.Generate(&node)

	outputFile, err := os.Create(*outputFileNamePtr)
	check(err)
	defer outputFile.Close()

	_, err = outputFile.WriteString(generatedHtml)
	check(err)
}
