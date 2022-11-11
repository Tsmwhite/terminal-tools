package git

import (
	"os/exec"
	"strings"
	"terminal-tools/clause"
)

/***
 * git 合并最新master
 * 命令顺序如下：
 * git pull
 * git checkout master
 * git pull
 * git checkout ~
 * git merge master
 */

type mergeMaster struct {
	clause.Command
}

func NewMergeMaster() *mergeMaster {
	return &mergeMaster{
		clause.Command{
			Commands: [][]string{
				{"git", "pull"},
				{"git", "checkout", "master"},
				{"git", "pull"},
				{"git", "checkout", "${branchName}"},
				{"git", "merge", "master"},
			},
			ArgKeys: []string{
				"${branchName}",
			},
		},
	}
}

func (m *mergeMaster) Handle() error {
	cmd := exec.Command("git", "status")
	outByte, err := cmd.Output()
	if err != nil {
		return err
	}
	content := string(outByte)
	firstLine := strings.Split(content, "\n")[0]
	firstLine = strings.ReplaceAll(firstLine, "On branch", "")
	branchName := strings.Trim(firstLine, " ")
	m.ArgsMap["${branchName}"] = branchName
	return m.Command.Handle()
}
