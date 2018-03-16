package data

import (
"archeboot/utils"
"gopkg.in/yaml.v2"
"io/ioutil"
"fmt"
)

type FType int

const (
	FTypeDir FType = iota
	FTypeXML
	FTypeJSON
	FTypeText
	FTypeJava
	FTypeJavaPackage
)

type ExecCommand struct {
	Command   string
	Arguments *[]string
}

type Creation struct {
	Name    string
	AbsPath string
	Type    FType
}

type Template struct {
	Name         string
	Arguments    *[]string
	Creations    *[]Creation
	ExecCommands *[]ExecCommand
}

type TemplateYaml struct {
	Name         string         `yaml:"name"`
	Arguments    *[]string      `yaml:"arguments"`
	Resources    *[]string      `yaml:"resources"`
	Creations    *[]interface{} `yaml:"create"`
	ExecCommands *[]string      `yaml:"exec"`
}

func ReadTemplateYaml(config *Config) (*TemplateYaml, error) {
	tempYaml := &TemplateYaml{}
	yamlFile, err := ioutil.ReadFile(config.Info.DefaultBootFileName)
	if err != nil {
		utils.LogError(err)
	}
	err = yaml.Unmarshal(yamlFile, tempYaml)
	if err != nil {
		utils.LogError(err)
	}
	return tempYaml, err
}


func parseCreation(ct interface{}) *Creation {
	switch ct.(type) {
	case string:
		fmt.Println("Creation String:", ct)
		return handleStringCreation(ct.(string))
	case []interface{}:
		fmt.Println("Creation list:")
		for i, item := range ct.([]interface{}) {
			fmt.Println("List Item [", i, "]:")
			parseCreation(item)
		}
	case map[interface{}]interface{}:
		fmt.Println("creation map:")
		for k, v := range ct.(map[interface{}]interface{}) {
			fmt.Println(" -- Key:", k)
			fmt.Println(" -- Value:")
			handleMapCreation(k.(string), v)
			parseCreation(v)
		}
	}
	return nil
}

func (ty *TemplateYaml) ToTemplate() *Template {
	template := &Template{
		Name: ty.Name,
		Arguments: ty.Arguments,
	}
	ctlist := *ty.Creations
	ctlen := len(ctlist)
	creations := make([]Creation, ctlen)
	parseCreation(ctlist)
	template.Creations = &creations
	return template
}

func (ty *TemplateYaml) Print() {
	fmt.Println("Template Name:", ty.Name)
	if ty.Arguments != nil {
		fmt.Println("Template Arguments:", *ty.Arguments)
	}
	if ty.Resources != nil {
		fmt.Println("Template Resources:", *ty.Resources)
	}
	if ty.Creations != nil {
		fmt.Println("Template Creations:", *ty.Creations)
	}
	if ty.ExecCommands != nil {
		fmt.Println("Template Exec Commands:", *ty.ExecCommands)
	}
}
