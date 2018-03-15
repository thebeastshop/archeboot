package data

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type FType int

const (
	FTypeDir FType = iota
	FTypeXML
	FTypeJSON
	FTypeText
)

type ExecCommand struct {
	Command   string
	Arguments *[]string
}

type File struct {
	Name    string
	AbsPath string
	Type    FType
}

type Template struct {
	Name         string
	Arguments    *[]string
	Files        *[]File
	ExecCommands *[]ExecCommand
}

type TemplateYaml struct {
	Name         string         `yaml:"name"`
	Arguments    *[]string      `yaml:"name"`
	Files        *[]interface{} `yaml:"files"`
	ExecCommands *[]string      `yaml:"exec"`
}

func ReadTemplateYaml() (*TemplateYaml, error) {
	return nil, nil
}
