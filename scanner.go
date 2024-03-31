package writan

type Scanner interface {
	fromString(plainMarkdown string) Token
}
