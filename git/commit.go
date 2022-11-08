package git

import (
	"errors"
	"terminal-tools/cmd"
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
	cmd.Command
	remark string
}

func (cm *Commit) Handle(args []string) error {
	if len(args) < 1 {
		return errors.New("请输入提交备注")
	}
	cm.remark = args[0]
	cm.Commands = [][]string{
		{"git", "add", "."},
		{"git", "commit", "-m", cm.remark},
		{"git", "pull"},
		{"git", "push"},
	}
	return nil
}
