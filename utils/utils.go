package utils

import (
	"log"
)

func LogError(err error) {
	log.Fatalf("[Archeboot] Error:  #%v", err)
}

func Println(info string) {
	log.Println("[Archboot] #" + info)
}
