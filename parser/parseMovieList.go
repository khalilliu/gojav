package parser

import (
	"fmt"
	"gojav/engine"
)

func ParseMovieList (content [] byte) engine.ParseResult {
	fmt.Println(string(content))
	return engine.ParseResult{}
}
