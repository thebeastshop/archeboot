package main

import (
	"archeboot/data"
	"fmt"
	"runtime"
)

func main() {
	var proj data.Project
	config := data.NewConfig()
	fmt.Println("Config Path:", config.ConfigPath)
	_, err := proj.Read(config)
	if err != nil {
		return
	}
	fmt.Println("OS:", runtime.GOOS)
	proj.Init()
	proj.Print()
	proj.Exec()
}
