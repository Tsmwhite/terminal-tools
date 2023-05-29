package config

import (
	"fmt"
	"strings"
	"testing"
)

var cmdStrings = map[string][]string{
	"git add .":                                       {},
	"git commit -m${remark}":                          {"${remark}"},
	"git checkout -b ${newBranchName}":                {"${newBranchName}"},
	"git push --set-upstream origin ${newBranchName}": {"${newBranchName}"},
	"ls ${option} | grep ${search}":                   {"${option}", "${search}"},
}

func TestSearchArg(t *testing.T) {
	for cmdStr, expected := range cmdStrings {
		fmt.Println("cmdStr:", cmdStr)
		response := searchArg(cmdStr)
		fmt.Println("response:", response)
		ok := strings.Join(response, ",") == strings.Join(expected, ",")
		fmt.Println("match:", ok)
		if !ok {
			t.Fail()
		}
	}
}
