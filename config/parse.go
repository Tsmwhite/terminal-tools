package config

import (
	"errors"
	"fmt"
	ym2 "gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
	"strings"
	"terminal-tools/clause"
)

const cmdConfigFile = "/usr/local/var/Go/terminal-tools/config/commands.yaml"
const searchCmdArgPattern = `\${\w*}`

func loadConfig(configFile string, dest interface{}) error {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return errors.New("loadConfig ReadFile Error:" + err.Error())
	}
	if err = ym2.Unmarshal(content, dest); err != nil {
		return errors.New("loadConfig Unmarshal Error:" + err.Error())
	}
	return nil
}

type CommandConfig struct {
	ShortKey  map[string]string   `yaml:"short-key"`
	CmdSource map[string][]string `yaml:"cmds"`
	Examples  map[string]string   `yaml:"examples"`
	Commands  map[string]*clause.Command
}

func LoadCmdConfig() *CommandConfig {
	configData := new(CommandConfig)
	if err := loadConfig(cmdConfigFile, configData); err != nil {
		fmt.Println(err)
	}
	return configData.parse()
}

func (cf *CommandConfig) parse() *CommandConfig {
	cf.Commands = make(map[string]*clause.Command)
	for cmdKey, cmdStrings := range cf.CmdSource {
		var commands [][]string
		var needArgs []string
		for _, cmdStr := range cmdStrings {
			res := searchArg(cmdStr)
			if len(res) > 0 {
				needArgs = append(needArgs, res...)
			}
			commands = append(commands, strings.Split(cmdStr, " "))
		}
		argsMap := make(map[string]string)
		var uniqueArgs []string
		for _, arg := range needArgs {
			if argsMap[arg] == "" {
				argsMap[arg] = "ok"
				uniqueArgs = append(uniqueArgs, arg)
			}
		}
		cf.Commands[cmdKey] = &clause.Command{
			Commands: commands,
			ArgKeys: uniqueArgs,
		}
	}
	return cf
}

func searchArg(cmdStr string) []string {
	var args []string
	ok, _ := regexp.MatchString(searchCmdArgPattern, cmdStr)
	if ok {
		reg := regexp.MustCompile(searchCmdArgPattern)
		argsBytes := reg.FindAll([]byte(cmdStr), -1)
		for _, argByte := range argsBytes {
			args = append(args, string(argByte))
		}
	}
	return args
}
