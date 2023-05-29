package clause

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type Cmder interface {
	Handle() error
	Exec() error
	SetArgsVal([]string)
	Print()
}

type Command struct {
	Commands [][]string
	ArgKeys  []string
	ArgsMap  map[string]string
}

func (cmd *Command) Handle() error {
	for _, argKey := range cmd.ArgKeys {
		if cmd.ArgsMap[argKey] == "" {
			cmd.Print()
			return errors.New("缺少必要参数-" + argKey)
		}
		for index, cmdArr := range cmd.Commands {
			cmdStr := strings.Join(cmdArr, " ")
			cmdStr = strings.ReplaceAll(cmdStr, argKey, cmd.ArgsMap[argKey])
			cmd.Commands[index] = strings.Split(cmdStr, " ")
		}
	}
	return nil
}

func (cmd *Command) Exec() error {
	return CommandsRun(cmd.Commands)
}

func (cmd *Command) SetArgsVal(values []string) {
	cmd.ArgsMap = make(map[string]string)
	for _, key := range cmd.ArgKeys {
		if len(values) > 0 {
			cmd.ArgsMap[key] = values[0]
			values = values[1:]
		} else {
			return
		}
	}
}

func (cmd *Command) Print() {
	for _, cmdArr := range cmd.Commands {
		fmt.Println("  ", strings.Join(cmdArr, " "))
	}
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
