package main

import (
	"archeboot/data"
        "fmt"
)

func main() {
	config := data.NewConfig()
	fmt.Println("Config Path:", config.ConfigPath)
        template, err := data.ReadTemplateYaml(config)
	//_, err := proj.Read(config)
	if err != nil {
		return
	}
        template.Print()
        fmt.Println("------------------------------------")
        _ = template.ToTemplate()
	//proj.Init()
	//proj.Print()
	//proj.Exec()
}
