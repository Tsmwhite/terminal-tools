package clause

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type Cmder interface {
	Handle(args []string) error
	Exec() error
	Print()
}

type Command struct {
	Commands [][]string
	ArgsMap  map[string]string
}

func (cmd *Command) Print() {
	for _, cmdArr := range cmd.Commands {
		fmt.Println("  ",strings.Join(cmdArr, " "))
	}
}

func (cmd *Command) Handle(args []string) error {
	if len(cmd.ArgsMap) > 0 {
		for argKey, _ := range cmd.ArgsMap {
			if len(args) < 1 {
				cmd.Print()
				return errors.New("缺少必要参数-" + argKey)
			}
			cmd.ArgsMap[argKey] = args[0]
			args = args[1:]
			for index, cmdArr := range cmd.Commands {
				cmdStr := strings.Join(cmdArr, " ")
				cmdStr = strings.ReplaceAll(cmdStr, argKey, cmd.ArgsMap[argKey])
				cmd.Commands[index] = strings.Split(cmdStr, " ")
			}
		}
	}
	return nil
}

func (cmd *Command) Exec() error {
	return CommandsRun(cmd.Commands)
}

func CommandsRun(commands [][]string) error {
	for _, cmdStrings := range commands {
		if err := CommandRun(cmdStrings); err != nil {
			return err
		}
	}
	fmt.Println("----------------OK----------------")
	return nil
}

func CommandRun(cmdStrings []string) error {
	fmt.Println(strings.Join(cmdStrings, " "))
	cmd := exec.Command(cmdStrings[0], cmdStrings[1:]...)
	outByte, err := cmd.Output()
	if err != nil {
		return err
	}
	if len(outByte) > 0 {
		fmt.Println(string(outByte))
	}
	return nil
}
