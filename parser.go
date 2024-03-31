package writan

import "fmt"

type Parser interface {
	match(token Token) Node
}

func TestTextParser() {
	parser := TextParser{}
	fmt.Println(parser.match(makeToken(TEXT_TOKEN, "hello")))
	fmt.Println(parser.match(makeToken(TEXT_TOKEN, "hello, world")))
	fmt.Println(parser.match(makeToken(UNDERSCORE_TOKEN, "_")))
}
