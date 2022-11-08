package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

type Cmder interface {
	Handle(args []string) error
	Exec() error
}

type Command struct {
	Commands [][]string
}

func (cmd *Command) Exec() error  {
	return CommandsRun(cmd.Commands)
}

func CommandsRun(commands [][]string) error {
	for _, cmdStrings := range commands {
		if err := CommandRun(cmdStrings); err != nil {
			return err
		}
	}
	return nil
}

func CommandRun(cmdStrings []string) error {
	cmd := exec.Command(cmdStrings[0], cmdStrings[1:]...)
	outByte, err := cmd.Output()
	if err != nil {
		return err
	}
	cmdStr := strings.Join(cmdStrings, " ")
	if len(outByte) > 0 {
		fmt.Println(cmdStr, "\n", string(outByte))
	} else {
		fmt.Println(cmdStr)
	}
	return nil
}
