package main

import (
	"archeboot/data"
	"fmt"
	"runtime"
)

func main() {
	var proj data.Project
	_, err := proj.Read()
	if err != nil {
		return
	}
	fmt.Println("OS:", runtime.GOOS)
	proj.Init()
	proj.Print()
	proj.Exec()
}
