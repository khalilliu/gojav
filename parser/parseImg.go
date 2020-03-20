package parser

import (
	"fmt"
	"github.com/fatih/color"
	"gojav/engine"
	"os"
)

func ParseImg(content []byte, path string) engine.ParseResult {
	// save content to path
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.Write(content)
	fmt.Println(color.GreenString("[图片写入完成]:"), path)
	return engine.ParseResult{}
}