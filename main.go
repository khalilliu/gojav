package main

import (
	"gojav/cmd"
	"gojav/engine"
)


func main() {
	cmd.Setup()
	engine.Execute()
}