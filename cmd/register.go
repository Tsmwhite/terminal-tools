package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"terminal-tools/git"
)

var cmdMap map[string]Cmder
var cmdShortMap map[string]string

func init() {
	cmdMap = map[string]Cmder{
		"git-commit-push": new(git.Commit),
	}
	cmdShortMap = map[string]string{
		"gcp": "git-commit-push",
	}
}

func Handle() {
	args := os.Args
	if len(args) < 2 {
		printlnErr("what are you doing")
		return
	}
	key := args[1]
	if cmdShortMap[key] != "" {
		key = cmdShortMap[key]
	}
	var err error
	if val, ok := cmdMap[key]; ok {
		if err = val.Handle(args[2:]); err != nil {
			printlnErr(err)
			return
		}
		err = val.Exec()
		if err != nil {
			printlnErr(err)
			return
		}
	}
}

func printlnErr(content interface{}) {
	_, err := color.New(color.FgRed).Println(content)
	if err != nil {
		fmt.Println("error：", err)
	}
}
