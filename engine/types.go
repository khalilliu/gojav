package engine

import "gojav/model"

const (
	IMG = iota
	HTML
)

var (
	//startUrl       string
	TargetHasFound = false
	//CurPage        = 1
	//end            = false
	//total          = 0
)


type Item struct {
	model.Movie
	//Url     string
	//Type    int
	//Path    string
	//Payload interface{}
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Request struct {
	Url       string
	Type      int
	ParseFunc func([]byte) ParseResult
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
