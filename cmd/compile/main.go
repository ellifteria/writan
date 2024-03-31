package main

import (
	"github.com/ellifteria/writan"
)

func main() {
	tokenizer := writan.MakeDefaultTokenizer()
	tokenizer.Test("**Hello, World!** My name is *Elli*.")

	writan.TestTextParser()
}
