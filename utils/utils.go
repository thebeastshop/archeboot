package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LogError(err error) {
	log.Fatalf("[Archeboot] Error:  #%v", err)
}

func Println(info string) {
	log.Println("[Archboot] #" + info)
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		LogError(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Mkdir(path string) error {
	exists, _ := PathExists(path)
	if exists {
		return nil
	}
	err := os.MkdirAll(path, 0777)
	if err != nil {
		LogError(err)
	} else {
		Println("Create Floder " + path)
	}
	return err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
