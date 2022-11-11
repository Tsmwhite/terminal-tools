package clause

import (
	"fmt"
	"testing"
)

func TestCommand_Handle(t *testing.T) {
	tsList := []*Command{
		{
			Commands: [][]string{
				{"git", "add", "."},
				{"git", "commit", "-m", "${remark}"},
				{"git", "pull"},
				{"git", "push"},
			},
			ArgKeys: []string{"${remark}"},
			ArgsMap: map[string]string{
				"${remark}": "First Commit By Test",
			},
		},
		{
			Commands: [][]string{
				{"git", "checkout", "master"},
				{"git", "pull"},
				{"git", "checkout", "-b", "${name}"},
				{"git", "push", "--set-upstream", "origin", "${name}"},
			},
			ArgKeys: []string{"${name}"},
			ArgsMap: map[string]string{
				"${name}": "hotfix/robin_testBranchName_20221109",
			},
		},
		{
			Commands: [][]string{
				{"ls", "${option}", "|", "grep", "${search}"},
			},
			ArgKeys: []string{"${option}", "${search}"},
			ArgsMap: map[string]string{
				"${option}": "-a",
				"${search}": "config.yaml",
			},
		},
	}
	for _, cmd := range tsList {
		if err := cmd.Handle(); err != nil {
			t.Fail()
		} else {
			fmt.Println(cmd)
		}
	}
}
