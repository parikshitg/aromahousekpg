package controllers

import (
	"encoding/json"
	"fmt"
)

func Debug(obj interface{}) {
	bytes, err := json.MarshalIndent(obj, "    ", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
