package main

import (
	"os"
	"terminal-tools/cmd"
)

func main() {
	args := os.Args
	cmd.Handle(args)
}
