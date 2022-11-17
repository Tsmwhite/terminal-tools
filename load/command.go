package load

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"terminal-tools/clause"
	"terminal-tools/config"
	"terminal-tools/tools/git"
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
	cmdMap["git-merge-master"] = git.NewMergeMaster()
	cmdShortMap["gmm"] = "git-merge-master"
}

func Handle() {
	args := os.Args
	if len(args) < 2 || args[1] == "help" {
		help()
		return
	}
	key := args[1]
	if cmdShortMap[key] != "" {
		key = cmdShortMap[key]
	}
	var err error
	if val, ok := cmdMap[key]; ok {
		val.SetArgsVal(args[2:])
		if err = val.Handle(); err != nil {
			colorPrint(err)
			return
		}
		err = val.Exec()
		if err != nil {
			colorPrint(err)
			return
		}
	} else {
		colorPrint("'"+key+"' is not command. See help")
	}
}

func colorPrint(content ...interface{}) {
	_, err := color.New(color.FgRed).Println(content...)
	if err != nil {
		fmt.Println("error：", err)
	}
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
