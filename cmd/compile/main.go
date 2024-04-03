package main

import (
	"github.com/ellifteria/writan"
)

func main() {
	tokenizer := writan.MakeDefaultTokenizer()
	// tokenizer.Test("@*Hello, World!@* My name is @_Elli@_.")
	// tokenizer.Test("@*@_Hello, World!@_@* My name is @_Elli@_.")
	// tokenizer.Test("@*@_Hello@_, World!@* My name is @_Elli@_.")
	// 	tokenizer.Test("@/hello@/")
	// 	plainMarkdown := `@*@_Hello, World!@/@/
	// My name is Elli, Beres. I'm a computer science student at Northwestern University.

	// Find me at @celli.beres@@u.northwestern.edu@/`

	// tokenizer.Test(plainMarkdown)

	// tokenizer.TestBTP("hello@*h@*")
	// tokenizer.TestBTP("@*hello@*")
	// tokenizer.TestBTP("@*hello@*, world")
	tokenizer.TestBTP("@*@_h@_ello@*, world")
}
