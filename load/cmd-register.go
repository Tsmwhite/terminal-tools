package load

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"terminal-tools/clause"
	"terminal-tools/config"
)

var cmdMap map[string]clause.Cmder
var cmdShortMap map[string]string
var cmdExamples map[string]string

func init() {
	cmdConfig := config.LoadCmdConfig()
	cmdExamples = cmdConfig.Examples
	cmdMap = make(map[string]clause.Cmder)
	for key, cmd := range cmdConfig.Commands {
		cmdMap[key] = cmd
	}
	cmdShortMap = cmdConfig.ShortKey
	//cmdMap = map[string]clause.Cmder{
	//	"git-commit-push":          new(git.Commit),
	//	"git-new-branch-by-master": new(git.Branch),
	//}
	//cmdShortMap = map[string]string{
	//	"gcp": "git-commit-push",
	//	"gcb": "git-new-branch-by-master",
	//}
}

func help() {
	fmt.Println("tools <command>")
	fmt.Println("Usage:")
	for key, cmd := range cmdMap {
		colorPrint(key)
		fmt.Println("example:", cmdExamples[key])
		fmt.Println("exec detail:")
		cmd.Print()
	}
	fmt.Println("--------quick-command--------")
	for short, cmdInfo := range cmdShortMap {
		colorPrint(short, " -> ", cmdInfo)
	}
}

func Handle() {
	args := os.Args
	if len(args) < 2 {
		help()
		return
	}
	key := args[1]
	if key == "help" {
		help()
		return
	}
	if cmdShortMap[key] != "" {
		key = cmdShortMap[key]
	}
	var err error
	if val, ok := cmdMap[key]; ok {
		if err = val.Handle(args[2:]); err != nil {
			colorPrint(err)
			return
		}
		err = val.Exec()
		if err != nil {
			colorPrint(err)
			return
		}
	}
}

func colorPrint(content ...interface{}) {
	_, err := color.New(color.FgRed).Println(content...)
	if err != nil {
		fmt.Println("error：", err)
	}
}
