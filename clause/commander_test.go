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
			ArgsMap: map[string]string{
				"${remark}": "",
			},
		},
		{
			Commands: [][]string{
				{"git", "checkout", "master"},
				{"git", "pull"},
				{"git", "checkout", "-b", "${name}"},
				{"git", "push", "--set-upstream", "origin", "${name}"},
			},
			ArgsMap: map[string]string{
				"${name}": "",
			},
		},
		{
			Commands: [][]string{
				{"ls", "${option}", "|", "grep", "${search}"},
			},
			ArgsMap: map[string]string{
				"${option}": "",
				"${search}": "",
			},
		},
	}
	testArgs := [][]string{
		{"First Commit By Test"},
		{"hotfix/robin_testBranchName_20221109"},
		{"-a", "config.yaml"},
	}
	for index, cmd := range tsList {
		if err := cmd.Handle(testArgs[index]); err != nil {
			t.Fail()
		} else {
			fmt.Println(cmd)
		}
	}
}
