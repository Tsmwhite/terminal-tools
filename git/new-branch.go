package git

import (
	"errors"
	"terminal-tools/cmd"
)

/***
 * git 快速切一个基于master的分支
 * 命令顺序如下：
 * git checkout master
 * git pull
 * git checkout -b "~"
 * git push --set-upstream origin "~"
 */

type Branch struct {
	cmd.Command
	name string
}

func (bh *Branch) Handle(args []string) error {
	if len(args) < 1 {
		return errors.New("请输入分支名")
	}
	bh.name = args[0]
	bh.Commands = [][]string{
		{"git", "checkout", "master"},
		{"git", "pull"},
		{"git", "checkout", "-b", bh.name},
		{"git", "push", "--set-upstream", "origin", bh.name},
	}
	return nil
}
