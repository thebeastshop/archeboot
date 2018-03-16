
package data

import (
"strings"
"fmt"
)

var (
	TypeMap = map[string]FType {
		"dir": FTypeDir,
		"xml": FTypeXML,
		"txt": FTypeText,
		"json": FTypeJSON,
		"java": FTypeJava,
		"javaPackage": FTypeJavaPackage,
	}

	SupportedKeys = map[string]map[string]string {
		"dir": map[string]string {
			"copyFrom": "optional",
			},
		}
	)

type TypeHandler interface {
	Handle(typeName string, value string, args ...interface{}) *Creation
}


func handleStringCreation(value string) *Creation {
	fmt.Println("Handle String Creation -->")
	strs := strings.Split(value, ":")
	fmt.Println(strs)
	if len(strs) > 1 {
		return createCreation(strs[0], strs[1], nil)
	}
	return nil
}

func handleMapCreation(key string, value interface{}) *Creation {
	fmt.Println("Handle Map Creation -->")
	strs := strings.Split(key, ":")
	fmt.Println(strs)
	if len(strs) > 1 {
		return createCreation(strs[0], strs[1], nil)
	}
	return nil
}

func createCreation(typeName string, value string, args map[interface{}]interface{}) *Creation {
	ct := &Creation {
		Name: value,
		Type: TypeMap[typeName],
	}
	return ct
}
