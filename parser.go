package writan

type Parser interface {
	match(token Token) Node
}
