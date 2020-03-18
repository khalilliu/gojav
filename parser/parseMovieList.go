package parser

import (
	"fmt"
	"gojav/engine"
)

func ParseMovieList (content [] byte) engine.ParseResult {
	fmt.Println(content)
	return engine.ParseResult{}
}
