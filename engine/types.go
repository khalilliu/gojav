package engine

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
