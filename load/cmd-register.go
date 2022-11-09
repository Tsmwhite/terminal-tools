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

func init() {
	cmdConfig := config.LoadCmdConfig()
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

func Handle() {
	args := os.Args
	if len(args) < 2 {
		println("what are you doing")
		return
	}
	key := args[1]
	if cmdShortMap[key] != "" {
		key = cmdShortMap[key]
	}
	var err error
	if val, ok := cmdMap[key]; ok {
		if err = val.Handle(args[2:]); err != nil {
			println(err)
			return
		}
		err = val.Exec()
		if err != nil {
			println(err)
			return
		}
	}
}

func println(content interface{}) {
	_, err := color.New(color.FgRed).Println(content)
	if err != nil {
		fmt.Println("error：", err)
	}
}
