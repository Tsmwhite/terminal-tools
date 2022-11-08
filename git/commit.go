package git

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

/***
 * git 提交后自动push
 * 命令顺序如下：
 * git add .
 * git commit -m"~"
 * git pull
 * git push
 */

type Commit struct {
	remark string
}

func (cm *Commit) Handle(args []string) error {
	if len(args) < 1 {
		return errors.New("请输入提交备注")
	}
	cm.remark = args[0]
	return nil
}

func (cm *Commit) Exec() error {
	commands := [][]string{
		{"git", "add", "."},
		{"git", "commit", "-m", cm.remark},
		{"git", "pull"},
		{"git", "push"},
	}
	for _, cmdStr := range commands {
		cmd := exec.Command(cmdStr[0], cmdStr[1:]...)
		outByte, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(cmdStr, " "), ":", string(outByte))
	}
	return nil
}
